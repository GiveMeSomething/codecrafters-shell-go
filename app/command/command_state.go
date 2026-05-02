package command

import (
	"errors"
	"fmt"
	"os"
	"path"
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

func (state *ShellState) GetFinalPath(userInput string) (string, error) {
	if strings.HasPrefix(userInput, "/") {
		return userInput, nil
	}

	if strings.HasPrefix(userInput, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		return path.Join(homeDir, strings.ReplaceAll(userInput, "~", ".")), nil
	}

	return path.Join(state.CurrentDir, userInput), nil
}

func (state *ShellState) Cd(cdPath string) {
	// Handle absolute path
	targetPath, err := state.GetFinalPath(cdPath)
	if err != nil {
		return
	}

	info, err := os.Stat(cdPath)
	if err != nil {
		// Print additional information when not exist
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("cd: %s: No such file or directory\n", cdPath)
			return
		}

		// Just skip the cd operation when there're error
		fmt.Printf("cd failed: %s", err)
		return
	}

	// Must be a directory
	if !info.IsDir() {
		return
	}

	state.CurrentDir = targetPath
}
