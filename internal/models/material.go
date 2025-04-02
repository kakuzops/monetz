package models

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Name string `gorm:"size:255;not null;uniqueIndex" json:"name" validate:"required,min=3,max=50,alphaunicode"`
}
