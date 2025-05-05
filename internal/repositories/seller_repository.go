package repositories

import (
	"monetz/internal/models"

	"gorm.io/gorm"
)

type SellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{db: db}
}

func (r *SellerRepository) Create(seller *models.Seller) error {
	return r.db.Create(seller).Error
}

func (r *SellerRepository) List(userID string) ([]models.Seller, error) {
	var sellers []models.Seller
	err := r.db.Where("user_id = ?", userID).Find(&sellers).Error
	return sellers, err
}

func (r *SellerRepository) Delete(sellerID string, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", sellerID, userID).Delete(&models.Seller{}).Error
}

func (r *SellerRepository) Update(sellerID string, userID string, seller *models.Seller) error {
	return r.db.Model(&models.Seller{}).Where("id = ? AND user_id = ?", sellerID, userID).Updates(seller).Error
}
