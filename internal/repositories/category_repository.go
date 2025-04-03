package repositories

import (
	"monetz/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) ListByUserID(userID string) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("user_id = ?", userID).Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetByID(id uint, userID string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&category).Error
	return &category, err
}

func (r *CategoryRepository) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Category{}).Error
}
