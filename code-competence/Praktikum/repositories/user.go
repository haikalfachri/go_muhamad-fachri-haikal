package repositories

import (
	"code_competence/database"
	"code_competence/models"
	"code_competence/models/input"
	
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImp struct {
}

func InitUserRepository() UserRepository {
	return &UserRepositoryImp{}
}

func (ur *UserRepositoryImp) Register(userInput input.UserInput) (models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost) 
	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name : userInput.Name,
		Email: userInput.Email,
		Password : string(hashedPass),
	}

	if err := database.ConnectDB().Create(&user).Error; err != nil {
		return models.User{}, err
	}

	if err := database.ConnectDB().Last(&user).Error; err != nil {
		return models.User{}, err
	}

    return user, nil
}

func (ur *UserRepositoryImp) Login(userInput input.UserInput) (models.User, error) {
	var user models.User

	if err := database.ConnectDB().First(&user, "email = ?", userInput.Email).Error; err != nil {
		return models.User{}, err
	}

	if err :=  bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		return models.User{}, err
	}

	return user, nil
}


