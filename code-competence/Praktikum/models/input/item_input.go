package input

import (
	"code_competence/models"
	"github.com/go-playground/validator/v10"
)

type ItemInput struct {
	Name     	string 	 		`json:"name" validate:"required"`
	Description string   		`json:"description" form:"description"`
	Stock       int      		`json:"stock" form:"stock"`
	Price       int      		`json:"price" form:"price"`
	CategoryID	uint			`json:"category_id" validate:"required"`
	Category    models.Category	`json:"-" gorm:"foreignKey:CategoryID"`
}

func (u *ItemInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	return err
}