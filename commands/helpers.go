package commands

import "GoShell/config"

func IsCommandBuiltIn(command string) (isCommandBuiltIn bool) {
	_, isCommandBuiltIn = config.GetBuiltInCommands()[command]
	return
}

func findHelper(directory string, file string) ([]string, error) {
	return []string{"Not implemented yet."}, nil
}
