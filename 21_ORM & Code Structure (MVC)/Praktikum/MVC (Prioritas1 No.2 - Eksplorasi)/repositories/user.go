package repositories

import (
	"mvc/database"
	"mvc/models"
	"errors"
	
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImp struct {
}

func InitUserRepository() UserRepository {
	return &UserRepositoryImp{}
}

func (ur *UserRepositoryImp) GetAll() ([]models.User, error) {
	var users []models.User

	if err := database.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (ur *UserRepositoryImp) GetById(id string) (models.User, error) {
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepositoryImp) Create(userInput models.UserInput) (models.User, error) {
    if userInput.Name == "" || userInput.Email == "" || userInput.Password == ""{
		return models.User{}, errors.New("All field must not be empty!")
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost) 
	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name : userInput.Name,
		Email: userInput.Email,
		Password : string(hashedPass),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	if err := database.DB.Last(&user).Error; err != nil {
		return models.User{}, err
	}
    return user, nil
}

func (ur *UserRepositoryImp) Update(userInput models.UserInput, id string) (models.User, error) {
	if userInput.Name == "" || userInput.Email == "" || userInput.Password == ""{
		return models.User{}, errors.New("All field must not be empty!")
	}

	user, err := ur.GetById(id)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost) 
	if err != nil {
		return models.User{}, err
	}
	user.Name = userInput.Name
	user.Email = userInput.Email
	user.Password = string(hashedPass)

	if err := database.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}
    return user, nil
}

func (ur *UserRepositoryImp) Delete(id string) error {
	user, err := ur.GetById(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}
    return nil
	
}


