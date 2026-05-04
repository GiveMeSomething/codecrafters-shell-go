package command

type ShellCommand string

const (
	CommandExit  ShellCommand = "exit"
	CommandEcho  ShellCommand = "echo"
	CommandType  ShellCommand = "type"
	CommandPwd   ShellCommand = "pwd"
	CommandCd    ShellCommand = "cd"
	CommandParse ShellCommand = "parse"
)

func (cmd ShellCommand) IsBuiltIn() bool {
	return cmd == CommandExit ||
		cmd == CommandEcho ||
		cmd == CommandType ||
		cmd == CommandPwd ||
		cmd == CommandParse
}

func IsStdoutRedirect(symbol string) bool {
	return symbol == ">" || symbol == "1>"
}
