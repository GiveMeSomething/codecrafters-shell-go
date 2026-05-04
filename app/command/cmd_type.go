package command

import (
	"fmt"
)

func HandleTypeCommand(cmd *CommandState) {
	targetCmd := string(cmd.Args[0])
	if cmd.Command.IsBuiltIn() {
		fmt.Printf("%s is a shell builtin\n", targetCmd)
		return
	}

	foundPath := SearchExecutable(targetCmd)
	if foundPath == nil {
		fmt.Printf("%s: not found\n", targetCmd)
		return
	}

	fmt.Printf("%s is %s\n", targetCmd, *foundPath)
}
