package utils

import (
	"strings"
	"testing"

	"github.com/tj/assert"
)

func TestDoesExist(t *testing.T) {
	const fName = "testing.txt"
	CreateTempFile(t, fName)
	assert.True(t, DoesExist(fName))
	RemoveTempFile(t, fName)
	assert.False(t, DoesExist(fName))
}

func TestReplaceRoot(t *testing.T) {
	instance := ReplaceRoot("~/.config")
	if !strings.Contains(instance, "/.config") {
		t.Error("Failed to replace root for .config folder")
	}
}
