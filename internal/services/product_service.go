package services

import (
	"monetz/internal/models"
	"monetz/internal/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.productRepo.Create(product)
}

func (s *ProductService) ListProducts(userID string) ([]models.Product, error) {
	return s.productRepo.ListByUserID(userID)
}
