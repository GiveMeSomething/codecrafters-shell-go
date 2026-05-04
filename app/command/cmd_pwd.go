package command

import (
	"fmt"
)

func HandlePwdCommand(cmd *CommandState) {
	fmt.Fprintln(cmd.Stdout, GetShellState().CurrentDir)
}
