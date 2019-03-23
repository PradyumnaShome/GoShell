package config

import (
	"bufio"
	"log"
	"os"
)

const (
	UserInputDelimiter  = " "
	UserInputTerminator = "\n"
)

var logger *log.Logger
var reader *bufio.Reader

func Init() error {
	logger = log.New(os.Stdin, "", 0)
	reader = bufio.NewReader(os.Stdin)
	return nil
}

func Contains(slice []string, query string) bool {
	for _, element := range slice {
		if element == query {
			return true
		}
	}
	return false
}

func GetBuiltInCommands() []string {
	return []string{"exit", "print", "find"}
}

func GetLogger() *log.Logger {
	return logger
}

func GetReader() *bufio.Reader {
	return reader
}
