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

func (r *ProductRepository) Delete(productID string) error {
	return r.db.Where("id = ?", productID).Delete(&models.Product{}).Error
}

func (r *ProductRepository) Update(productID string, name string, price string, stock int) error {
	return r.db.Model(&models.Product{}).Where("id = ?", productID).Updates(map[string]interface{}{"name": name, "price": price, "stock": stock}).Error
}
