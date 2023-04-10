package models

type BookInput struct {
	Title    	string 	`json:"title"`
	Writer   	string 	`json:"writer"`
	Publisher	string 	`json:"publisher"`
}