package services

import (
	"mvc/repositories"
	"mvc/models"
)

type UserService struct {
	repository repositories.UserRepository
}

func InitUserService() UserService {
	return UserService{
		repository: &repositories.UserRepositoryImp{},
	}
}

func (us *UserService) GetAll() ([]models.User, error){
	return us.repository.GetAll()
}

func (us *UserService) GetById(id string) (models.User, error){
	return us.repository.GetById(id)
}

func (us *UserService) Create(userInput models.UserInput) (models.User, error){
	return us.repository.Create(userInput)
}

func (us *UserService) Update(userInput models.UserInput, id string) (models.User, error){
	return us.repository.Update(userInput, id)
}

func (us *UserService) Delete(id string) (error){
	return us.repository.Delete(id)
}



