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
)

func handleCommand(command ShellCommand, args []string) {
	switch command {
	case CommandExit:
		os.Exit(0)
		return
	case CommandEcho:
		fmt.Println(strings.Join(args, " "))
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
				return
			}

			handleCommand(ShellCommand(args[0]), args[1:])
		}
	}
}
