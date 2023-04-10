package users

import (
	"belajar-go-docker/businesses/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		Email:     rec.Email,
		Password:  rec.Password,
	}
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		Email:     domain.Email,
		Password:  domain.Password,
	}
}