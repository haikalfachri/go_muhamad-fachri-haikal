package repositories

import (
	"mvc/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id string) (models.User, error)
	Create(userInput models.UserInput) (models.User, error)
	Update(userInput models.UserInput, id string) (models.User, error)
	Delete(id string) error
}

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id string) (models.Book, error)
	Create(bookInput models.BookInput) (models.Book, error)
	Update(bookInput models.BookInput, id string) (models.Book, error)
	Delete(id string) error
}