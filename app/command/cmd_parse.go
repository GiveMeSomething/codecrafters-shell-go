package command

import (
	"fmt"
	"strings"
)

func HandleParseCommand(args []string) {
	if len(args) != 1 {
		fmt.Println("parse only accept 1 argument")
		return
	}

	parsedCommand, err := ParseCommand(args[0])
	if err != nil {
		fmt.Printf("parse failed. error: %s\n", err)
		return
	}

	output := strings.Builder{}
	for i, arg := range parsedCommand.Args {
		fmt.Fprintf(&output, "\"%s\"", arg)
		if i != len(parsedCommand.Args)-1 {
			output.WriteString(", ")
		}
	}

	fmt.Printf("[%s]\n", output.String())
}
