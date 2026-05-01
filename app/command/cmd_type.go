package command

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func HandleTypeCommand(args []string) {
	targetCmd := args[0]

	shellCmd := ShellCommand(targetCmd)
	if shellCmd.IsBuiltIn() {
		fmt.Printf("%s is a shell builtin", shellCmd)
		return
	}

	pathEnv := os.Getenv("PATH")
	folders := strings.SplitSeq(pathEnv, ":")
	for folderName := range folders {
		checkPath := path.Join(folderName, targetCmd)
		fileInfo, err := os.Stat(checkPath)
		if err != nil {
			// Ignore error, continue checking
			continue
		}

		if !IsExecAny(fileInfo.Mode()) {
			// If the file exists but lacks execute permissions, skip it and continue to the next directory
			continue
		}

		fmt.Printf("%s is %s\n", targetCmd, checkPath)
		return
	}

	fmt.Printf("%s: not found\n", targetCmd)
}
