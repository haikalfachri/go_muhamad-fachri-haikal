package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name"`
	Item []Item `gorm:"foreignkey:CategoryID"`
}
