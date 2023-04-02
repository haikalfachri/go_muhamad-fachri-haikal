package services

import (
	"mvc/repositories"
	"mvc/models"
)

type BookService struct {
	repository repositories.BookRepository
}

func InitBookService() BookService {
	return BookService{
		repository: &repositories.BookRepositoryImp{},
	}
}

func (sb *BookService) GetAll() ([]models.Book, error){
	return sb.repository.GetAll()
}

func (sb *BookService) GetById(id string) (models.Book, error){
	return sb.repository.GetById(id)
}

func (sb *BookService) Create(bukuInput models.BookInput) (models.Book, error){
	return sb.repository.Create(bukuInput)
}

func (sb *BookService) Update(bukuInput models.BookInput, id string) (models.Book, error){
	return sb.repository.Update(bukuInput, id)
}

func (sb *BookService) Delete(id string) (error){
	return sb.repository.Delete(id)
}



