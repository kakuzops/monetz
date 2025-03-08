package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;not null;unique"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
