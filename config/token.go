package config


func AddToken(t string, u string) {

	// add token
	AppConfig.Tokens = append(AppConfig.Tokens, Token{t, u})
	WriteConfig()

}
