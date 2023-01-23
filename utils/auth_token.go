package utils

import (
	"errors"

	"github.com/parthkax70/slect/config"
)


// auth token (token is valid or not)
func AuthToken(token string) (string, error) {

	for _, v := range config.AppConfig.Tokens {
		if v.Token == token {
			return v.Username, nil
		}
	}

	return "", errors.New("login error")

}
