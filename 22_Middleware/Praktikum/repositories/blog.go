package repositories

import (
	"errors"
	"mvc_middleware/database"
	"mvc_middleware/models"

	"golang.org/x/crypto/bcrypt"
)

type BlogRepositoryImp struct {
}

func InitBlogRepository() BlogRepository {
	return &BlogRepositoryImp{}
}

func (br *BlogRepositoryImp) GetAll() ([]models.Blog, error) {
	var blog []models.Blog

	if err := database.DB.Preload("User").Find(&blog).Error; err != nil {
		return blog, err
	}
	return blog, nil
}

func (br *BlogRepositoryImp) GetById(id string) (models.Blog, error) {
	var blog models.Blog

	if err := database.DB.Preload("User").First(&blog, id).Error; err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func (br *BlogRepositoryImp) Create(blogInput models.BlogInput) (models.Blog, error) {
	if blogInput.Title == "" || blogInput.Content == "" {
		return models.Blog{}, errors.New("All field must not be empty")
	}

	var blog models.Blog = models.Blog{
		Title:   blogInput.Title,
		Content: blogInput.Content,
		User:    blogInput.User,
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(blogInput.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.Blog{}, err
	}
	blog.User.Password = string(hashedPass)

	if err := database.DB.Create(&blog).Error; err != nil {
		return models.Blog{}, err
	}

	if err := database.DB.Last(&blog).Error; err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func (br *BlogRepositoryImp) Update(blogInput models.BlogInput, id string) (models.Blog, error) {
	if blogInput.Title == "" || blogInput.Content == "" {
		return models.Blog{}, errors.New("All field must not be empty")
	}

	blog, err := br.GetById(id)

	if err != nil {
		return models.Blog{}, err
	}
	blog.Title = blogInput.Title
	blog.Content = blogInput.Content

	if err := database.DB.Save(&blog).Error; err != nil {
		return models.Blog{}, err
	}
	return blog, nil
}

func (br *BlogRepositoryImp) Delete(id string) error {
	blog, err := br.GetById(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&blog).Error; err != nil {
		return err
	}
	return nil
}
