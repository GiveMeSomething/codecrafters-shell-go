package command

import (
	"fmt"
	"strings"
)

func HandleEchoCommand(cmd *CommandState) {
	_, err := fmt.Fprintln(cmd.Stdout, strings.Join(cmd.Args, " "))
	if err != nil {
		fmt.Println(err)
	}
}
