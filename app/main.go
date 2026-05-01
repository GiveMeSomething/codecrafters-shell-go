package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ShellCommand string

const (
	CommandExit ShellCommand = "exit"
	CommandEcho ShellCommand = "echo"
	CommandType ShellCommand = "type"
)

func (cmd ShellCommand) IsBuiltIn() bool {
	return cmd == CommandExit || cmd == CommandEcho || cmd == CommandType
}

func handleCommand(command ShellCommand, args []string) {
	switch command {
	case CommandExit:
		os.Exit(0)
		return
	case CommandEcho:
		fmt.Println(strings.Join(args, " "))
		return
	case CommandType:
		typeArg := ShellCommand(args[0])
		if typeArg.IsBuiltIn() {
			fmt.Printf("%s is a shell builtin\n", typeArg)
			return
		}
		fmt.Printf("%s: not found\n", typeArg)
		return
	}

	fmt.Printf("%s: command not found\n", command)
}

func main() {
	for {
		fmt.Print("$ ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			args := strings.Split(scanner.Text(), " ")
			if len(args) == 1 {
				handleCommand(ShellCommand(args[0]), []string{})
				continue
			}

			handleCommand(ShellCommand(args[0]), args[1:])
		}
	}
}
