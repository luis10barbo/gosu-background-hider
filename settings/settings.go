package settings

import (
	"encoding/json"
	"os"
)

var Config ConfigStruct
type ConfigStruct struct {
	OsuPath           string `json:"OsuPath"`
	RemoveBackgrounds int    `json:"RemoveBackgrounds"`
	DesktopNotifications int `json:"DesktopNotifications"`
}

func LoadSettings() (ConfigStruct, error) {
	
	configFilePath := "settings.json"
	
	// Open Config File
	configFile, err := os.Open(configFilePath)
	// if os.IsNotExist(err) {
	// 	// saveSettings(configFilePath, defaultConfig, true)
		
	// 	return config, err
	// }

	if err != nil {
		return Config, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&Config)

	return Config, err
}

