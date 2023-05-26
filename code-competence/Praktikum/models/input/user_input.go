package input

import "github.com/go-playground/validator/v10"

type UserInput struct {
	Name     	string `json:"name" validate:"required"`
	Email    	string `json:"email" gorm:"unique" validate:"required,email"`
	Password 	string `json:"password" validate:"required"`
}

func (u *UserInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}