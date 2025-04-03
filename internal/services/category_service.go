package services

import (
	"monetz/internal/models"
	"monetz/internal/repositories"
)

type CategoryService struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.categoryRepo.Create(category)
}

func (s *CategoryService) ListCategories(userID string) ([]models.Category, error) {
	return s.categoryRepo.ListByUserID(userID)
}

func (s *CategoryService) GetCategory(id uint, userID string) (*models.Category, error) {
	return s.categoryRepo.GetByID(id, userID)
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	return s.categoryRepo.Update(category)
}

func (s *CategoryService) DeleteCategory(id uint, userID string) error {
	return s.categoryRepo.Delete(id, userID)
}
