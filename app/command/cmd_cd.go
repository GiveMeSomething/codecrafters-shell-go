package command

import "fmt"

func HandleCdCommand(cmd *CommandState) {
	if len(cmd.Args) != 1 {
		fmt.Println("cd accept only one parameter")
		return
	}
	GetShellState().Cd(cmd.Args[0])
}
