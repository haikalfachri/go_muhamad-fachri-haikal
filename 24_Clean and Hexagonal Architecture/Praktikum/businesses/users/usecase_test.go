package users_test

import (
	"belajar-go-echo/app/middlewares"
	"belajar-go-echo/businesses/users"
	"context"
	_userMock "belajar-go-echo/businesses/users/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository _userMock.Repository
	userUsecase    users.Usecase

	userDomain users.Domain
	ctx        context.Context
)

func TestMain(m *testing.M) {
	userUsecase = users.NewUserUseCase(&userRepository, &middlewares.JWTConfig{})
	userDomain = users.Domain{
		Email:    "dummy@example.com",
		Password: "dummyPassword",
	}

	ctx = context.TODO()

	m.Run()
}

func TestCreateUser(t *testing.T) {
	t.Run("CreateUser | Valid", func(t *testing.T) {
		userRepository.On("CreateUser", ctx, &userDomain).Return(userDomain, nil).Once()

		result, err := userUsecase.CreateUser(ctx, &userDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("CreateUser |  Invalid", func(t *testing.T) {
		userRepository.On("CreateUser", ctx, &users.Domain{}).Return(users.Domain{}, errors.New("failed")).Once()

		result, err := userUsecase.CreateUser(ctx, &users.Domain{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("GetAllUsers | Valid", func(t *testing.T) {
		userRepository.On("GetAllUsers", ctx).Return([]users.Domain{userDomain}, nil).Once()

		result, err := userUsecase.GetAllUsers(ctx)

		assert.Equal(t, 1, len(result))
		assert.Nil(t, err)
	})

	t.Run("GetAllUsers |  Invalid", func(t *testing.T) {
		userRepository.On("GetAllUsers", ctx).Return([]users.Domain{}, nil).Once()

		result, err := userUsecase.GetAllUsers(ctx)

		assert.Equal(t, 0, len(result))
		assert.Nil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		userRepository.On("GetByEmail", ctx, &userDomain).Return(userDomain, nil).Once()

		result, err := userUsecase.Login(ctx, &userDomain)

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Login |  Invalid", func(t *testing.T) {
		userRepository.On("GetByEmail", ctx, &users.Domain{}).Return(users.Domain{}, errors.New("failed")).Once()

		result, err := userUsecase.Login(ctx, &users.Domain{})

		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})
}