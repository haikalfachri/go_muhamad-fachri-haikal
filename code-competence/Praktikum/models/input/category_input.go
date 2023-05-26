package input

import (
	"github.com/go-playground/validator/v10"
)

type CategoryInput struct {
	Name	string 	 `json:"name" validate:"required"`
}

func (u *CategoryInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}