package utils

import (
	"os"
	"testing"
)

// Create a temporary environment for testing
func CreateTempFile(t *testing.T, fName string) {
	f, err := os.Create(fName)
	CheckTestingErr(t, err)
	err = f.Close()
	CheckTestingErr(t, err)
}

// Remove the temporary environment for testing
func RemoveTempFile(t *testing.T, fName string) {
	err := os.Remove(fName)
	CheckTestingErr(t, err)
}

func CheckTestingErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
