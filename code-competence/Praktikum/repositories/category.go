package repositories

import (
	"code_competence/database"
	"code_competence/models"
	"code_competence/models/input"
)

type CategoryRepositoryImp struct {
}

func InitCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImp{}
}

func (ur *CategoryRepositoryImp) Create(categoryInput input.CategoryInput) (models.Category, error) {
	var category models.Category = models.Category{
		Name: categoryInput.Name,
	}

	if err := database.ConnectDB().Create(&category).Error; err != nil {
		return models.Category{}, err
	}

	if err := database.ConnectDB().Last(&category).Error; err != nil {
		return models.Category{}, err
	}

    return category, nil
}

func (ur *CategoryRepositoryImp) GetAll() ([]models.Category, error) {
	var categorys []models.Category

	if err := database.ConnectDB().Find(&categorys).Error; err != nil {
		return categorys, err
	}
	return categorys, nil
}

func (ur *CategoryRepositoryImp) GetById(id string) (models.Category, error) {
	var category models.Category

	if err := database.ConnectDB().First(&category, id).Error; err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (ur *CategoryRepositoryImp) Update(categoryInput input.CategoryInput, id string) (models.Category, error) {
	category, err := ur.GetById(id)

	if err != nil {
		return models.Category{}, err
	}

	category.Name = categoryInput.Name

	if err := database.ConnectDB().Save(&category).Error; err != nil {
		return models.Category{}, err
	}

    return category, nil
}

func (ur *CategoryRepositoryImp) Delete(id string) error {
	category, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.ConnectDB().Delete(&category).Error; err != nil {
		return err
	}

    return nil
}


