package commands

import "GoShell/config"

func IsCommandBuiltIn(command string) bool {
	return config.Contains(config.GetBuiltInCommands(), command)
}

func findHelper(directory string, file string) ([]string, error) {
	return []string{"Not implemented yet."}, nil
}
