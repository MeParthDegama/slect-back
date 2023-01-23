package utils

import "github.com/parthkax70/slect/config"

// get user fullname
func FullName(username string) string {
	for _, v := range config.AppConfig.FullNames {
		if v.Username == username {
			return v.FullName
		}
	}
	return username
}

// set user fullname
func SetFullName(username string, fullname string) {

	for i, v := range config.AppConfig.FullNames {
		if v.Username == username {
			config.AppConfig.FullNames[i].FullName = fullname // update user full name
			config.WriteConfig()                              // write config file
			return
		}
	}

	// if user fullname is not exist than create new username
	config.AppConfig.FullNames = append(config.AppConfig.FullNames, config.FullName{Username: username, FullName: fullname})
	config.WriteConfig() // write to config file

}
