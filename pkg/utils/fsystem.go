package utils

import (
	"os"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
)

// Ensure that a file or folder exists
func DoesExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Replace ~ in a path to actual root path
func ReplaceRoot(path string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get user home directory", err, 1)
	}
	if string([]rune(path)[0]) == "~" {
		return strings.Replace(path, "~", homeDir, 1)
	}
	return path
}
