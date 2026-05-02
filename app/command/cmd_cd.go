package command

import "fmt"

func HandleCdCommand(args []string) {
	if len(args) != 1 {
		fmt.Println("cd accept only one parameter")
		return
	}
	GetShellState().Cd(args[0])
}
