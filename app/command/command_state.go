package command

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

// This is to control the current shell state

type ShellState struct {
	CurrentDir string
}

var lock sync.Mutex

var currentShellState *ShellState

func GetShellState() *ShellState {
	if currentShellState == nil {
		lock.Lock()
		defer lock.Unlock()

		currentShellState = &ShellState{}
		currentShellState.Init()
	}

	return currentShellState
}

func (state *ShellState) Init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	state.CurrentDir = cwd
}

func (state *ShellState) Cd(path string) {
	// Handle absolute path
	if strings.HasPrefix(path, "/") {
		// Check path exist
		_, err := os.Stat(path)
		if err != nil {
			// Print additional information when not exist
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("cd: %s: No such file or directory\n", path)
				return
			}

			// Just skip the cd operation when there're error
			fmt.Printf("cd failed: %s", err)
			return
		}
		state.CurrentDir = path
	}
}
