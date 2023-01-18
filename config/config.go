package config

type Config struct {
	Tokens []Token `json:"tokens"`
}

type Token struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

