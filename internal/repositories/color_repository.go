package repositories

import (
	"monetz/internal/models"

	"gorm.io/gorm"
)

type ColorRepository struct {
	db *gorm.DB
}

func NewColorRepository(db *gorm.DB) *ColorRepository {
	return &ColorRepository{db: db}
}

func (r *ColorRepository) Create(color *models.Color) error {
	return r.db.Create(color).Error
}

func (r *ColorRepository) List() ([]models.Color, error) {
	var colors []models.Color
	err := r.db.Find(&colors).Error
	return colors, err
}
