package command

import (
	"fmt"
	"os"
)

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
	case CommandParse:
		HandleParseCommand(cmdArgs)
		return
	default:
		HandleOtherCommand(string(cmd), cmdArgs)
		return
	}

	fmt.Printf("%s: command not found\n", cmd)
}
