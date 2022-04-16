package settings

import (
	"encoding/json"
	"os"
)

type Config struct {
	OsuPath           string `json:"OsuPath"`
	RemoveBackgrounds int    `json:"RemoveBackgrounds"`
}

func LoadSettings(configFilePath string) (Config, error) {
	var config Config

	// Open Config File
	configFile, err := os.Open(configFilePath)
	// if os.IsNotExist(err) {
	// 	// saveSettings(configFilePath, defaultConfig, true)
		
	// 	return config, err
	// }

	if err != nil {
		return config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}

