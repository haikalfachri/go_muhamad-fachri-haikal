package users

import (
	"context"
	"belajar-go-echo/businesses/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	if err != nil {
		return users.Domain{}, err
	}

	record := FromDomain(userDomain)

	record.Password = string(password)

	result := ur.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return users.Domain{}, err
	}

	err = result.Last(&record).Error

	if err != nil {
		return users.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ur *userRepository) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
	var records []User

	if err := ur.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	users := []users.Domain{}

	for _, user := range records {
		users = append(users, user.ToDomain())
	}

	return users, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	var user User

	if err := ur.conn.WithContext(ctx).First(&user, "email = ?", userDomain.Email).Error; err != nil {
		return users.Domain{}, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

