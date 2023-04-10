package repositories

import (
	"restful_api_testing/database"
	"restful_api_testing/models"
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
	var book models.Book

	if err := database.DB.First(&book, id).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (br *BookRepositoryImp) Create(bookInput models.BookInput) (models.Book, error) {
    if bookInput.Title == "" || bookInput.Writer == "" || bookInput.Publisher == ""{
		return models.Book{}, errors.New("All field must not be empty!")
	}

	var book models.Book = models.Book{
		Title : bookInput.Title,
		Writer: bookInput.Writer,
		Publisher : bookInput.Publisher,
	}

	if err := database.DB.Create(&book).Error; err != nil {
		return models.Book{}, err
	}

	if err := database.DB.Last(&book).Error; err != nil {
		return models.Book{}, err
	}
    return book, nil
}

func (br *BookRepositoryImp) Update(bookInput models.BookInput, id string) (models.Book, error) {
	if bookInput.Title == "" || bookInput.Writer == "" || bookInput.Publisher == ""{
		return models.Book{}, errors.New("All field must not be empty!")
	}

	book, err := br.GetById(id)

	if err != nil {
		return models.Book{}, err
	}
	book.Title = bookInput.Title
	book.Writer = bookInput.Writer
	book.Publisher = bookInput.Publisher

	if err := database.DB.Save(&book).Error; err != nil {
		return models.Book{}, err
	}
    return book, nil
}

func (br *BookRepositoryImp) Delete(id string) error {
	book, err := br.GetById(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		return err
	}
    return nil
	
}


