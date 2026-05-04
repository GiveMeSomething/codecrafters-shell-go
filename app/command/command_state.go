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
}

func ParseCommand(input string) (*CommandState, error) {
	parseResult := []string{}
	buffer := strings.Builder{}

	singleQuoteOpen := false
	doubleQuoteOpen := false
	backlashEnabled := false

	stdout := os.Stdout

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

	stdout, err := GetStdout(parseResult)
	if err != nil {
		return nil, err
	}

	if len(parseResult) == 1 {
		return &CommandState{
			Command: ShellCommand(parseResult[0]),
			Args:    []string{},
			Stdout:  stdout,
		}, nil
	}

	return &CommandState{
		Command: ShellCommand(parseResult[0]),
		Args:    parseResult[1:],
		Stdout:  stdout,
	}, nil
}

func GetStdout(args []string) (*os.File, error) {
	for i, token := range args {
		if !IsStdoutRedirect(token) {
			continue
		}

		// The next arg must be a path to output
		if i+1 >= len(args) {
			return nil, errors.New("Missing argument for redirection")
		}

		outputPath, err := GetShellState().GetFinalPath(args[i+1])
		if err != nil {
			return nil, err
		}

		redirectTo, err := os.OpenFile(outputPath, os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			return nil, err
		}
		return redirectTo, nil
	}

	// Default to return to stdout
	return os.Stdout, nil
}
