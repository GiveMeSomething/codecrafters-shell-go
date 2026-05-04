package command

import (
	"fmt"
	"os"
)

func HandleCommand(input string) {
	parsedCommand, err := ParseCommand(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch parsedCommand.Command {
	case CommandExit:
		os.Exit(0)
	case CommandEcho:
		HandleEchoCommand(parsedCommand)
		return
	case CommandType:
		HandleTypeCommand(parsedCommand)
		return
	case CommandPwd:
		HandlePwdCommand(parsedCommand)
		return
	case CommandCd:
		HandleCdCommand(parsedCommand)
		return
	case CommandParse:
		HandleParseCommand(parsedCommand)
		return
	default:
		HandleOtherCommand(parsedCommand)
		return
	}

	fmt.Printf("%s: command not found\n", parsedCommand.Command)
}
