package config

import (
	"io/ioutil"

	"github.com/Matt-Gleich/octo_airport/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"gopkg.in/yaml.v3"
)

type Outline struct {
	PAT      string `yaml:"token"`
	Username string `yaml:"username"`
}

var validFilePaths = []string{"~/.config/octo_airport/config.yml", "~/.octo_airport.yml"}

// Get the config from the config file
func Get() Outline {
	// Getting file path for config
	var configPath string
	for i, path := range validFilePaths {
		fixedPath := utils.ReplaceRoot(path)
		if utils.DoesExist(fixedPath) {
			configPath = fixedPath
			break
		}
		if i == len(validFilePaths)-1 {
			statuser.ErrorMsg("Failed to find config file", 1)
		}
	}
	// Reading from the file
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		statuser.Error("Failed to read from config file", err, 1)
	}
	// Parsing the yaml
	var yamlFile Outline
	err = yaml.Unmarshal(bytes, &yamlFile)
	if err != nil {
		statuser.Error("Failed to parse the yaml from the config file", err, 1)
	}
	return yamlFile
}
