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

func (s *ProductService) DeleteProduct(productID string) error {
	return s.productRepo.Delete(productID)
}

func (s *ProductService) UpdateProduct(productID string, name string, price string, stock int) error {
	return s.productRepo.Update(productID, name, price, stock)
}
