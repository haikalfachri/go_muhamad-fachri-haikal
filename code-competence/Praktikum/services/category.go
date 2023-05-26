package services

import (
	"code_competence/repositories"
	"code_competence/models"
	"code_competence/models/input"
	"code_competence/middlewares"
)

type CategoryService struct {
	repository repositories.CategoryRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitCategoryService(jwtAuth *middlewares.JWTConfig) CategoryService {
	return CategoryService{
		repository: &repositories.CategoryRepositoryImp{},
		jwtAuth: jwtAuth,
	}
}

func (us *CategoryService) Create(categoryInput input.CategoryInput) (models.Category, error){
	return us.repository.Create(categoryInput)
}

func (us *CategoryService) GetAll() ([]models.Category, error){
	return us.repository.GetAll()
}

func (us *CategoryService) GetById(id string) (models.Category, error){
	return us.repository.GetById(id)
}

func (us *CategoryService) Update(categoryInput input.CategoryInput, id string) (models.Category, error){
	return us.repository.Update(categoryInput, id)
}

func (us *CategoryService) Delete(id string) (error){
	return us.repository.Delete(id)
}