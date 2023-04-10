package models

type UserInput struct {
	Name     	string `json:"name"`
	Email    	string `json:"email" gorm:"unique"`
	Password 	string `json:"password"`
}