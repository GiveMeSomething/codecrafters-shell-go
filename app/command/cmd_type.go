package command

import (
	"fmt"
)

func HandleTypeCommand(cmd *CommandState) {
	targetCmd := ShellCommand(cmd.Args[0])
	if targetCmd.IsBuiltIn() {
		fmt.Printf("%s is a shell builtin\n", targetCmd)
		return
	}

	foundPath := SearchExecutable(string(targetCmd))
	if foundPath == nil {
		fmt.Printf("%s: not found\n", targetCmd)
		return
	}

	fmt.Printf("%s is %s\n", targetCmd, *foundPath)
}
