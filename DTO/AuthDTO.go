package dto

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Register struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
