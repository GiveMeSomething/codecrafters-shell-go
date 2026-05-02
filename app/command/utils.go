package command

import (
	"os"
	"path"
	"strings"
)

// UNIX permission bits
// Owner - Group - Other
// rwx - rwx - rwx

func IsExecOwner(mode os.FileMode) bool {
	return mode&0100 != 0
}

func IsExecGroup(mode os.FileMode) bool {
	return mode&0010 != 0
}

func IsExecOther(mode os.FileMode) bool {
	return mode&0001 != 0
}

func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

func SearchExecutable(targetCmd string) *string {
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

		return &checkPath
	}
	return nil
}
