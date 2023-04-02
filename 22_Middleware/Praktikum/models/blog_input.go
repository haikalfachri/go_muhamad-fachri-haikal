package models

type BlogInput struct {
	Title    	string 	`json:"title"`
	Content   	string 	`json:"content"`
	User    	User   	`json:"user"`
}