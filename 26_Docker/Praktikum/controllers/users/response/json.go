package response

import (
	"belajar-go-docker/businesses/users"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FromDomain(domain users.Domain) User {
	return User{
		Email    : domain.Email,
		Password : domain.Password,
	}
}