package models

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"json"`
	Password string `json:"password"`
}
