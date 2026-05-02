package command_test

import (
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/command"
)

func TestGetFinalPath_Absolute(t *testing.T) {
	pwd := "/Users/mint/code/codecrafters-shell-go"
	input := "/Users/mock"
	expected := "/Users/mock"

	state := command.ShellState{
		CurrentDir: pwd,
	}

	output, err := state.GetFinalPath(input)
	if err != nil {
		t.Error(err)
		return
	}

	if output != expected {
		t.Errorf("Final path mismatched. Expected: %s. Received: %s", expected, output)
	}
}

// This test only work on my machine lol
func TestGetFinalPath_HomeDir(t *testing.T) {
	pwd := "/Users/mint/code/codecrafters-shell-go"

	resultMap := map[string]string{
		"~":           "/Users/mint",
		"~/code":      "/Users/mint/code",
		"~/code/test": "/Users/mint/code/test",
	}

	for input, expected := range resultMap {
		state := command.ShellState{
			CurrentDir: pwd,
		}

		output, err := state.GetFinalPath(input)
		if err != nil {
			t.Error(err)
			return
		}

		if output != expected {
			t.Errorf("Final path mismatched. Expected: %s. Received: %s", expected, output)
		}
	}
}

func TestGetFinalPath_Relative(t *testing.T) {
	pwd := "/Users/mint/code/codecrafters-shell-go"
	resultMap := map[string]string{
		"..":     "/Users/mint/code",
		".":      "/Users/mint/code/codecrafters-shell-go",
		"./":     "/Users/mint/code/codecrafters-shell-go",
		"./test": "/Users/mint/code/codecrafters-shell-go/test",
	}

	for input, expected := range resultMap {
		state := command.ShellState{
			CurrentDir: pwd,
		}

		output, err := state.GetFinalPath(input)
		if err != nil {
			t.Error(err)
			return
		}

		if output != expected {
			t.Errorf("Final path mismatched. Expected: %s. Received: %s", expected, output)
		}
	}
}
