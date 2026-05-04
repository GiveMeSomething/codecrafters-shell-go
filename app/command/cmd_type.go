package command

import (
	"fmt"
)

func HandleTypeCommand(cmd *CommandState) {
	if cmd.Command.IsBuiltIn() {
		fmt.Printf("%s is a shell builtin\n", cmd.Command)
		return
	}

	targetCmd := string(cmd.Command)
	foundPath := SearchExecutable(targetCmd)
	if foundPath == nil {
		fmt.Printf("%s: not found\n", targetCmd)
		return
	}

	fmt.Printf("%s is %s\n", targetCmd, *foundPath)
}
