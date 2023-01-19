package config

import (
	"encoding/json"
	"os"
)

var configFilePath = "./sconfig" // tmp

// if config file not exist than make config file
func MakeConfigFile() {

	// if file exist
	_, err := os.Stat(configFilePath)
	if err == nil {
		return
	}

	// make config structure
	sConfig := &Config{
		Tokens: []Token{},
	}

	confFileByte, err := json.Marshal(sConfig)
	if err != nil {
		// if error
		panic(err)
	}

	// wirte filee
	os.WriteFile(configFilePath, confFileByte, 0777)

}

func WriteConfig() {

	confFileByte, err := json.Marshal(AppConfig)
	if err != nil {
		// if error
		panic(err)
	}

	// wirte filee
	os.WriteFile(configFilePath, confFileByte, 0777)

}
