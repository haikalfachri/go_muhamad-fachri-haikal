package repositories

import (
	"mvc_middleware/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id string) (models.User, error)
	Create(userInput models.UserInput) (models.User, error)
	Update(userInput models.UserInput, id string) (models.User, error)
	Delete(id string) error
	GetByEmail(userInput models.UserInput) (models.User, error)
}

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id string) (models.Book, error)
	Create(bookInput models.BookInput) (models.Book, error)
	Update(bookInput models.BookInput, id string) (models.Book, error)
	Delete(id string) error
}

type BlogRepository interface {
	GetAll() ([]models.Blog, error)
	GetById(id string) (models.Blog, error)
	Create(blogInput models.BlogInput) (models.Blog, error)
	Update(blogInput models.BlogInput, id string) (models.Blog, error)
	Delete(id string) error
}