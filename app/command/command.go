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
		HandleEchoCommand(parsedCommand.Args)
		return
	case CommandType:
		HandleTypeCommand(parsedCommand.Args)
		return
	case CommandPwd:
		HandlePwdCommand()
		return
	case CommandCd:
		HandleCdCommand(parsedCommand.Args)
		return
	case CommandParse:
		HandleParseCommand(parsedCommand.Args)
		return
	default:
		HandleOtherCommand(string(parsedCommand.Command), parsedCommand.Args)
		return
	}

	fmt.Printf("%s: command not found\n", parsedCommand.Command)
}
