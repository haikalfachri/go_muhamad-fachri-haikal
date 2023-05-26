package repositories

import (
	"code_competence/models"
	"code_competence/models/input"
)

type UserRepository interface {
	Register(userInput input.UserInput) (models.User, error)
	Login(userInput input.UserInput) (models.User, error)
}

type ItemRepository interface {
	Create(itemInput input.ItemInput) (models.Item, error)
	GetAll() ([]models.Item, error)
	GetById(id string) (models.Item, error)
	GetItemsByCategoryId(id string) ([]models.Item, error)
	GetItemsByName(name string) ([]models.Item, error)
	Update(itemInput input.ItemInput, id string) (models.Item, error)
	Delete(id string) error
}

type CategoryRepository interface {
	Create(categoryInput input.CategoryInput) (models.Category, error)
	GetAll() ([]models.Category, error)
	GetById(id string) (models.Category, error)
	Update(categoryInput input.CategoryInput, id string) (models.Category, error)
	Delete(id string) error
}