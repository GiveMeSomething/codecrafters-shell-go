package command

import (
	"fmt"
	"os"
	"os/exec"
)

func HandleOtherCommand(cmd string, args []string) {
	foundCmd := SearchExecutable(cmd)
	if foundCmd == nil {
		fmt.Printf("%s: command not found\n", cmd)
		return
	}

	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		fmt.Println(err)
	}
}
