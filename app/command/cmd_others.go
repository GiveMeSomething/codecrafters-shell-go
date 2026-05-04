package command

import (
	"fmt"
	"os"
	"os/exec"
)

func HandleOtherCommand(cmd *CommandState) {
	foundCmd := SearchExecutable(string(cmd.Command))
	if foundCmd == nil {
		fmt.Printf("%s: command not found\n", string(cmd.Command))
		return
	}

	command := exec.Command(string(cmd.Command), cmd.Args...)
	command.Stdout = cmd.Stdout
	command.Stderr = os.Stderr

	_ = command.Run()
}
