package main

import (
	"GoShell/commands"
	"GoShell/config"
	"GoShell/format"
	"errors"
	"fmt"
	"io"
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
	userInput := ""
	for {
		rune, _, err := config.GetReader().ReadRune()

		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			runeString := fmt.Sprintf("%c", rune)
			fmt.Printf("Rune String: {%v}\n", runeString)

			switch runeString {
			case "\t":
				TabComplete(userInput)
			case "\n":
				break
			default:
				userInput += runeString
			}
		}
	}

	if len(userInput) == 0 {
		return "", nil
	}

	return userInput, nil
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

func TabComplete(userInput string) {
	// searchList := []string{}

	// Look through internal commands, and add those to the list
	// Sort the list, and autofill
	for _, element := range config.GetBuiltInCommands() {
		if strings.HasPrefix(element, userInput) {
			// Write the remaining to standard out
			startingIndex := len(userInput)
			remaining := element[startingIndex:]
			fmt.Print(remaining)
			return
		}
	}
	// Look in $PATH, and find all commands starting with userInput
	return
}
