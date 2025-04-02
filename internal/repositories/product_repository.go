package repositories

import (
	"monetz/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) ListByUserID(userID string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("user_id = ?", userID).Find(&products).Error
	return products, err
}
