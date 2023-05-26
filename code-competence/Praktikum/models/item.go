package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Stock       int      `json:"stock"`
	Price       int      `json:"price"`
	CategoryID	uint	 `json:"category_id"`
	Category    Category `json:"-" gorm:"foreignKey:CategoryID"`
}
