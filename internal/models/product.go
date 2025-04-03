package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	Price       float64    `json:"price" validate:"required"`
	UserID      string     `json:"user_id"`
	Categories  []Category `gorm:"many2many:product_categories;" json:"categories"`
}
