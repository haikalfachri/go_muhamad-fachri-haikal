package users

import (
	"context"
	"belajar-go-docker/app/middlewares"
)

type userUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.JWTConfig
}

func NewUserUseCase(repository Repository, jwtAuth *middlewares.JWTConfig) Usecase {
	return &userUsecase{
		userRepository: repository,
		jwtAuth: jwtAuth,
	}
}

func (usecase *userUsecase) CreateUser(ctx context.Context, userDomain *Domain) (Domain, error) {
	return usecase.userRepository.CreateUser(ctx, userDomain)
}

func (usecase *userUsecase) GetAllUsers(ctx context.Context) ([]Domain, error) {
	return usecase.userRepository.GetAllUsers(ctx)
}

func (usecase *userUsecase) Login(ctx context.Context, userDomain *Domain) (string, error) {
	user, err := usecase.userRepository.GetByEmail(ctx, userDomain)

	if err != nil {
		return "", err
	}

	token, err := usecase.jwtAuth.GenerateToken(int(user.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
