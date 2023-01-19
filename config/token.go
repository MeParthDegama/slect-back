package config


func AddToken(t string, u string) {

	// add token
	appConfig.Tokens = append(appConfig.Tokens, Token{t, u})
	WriteConfig()

}
