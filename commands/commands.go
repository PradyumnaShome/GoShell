package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Find recursively looks for a file in `directory` and all subdirectories.
// The second element of the slice should be the directory in which to search
// The third element of the slice should be the file / directory being searched for
// This function returns all matches.
func Find(inputTokens []string) error {
	if len(inputTokens) < 3 {
		return errors.New("Not enough arguments to find.")
	}

	rootDirectory := inputTokens[1]
	file := inputTokens[2]

	matches, err := findHelper(rootDirectory, file)

	if err != nil {
		return err
	}

	if len(matches) == 0 {
		fmt.Println("Could not find the file.")
		return nil
	}

	for _, match := range matches {
		fmt.Println(match)
	}

	return nil
}

func Print(inputTokens []string, userInputDelimiter string) error {
	for _, token := range inputTokens[1:] {
		fmt.Print(token, userInputDelimiter)
	}
	fmt.Println()
	return nil
}

func ExternalCommand(inputTokens []string) error {
	log.Println("External command")

	if len(inputTokens) < 1 {
		return nil
	}

	command := exec.Command(inputTokens[0], inputTokens[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func Exit() {
	os.Exit(0)
}
