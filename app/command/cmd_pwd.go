package command

import (
	"fmt"
	"os"
)

func HandlePwdCommand() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to get current working directory: %s", err)
		return
	}

	fmt.Println(pwd)
}
