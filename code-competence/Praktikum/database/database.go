package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"code_competence/models"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func ConnectDB() *gorm.DB {
	config := Config{
	  DB_Username: "root",
	  DB_Password: "root123",
	  DB_Port:     "3306",
	  DB_Host:     "localhost",
	  DB_Name:     "cc_db",
	}
  
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	  config.DB_Username,
	  config.DB_Password,
	  config.DB_Host,
	  config.DB_Port,
	  config.DB_Name,
	)
  
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
	  panic(err)
	}

	log.Printf("successfully connected to database\n")

	return db
}

func MigrateDB(db *gorm.DB) {

	err := db.AutoMigrate(&models.User{}, &models.Item{}, &models.Category{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}

	log.Printf("successfully database migration\n")
}
