package format

import (
	"fmt"
	"os"
)

func PrintShellPrompt() {
	currentWorkingDirectory, err := os.Getwd()

	if err != nil {
		currentWorkingDirectory = "CurrentWorkingDirectory"
	}

	shellPrompt := currentWorkingDirectory + ">"
	fmt.Print(shellPrompt)
}

func PrintInvalidCommand(userInput string) {
	fmt.Printf("The command `%v` was invalid.\n", userInput)
}
