package command_test

import (
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/command"
)

func TestParseCommand_HappyCase(t *testing.T) {
	resultMap := map[string][]string{
		// Should trim all white space and keep only the good command part
		"hello world":       {"hello", "world"},
		"hello    world":    {"hello", "world"},
		"  hello world":     {"hello", "world"},
		"  hello   world":   {"hello", "world"},
		"  hello   world  ": {"hello", "world"},

		// Should understand single quote
		"echo 'a very   long hello world'": {"echo", "a very   long hello world"},
		"echo 'hello''world'":              {"echo", "helloworld"},
		"echo hello''world":                {"echo", "helloworld"},

		// Shoulder understand double quote
		"echo \"hello    world\"": {"echo", "hello    world"},
		"echo \"hello\"\"world\"": {"echo", "helloworld"},
		"echo \"hello\"world":     {"echo", "helloworld"},
	}

	for input, expected := range resultMap {
		output, err := command.ParseCommand(input)
		if err != nil {
			t.Error(err)
			return
		}

		if len(output.Args)+1 != len(expected) {
			t.Errorf("output length mismatched. expected %+v. received %+v", expected, output.Args)
			return
		}

		for i, value := range expected {
			// The first expected value should be the command
			if i == 0 && value != string(output.Command) {
				t.Errorf("parse failed. expected %+v. received %+v", expected, output.Args)
				return
			}

			// The rest should be the arguments
			if i != 0 && expected[i] != value {
				t.Errorf("parse failed. expected %+v. received %+v", expected, output.Args)
				return
			}
		}
	}
}
