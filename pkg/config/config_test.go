package config

import (
	"io/ioutil"
	"testing"

	"github.com/Matt-Gleich/octo_airport/pkg/utils"
	"github.com/tj/assert"
)

const fakeToken = "18947302981074238973412"

func TestGetFile(t *testing.T) {
	for _, filePath := range validFilePaths {
		utils.CreateTempFile(t, filePath)
		err := ioutil.WriteFile(filePath, []byte("token: \""+fakeToken+"\"\n"), 0644)
		utils.CheckTestingErr(t, err)
		instance := getFile()
		assert.Equal(t, ConfigOutline{PAT: fakeToken}, instance)
		utils.RemoveTempFile(t, filePath)
	}
}
