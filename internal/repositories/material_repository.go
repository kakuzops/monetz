package repositories

import (
	"monetz/internal/models"

	"gorm.io/gorm"
)

type MaterialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *MaterialRepository {
	return &MaterialRepository{db: db}
}

func (r *MaterialRepository) Create(material *models.Material) error {
	return r.db.Create(material).Error
}

func (r *MaterialRepository) List() ([]models.Material, error) {
	var materials []models.Material
	err := r.db.Find(&materials).Error
	return materials, err
}
