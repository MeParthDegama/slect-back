package config

import (
	"encoding/json"
	"os"
)

func AddToken(t string, u string) {

	conf := &Config{}

	fbyte, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			MakeConfigFile()
		} else {
			panic(err)
		}
	}

	// add token
	json.Unmarshal(fbyte, &conf)
	conf.Tokens = append(conf.Tokens, Token{t, u})
	WriteConfig(conf)

}
