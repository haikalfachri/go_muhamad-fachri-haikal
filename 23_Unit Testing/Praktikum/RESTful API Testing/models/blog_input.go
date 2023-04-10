package models

type BlogInput struct {
	Title    	string 	`json:"title"`
	Content   	string 	`json:"content"`
	UserID    	uint   	`json:"userId"`
}