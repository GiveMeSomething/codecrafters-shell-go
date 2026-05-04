package command

import (
	"fmt"
	"strings"
)

func HandleEchoCommand(cmd *CommandState) {
	fmt.Println(strings.Join(cmd.Args, " "))
}
