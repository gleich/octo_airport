package utils

import (
	"os"
)

// Ensure that a file or folder exists
func DoesExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
