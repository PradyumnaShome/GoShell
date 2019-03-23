package config

import (
	"bufio"
	"log"
	"os"
)

const (
	UserInputTerminator = '\n'
	UserInputDelimiter  = " "
)

var logger *log.Logger
var reader *bufio.Reader

func Init() error {
	logger = log.New(os.Stdin, "", 0)
	reader = bufio.NewReader(os.Stdin)
	return nil
}

func GetBuiltInCommands() map[string]bool {
	return map[string]bool{"exit": true, "print": true, "find": true}
}

func GetLogger() *log.Logger {
	return logger
}

func GetReader() *bufio.Reader {
	return reader
}
