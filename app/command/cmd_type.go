package command

import (
	"fmt"
)

func HandleTypeCommand(args []string) {
	targetCmd := args[0]

	shellCmd := ShellCommand(targetCmd)
	if shellCmd.IsBuiltIn() {
		fmt.Printf("%s is a shell builtin\n", shellCmd)
		return
	}

	foundPath := SearchExecutable(targetCmd)
	if foundPath == nil {
		fmt.Printf("%s: not found\n", targetCmd)
		return
	}

	fmt.Printf("%s is %s\n", targetCmd, *foundPath)
}
