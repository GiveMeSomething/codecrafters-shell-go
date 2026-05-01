package main

import (
	"bufio"
	"fmt"
	"os"
)

type ShellCommand string

const (
	CommandExit ShellCommand = "exit"
)

func handleCommand(command ShellCommand) {
	switch command {
	case CommandExit:
		os.Exit(0)
		return
	}

	fmt.Printf("%s: command not found\n", command)
}

func main() {
	for {
		fmt.Print("$ ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			command := scanner.Text()
			handleCommand(ShellCommand(command))
		}
	}
}
