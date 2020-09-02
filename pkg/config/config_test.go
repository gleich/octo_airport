package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Matt-Gleich/octo_airport/pkg/utils"
	"github.com/tj/assert"
)

const fakeToken = "18947302981074238973412"

func TestGet(t *testing.T) {
	err := os.MkdirAll(utils.ReplaceRoot("~/.config/octo_airport/"), 0700)
	utils.CheckTestingErr(t, err)
	for _, filePath := range validFilePaths {
		fixedPath := utils.ReplaceRoot(filePath)
		utils.CreateTempFile(t, fixedPath)
		err := ioutil.WriteFile(fixedPath, []byte("token: \""+fakeToken+"\"\n"), 0644)
		utils.CheckTestingErr(t, err)
		instance := Get()
		assert.Equal(t, Outline{PAT: fakeToken}, instance)
		utils.RemoveTempFile(t, fixedPath)
	}
}
