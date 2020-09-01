package utils

import (
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
