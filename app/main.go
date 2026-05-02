package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/command"
)

func main() {
	// Init a new shell state
	_ = command.GetShellState()

	for {
		fmt.Print("$ ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			// Read and handle user input
			command.HandleCommand(scanner.Text())
		}
	}
}
