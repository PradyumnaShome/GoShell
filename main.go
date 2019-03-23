package main

import (
	"GoShell/commands"
	"GoShell/config"
	"GoShell/format"
	"errors"
	"os"
	"strings"
)

func main() {
	err := InitializeShell()

	if err != nil {
		os.Exit(1)
	}

	for {
		format.PrintShellPrompt()

		userInput, err := GetUserInput()

		if err != nil {
			format.PrintInvalidCommand(userInput)
			continue
		}

		ProcessUserInput(userInput)
	}
}

func InitializeShell() error {
	// Potentially load a configuration file
	return config.Init()
}

// Returns everything inputs up to and excluding the userInputTerminator
func GetUserInput() (string, error) {
	userInput, err := config.GetReader().ReadString(config.UserInputTerminator)

	if err != nil {
		return "", err
	}

	if len(userInput) == 1 {
		return "", nil
	}

	return userInput[:len(userInput)-1], nil
}

func ProcessUserInput(userInput string) error {
	tokens := strings.Split(userInput, config.UserInputDelimiter)

	if len(tokens) < 1 {
		return nil
	}

	command := tokens[0]

	isCommandBuiltIn := commands.IsCommandBuiltIn(command)

	if isCommandBuiltIn {
		switch command {
		case "print":
			return commands.Print(tokens, config.UserInputDelimiter)
		case "find":
			return commands.Find(tokens)
		case "exit":
			commands.Exit()
		default:
			format.PrintInvalidCommand(userInput)
			return errors.New("Invalid command.")
		}
	} else {
		return commands.ExternalCommand(tokens)
	}

	return nil
}
