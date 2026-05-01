package command

import (
	"fmt"
	"strings"
)

func HandleEchoCommand(args []string) {
	fmt.Println(strings.Join(args, " "))
}
