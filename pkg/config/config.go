package config

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/octo_airport/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"gopkg.in/yaml.v3"
)

type ConfigOutline struct {
	PAT string
}

type yamlConfigOutline struct {
	PAT string `yaml:"token"`
}

// Get the configuration
func Get() ConfigOutline {
	envToken := os.Getenv("OCTO_AIRPORT_TOKEN")
	if envToken != "" {
		return ConfigOutline{PAT: envToken}
	}
	return getFile()
}

var validFilePaths = []string{"/Users/matthewgleich/.config/octo_airport/config.yml", "/Users/matthewgleich/.octo_airport.yml"}

// Get the config from the config file
func getFile() ConfigOutline {
	// Getting file path for config
	var configPath string
	for i, path := range validFilePaths {
		if utils.DoesExist(path) {
			configPath = path
			break
		}
		if i == len(validFilePaths) {
			statuser.ErrorMsg("Failed to find config file", 1)
		}
	}
	// Reading from the file
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		statuser.Error("Failed to read from config file", err, 1)
	}
	// Parsing the yaml
	var yamlFile yamlConfigOutline
	err = yaml.Unmarshal(bytes, &yamlFile)
	if err != nil {
		statuser.Error("Failed to parse the yaml from the config file", err, 1)
	}
	return ConfigOutline(yamlFile)
}
