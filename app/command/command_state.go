package command

import (
	"errors"
	"os"
	"strings"
)

type CommandState struct {
	Command ShellCommand
	Args    []string
	Stdout  *os.File
	Stderr  *os.File
}

func ParseCommand(input string) (*CommandState, error) {
	parseResult := []string{}
	buffer := strings.Builder{}

	singleQuoteOpen := false
	doubleQuoteOpen := false
	backlashEnabled := false

	for _, char := range input {
		// Empty space usually mean the next part of the command
		// Unless there's opening single/double quote
		if char == ' ' {
			// Ignore empty part of the command
			if buffer.Len() == 0 {
				continue
			}

			if singleQuoteOpen || doubleQuoteOpen {
				buffer.WriteRune(char)
				continue
			}

			if backlashEnabled {
				buffer.WriteRune(char)
				backlashEnabled = false
				continue
			}

			parseResult = append(parseResult, buffer.String())
			buffer.Reset()
			continue
		}

		if char == '\'' {
			if doubleQuoteOpen || backlashEnabled {
				buffer.WriteRune(char)
				backlashEnabled = false
				continue
			}
			singleQuoteOpen = !singleQuoteOpen
			continue
		}

		if char == '"' {
			if singleQuoteOpen || backlashEnabled {
				buffer.WriteRune(char)
				backlashEnabled = false
				continue
			}
			doubleQuoteOpen = !doubleQuoteOpen
			continue
		}

		if char == '\\' {
			if singleQuoteOpen {
				buffer.WriteRune(char)
				continue
			}
			if backlashEnabled {
				buffer.WriteRune(char)
				backlashEnabled = false
				continue
			}
			backlashEnabled = true
			continue
		}

		buffer.WriteRune(char)
		backlashEnabled = false
	}

	if buffer.Len() != 0 {
		parseResult = append(parseResult, buffer.String())
	}

	stdout, stdoutCutAt, err := GetStdout(parseResult)
	if err != nil {
		return nil, err
	}
	stderr, stderrCutAt, err := GetStderr(parseResult)
	if err != nil {
		return nil, err
	}

	if len(parseResult) == 1 {
		return &CommandState{
			Command: ShellCommand(parseResult[0]),
			Args:    []string{},
			Stdout:  stdout,
			Stderr:  stderr,
		}, nil
	}

	return &CommandState{
		Command: ShellCommand(parseResult[0]),
		Args:    parseResult[1:min(stdoutCutAt, stderrCutAt)],
		Stdout:  stdout,
		Stderr:  stderr,
	}, nil
}

// Return output, cut-off index, error
// The cut-off index is to determine where is the command, and the rest of the redirection
func GetStdout(args []string) (*os.File, int, error) {
	for i, token := range args {
		if !IsStdoutRedirect(token) && !IsStdoutAppend(token) {
			continue
		}

		// The next arg must be a path to output
		if i+1 >= len(args) {
			return nil, 0, errors.New("Missing argument for redirection")
		}

		outputPath, err := GetShellState().GetFinalPath(args[i+1])
		if err != nil {
			return nil, 0, err
		}

		var fileMode = os.O_CREATE | os.O_APPEND | os.O_WRONLY

		var redirectTo *os.File
		if IsStderrRedirect(token) {
			redirectTo, err = os.OpenFile(outputPath, fileMode|os.O_TRUNC, os.ModePerm)
		} else {
			// Handle append to stdout
			redirectTo, err = os.OpenFile(outputPath, fileMode, os.ModePerm)
		}

		if err != nil {
			return nil, 0, err
		}
		return redirectTo, i, nil

	}

	// Default to return to stdout
	return os.Stdout, len(args), nil
}

func GetStderr(args []string) (*os.File, int, error) {
	for i, token := range args {
		if !IsStderrRedirect(token) {
			continue
		}

		// The next arg must be a path to output
		if i+1 >= len(args) {
			return nil, 0, errors.New("Missing argument for redirection")
		}

		outputPath, err := GetShellState().GetFinalPath(args[i+1])
		if err != nil {
			return nil, 0, err
		}

		redirectTo, err := os.OpenFile(outputPath, os.O_CREATE|os.O_APPEND|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return nil, 0, err
		}
		return redirectTo, i, nil
	}

	// Default to return to stdout
	return os.Stderr, len(args), nil
}
