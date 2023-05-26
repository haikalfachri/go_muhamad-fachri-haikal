package services

import (
	"code_competence/repositories"
	"code_competence/models"
	"code_competence/models/input"
	"code_competence/middlewares"
)

type ItemService struct {
	repository repositories.ItemRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitItemService(jwtAuth *middlewares.JWTConfig) ItemService {
	return ItemService{
		repository: &repositories.ItemRepositoryImp{},
		jwtAuth: jwtAuth,
	}
}

func (us *ItemService) Create(itemInput input.ItemInput) (models.Item, error){
	return us.repository.Create(itemInput)
}

func (us *ItemService) GetAll() ([]models.Item, error){
	return us.repository.GetAll()
}

func (us *ItemService) GetById(id string) (models.Item, error){
	return us.repository.GetById(id)
}

func (us *ItemService) GetItemsByCategoryId(id string) ([]models.Item, error){
	return us.repository.GetItemsByCategoryId(id)
}

func (us *ItemService) GetItemsByName(name string) ([]models.Item, error){
	return us.repository.GetItemsByName(name)
}

func (us *ItemService) Update(itemInput input.ItemInput, id string) (models.Item, error){
	return us.repository.Update(itemInput, id)
}

func (us *ItemService) Delete(id string) (error){
	return us.repository.Delete(id)
}