package services

import (
	"monetz/internal/models"
	"monetz/internal/repositories"
)

type SellerService struct {
	sellerRepo *repositories.SellerRepository
}

func NewSellerService(sellerRepo *repositories.SellerRepository) *SellerService {
	return &SellerService{sellerRepo: sellerRepo}
}

func (s *SellerService) CreateSeller(seller *models.Seller) error {
	return s.sellerRepo.Create(seller)
}

func (s *SellerService) ListSellers(userID string) ([]models.Seller, error) {
	return s.sellerRepo.List(userID)
}

func (s *SellerService) DeleteSeller(sellerID string, userID string) error {
	return s.sellerRepo.Delete(sellerID, userID)
}

func (s *SellerService) UpdateSeller(sellerID string, userID string, seller *models.Seller) error {
	return s.sellerRepo.Update(sellerID, userID, seller)
}
