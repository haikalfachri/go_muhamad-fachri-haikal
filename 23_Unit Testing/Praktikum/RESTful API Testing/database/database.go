package database

import (
	"errors"
	"fmt"
	"log"
	"restful_api_testing/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}
  
func InitDB() {
	config := Config{
	  DB_Username: "root",
	  DB_Password: "",
	  DB_Port:     "3306",
	  DB_Host:     "localhost",
	  DB_Name:     "crud_go",
	}
  
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	  config.DB_Username,
	  config.DB_Password,
	  config.DB_Host,
	  config.DB_Port,
	  config.DB_Name,
	)
  
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
	  panic(err)
	}
}

func InitialMigration() {
	err := DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Blog{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}

func SeedUser() (models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte("dummy"), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Name: "dummy",
		Email: "dummy",
		Password: string(hashedPass),
	}

	result := DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	if err := result.Last(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func SeedBook() (models.Book, error) {
	var book models.Book = models.Book{
		Title:  "dummy",
		Writer: "dummy",
		Publisher: "dummy",
	}

	result := DB.Create(&book)

	if err := result.Error; err != nil {
		return models.Book{}, err
	}

	if err := result.Last(&book).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func SeedBlog() (models.Blog, error) {
	user, err := SeedUser()

	if err != nil {
		return models.Blog{}, err
	}

	var blog models.Blog = models.Blog{
		Title: "dummy",
		Content: "dummy",
		UserID: user.ID,
	}

	result := DB.Create(&blog)

	if err := result.Error; err != nil {
		return models.Blog{}, err
	}

	if err := result.Last(&blog).Error; err != nil {
		return models.Blog{}, err
	}

	return blog, nil
}

func CleanSeeders() error {
	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	userErr := DB.Exec("DELETE FROM users").Error
	bookErr := DB.Exec("DELETE FROM books").Error
	blogErr := DB.Exec("DELETE FROM blogs").Error

	if isFailed := userErr != nil || bookErr != nil || blogErr != nil; isFailed{
		return errors.New("cleaning failed")
	}

	log.Println("seeders are cleaned up successfully")

	return nil
}