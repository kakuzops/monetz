package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UserID string `gorm:"size:255;not null" json:"user_id"`
	Name   string `gorm:"size:255;not null" json:"name" validate:"required,min=3,max=50"`
}
