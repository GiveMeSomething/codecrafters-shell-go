package command

import (
	"fmt"
	"os"
	"strings"
)

type ShellCommand string

const (
	CommandExit ShellCommand = "exit"
	CommandEcho ShellCommand = "echo"
	CommandType ShellCommand = "type"
	CommandPwd  ShellCommand = "pwd"
	CommandCd   ShellCommand = "cd"
)

func (cmd ShellCommand) IsBuiltIn() bool {
	return cmd == CommandExit || cmd == CommandEcho || cmd == CommandType || cmd == CommandPwd
}

func ParseCommand(input string) []string {
	parseResult := []string{}
	buffer := strings.Builder{}

	singleQuoteOpen := false

	for _, char := range input {
		// Empty space usually mean the next part of the command
		// Unless there's opening single/double quote
		if char == ' ' {
			// Ignore empty part of the command
			if buffer.Len() == 0 {
				continue
			}

			if singleQuoteOpen {
				buffer.WriteRune(char)
				continue
			}

			parseResult = append(parseResult, buffer.String())
			buffer.Reset()
			continue
		}

		if char == '\'' {
			singleQuoteOpen = !singleQuoteOpen
			continue
		}

		buffer.WriteRune(char)
	}

	if buffer.Len() != 0 {
		parseResult = append(parseResult, buffer.String())
	}

	return parseResult
}

func HandleCommand(input string) {
	args := ParseCommand(input)

	cmd := ShellCommand(args[0])
	cmdArgs := func() []string {
		if len(args) == 0 {
			return []string{}
		}
		return args[1:]
	}()

	switch cmd {
	case CommandExit:
		os.Exit(0)
	case CommandEcho:
		HandleEchoCommand(cmdArgs)
		return
	case CommandType:
		HandleTypeCommand(cmdArgs)
		return
	case CommandPwd:
		HandlePwdCommand()
		return
	case CommandCd:
		HandleCdCommand(cmdArgs)
		return
	default:
		HandleOtherCommand(string(cmd), cmdArgs)
		return
	}

	fmt.Printf("%s: command not found\n", cmd)
}
