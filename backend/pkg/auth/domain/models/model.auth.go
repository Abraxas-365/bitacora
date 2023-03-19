package models

type RegisterWithPlatform struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
