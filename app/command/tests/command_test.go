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
	}

	for input, expected := range resultMap {
		output := command.ParseCommand(input)
		if len(output) != len(expected) {
			t.Errorf("output length mismatched. expected %+v. received %+v", expected, output)
			return
		}

		for i, value := range output {
			if expected[i] != value {
				t.Errorf("parse failed. expected %+v. received %+v", expected, output)
				return
			}
		}
	}
}
