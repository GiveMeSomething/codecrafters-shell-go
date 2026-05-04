package command

import (
	"os"
	"strings"
)

type CommandState struct {
	Command string
	Args    []string
	Stdout  *os.File
}

func ParseCommand(input string) []string {
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

	return parseResult
}
