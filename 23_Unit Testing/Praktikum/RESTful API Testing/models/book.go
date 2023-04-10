package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model		
	Title    	string 	`json:"title"`
	Writer   	string 	`json:"writer"`
	Publisher	string 	`json:"publisher"`
}