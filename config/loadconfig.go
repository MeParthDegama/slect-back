package config

import (
	"encoding/json"
	"os"
)

var AppConfig Config

func LoadConfig() {
	
	fbyte, err := os.ReadFile(configFilePath)
	
	if err != nil {
		if os.IsNotExist(err) {
			MakeConfigFile()
		} else {
			panic(err)
		}
	}

	json.Unmarshal(fbyte, &AppConfig)

}
