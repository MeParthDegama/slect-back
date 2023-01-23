package config

type Config struct {
	Tokens    []Token    `json:"tokens"`
	FullNames []FullName `json:"fullnames"`
}

type Token struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

type FullName struct {
	Username string `json:"username"`
	FullName string `json:"fullname"`
}
