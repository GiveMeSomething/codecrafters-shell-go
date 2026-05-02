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
)

func (cmd ShellCommand) IsBuiltIn() bool {
	return cmd == CommandExit || cmd == CommandEcho || cmd == CommandType || cmd == CommandPwd
}

func HandleCommand(input string) {
	args := strings.Split(input, " ")

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
	default:
		HandleOtherCommand(string(cmd), cmdArgs)
		return
	}

	fmt.Printf("%s: command not found\n", cmd)
}
