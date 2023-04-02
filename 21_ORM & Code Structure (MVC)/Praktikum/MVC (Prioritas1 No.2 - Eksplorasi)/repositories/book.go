package repositories

import (
	"mvc/database"
	"mvc/models"
	"errors"
)

type BookRepositoryImp struct {
}

func InitBookRepository() BookRepository {
	return &BookRepositoryImp{}
}

func (br *BookRepositoryImp) GetAll() ([]models.Book, error) {
	var books []models.Book

	if err := database.DB.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (br *BookRepositoryImp) GetById(id string) (models.Book, error) {
	var buku models.Book

	if err := database.DB.First(&buku, id).Error; err != nil {
		return models.Book{}, err
	}
	return buku, nil
}

func (br *BookRepositoryImp) Create(bukuInput models.BookInput) (models.Book, error) {
    if bukuInput.Title == "" || bukuInput.Writer == "" || bukuInput.Publisher == ""{
		return models.Book{}, errors.New("All field must not be empty!")
	}

	var buku models.Book = models.Book{
		Title : bukuInput.Title,
		Writer: bukuInput.Writer,
		Publisher : bukuInput.Publisher,
	}

	if err := database.DB.Create(&buku).Error; err != nil {
		return models.Book{}, err
	}

	if err := database.DB.Last(&buku).Error; err != nil {
		return models.Book{}, err
	}
    return buku, nil
}

func (br *BookRepositoryImp) Update(bukuInput models.BookInput, id string) (models.Book, error) {
	if bukuInput.Title == "" || bukuInput.Writer == "" || bukuInput.Publisher == ""{
		return models.Book{}, errors.New("All field must not be empty!")
	}

	buku, err := br.GetById(id)

	if err != nil {
		return models.Book{}, err
	}
	buku.Title = bukuInput.Title
	buku.Writer = bukuInput.Writer
	buku.Publisher = bukuInput.Publisher

	if err := database.DB.Save(&buku).Error; err != nil {
		return models.Book{}, err
	}
    return buku, nil
}

func (br *BookRepositoryImp) Delete(id string) error {
	buku, err := br.GetById(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&buku).Error; err != nil {
		return err
	}
    return nil
	
}


