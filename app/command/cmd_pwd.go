package command

import (
	"fmt"
)

func HandlePwdCommand() {
	fmt.Println(GetShellState().CurrentDir)
}
