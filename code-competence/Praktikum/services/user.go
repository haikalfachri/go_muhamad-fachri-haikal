package services

import (
	"code_competence/repositories"
	"code_competence/models"
	"code_competence/models/input"
	"code_competence/middlewares"
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

func (us *UserService) Register(userInput input.UserInput) (models.User, error){
	return us.repository.Register(userInput)
}

func (us *UserService) Login(userInput input.UserInput) (string, error) {
	user, err := us.repository.Login(userInput)
	if err != nil {
		return "", err
	}

	token, err := us.jwtAuth.GenerateToken(int(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}