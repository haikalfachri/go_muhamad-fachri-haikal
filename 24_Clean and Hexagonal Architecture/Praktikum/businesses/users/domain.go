package users

import (
	"context"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	Email    string 
	Password string
}

type Usecase interface {
	CreateUser(ctx context.Context, userDomain *Domain) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
	Login(ctx context.Context, userDomain *Domain) (string, error)
}

type Repository interface {
	CreateUser(ctx context.Context, userDomain *Domain) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
	GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
}