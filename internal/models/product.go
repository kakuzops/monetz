package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID         string    `gorm:"size:255;not null" json:"user_id"`
	Name           string    `gorm:"size:255;not null" json:"name"`
	Description    string    `gorm:"size:1024" json:"description"`
	ProductionCost float64   `gorm:"type:decimal(10,2);not null" json:"production_cost"`
	SalePrice      float64   `gorm:"type:decimal(10,2);not null" json:"sale_price"`
	MaterialID     uint      `gorm:"not null" json:"material_id"`
	ColorID        uint      `gorm:"not null" json:"color_id"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
