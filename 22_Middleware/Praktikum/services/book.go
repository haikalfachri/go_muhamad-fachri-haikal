package services

import (
	"mvc_middleware/repositories"
	"mvc_middleware/models"
)

type BookService struct {
	repository repositories.BookRepository
}

func InitBookService() BookService {
	return BookService{
		repository: &repositories.BookRepositoryImp{},
	}
}

func (bs *BookService) GetAll() ([]models.Book, error){
	return bs.repository.GetAll()
}

func (bs *BookService) GetById(id string) (models.Book, error){
	return bs.repository.GetById(id)
}

func (bs *BookService) Create(bookInput models.BookInput) (models.Book, error){
	return bs.repository.Create(bookInput)
}

func (bs *BookService) Update(bookInput models.BookInput, id string) (models.Book, error){
	return bs.repository.Update(bookInput, id)
}

func (bs *BookService) Delete(id string) (error){
	return bs.repository.Delete(id)
}



