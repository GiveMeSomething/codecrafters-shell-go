package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Print("$ ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			command := scanner.Text()
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
