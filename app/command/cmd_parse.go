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

	parseArgs := ParseCommand(args[0])
	output := strings.Builder{}
	for i, arg := range parseArgs {
		fmt.Fprintf(&output, "\"%s\"", arg)
		if i != len(parseArgs)-1 {
			output.WriteString(", ")
		}
	}

	fmt.Printf("[%s]\n", output.String())
}
