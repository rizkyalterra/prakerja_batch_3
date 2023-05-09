package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Address  string `json:"address"`
	Password string `jsin:"string"`
}
