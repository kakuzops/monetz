package models

import "gorm.io/gorm"

type Color struct {
	gorm.Model
	Name string `gorm:"size:255;not null;uniqueIndex" json:"name"`
}
