package utils

import "testing"

func TestTempFuncs(t *testing.T) {
	const fName = "testing.txt"
	CreateTempFile(t, fName)
	if !DoesExist(fName) {
		t.Error("Failed to create temp file")
	}
	RemoveTempFile(t, fName)
	if DoesExist(fName) {
		t.Error("Failed to remove temp file")
	}
}
