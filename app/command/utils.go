package command

import "os"

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
