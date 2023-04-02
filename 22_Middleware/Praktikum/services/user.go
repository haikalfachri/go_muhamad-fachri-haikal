package services

import (
	"mvc_middleware/repositories"
	"mvc_middleware/models"
	"mvc_middleware/middlewares"
)

type UserService struct {
	repository repositories.UserRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitUserService(jwtAuth *middlewares.JWTConfig) UserService {
	return UserService{
		repository: &repositories.UserRepositoryImp{},
		jwtAuth: jwtAuth,
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

func (us *UserService) Login(userInput models.UserInput) (string, error) {
	user, err := us.repository.GetByEmail(userInput)
	if err != nil {
		return "", err
	}

	token, err := us.jwtAuth.GenerateToken(int(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}


