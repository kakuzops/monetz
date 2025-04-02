package services

import (
	"monetz/internal/models"
	"monetz/internal/repositories"
)

type MaterialService struct {
	materialRepo *repositories.MaterialRepository
}

func NewMaterialService(materialRepo *repositories.MaterialRepository) *MaterialService {
	return &MaterialService{materialRepo: materialRepo}
}

func (s *MaterialService) CreateMaterial(material *models.Material) error {
	return s.materialRepo.Create(material)
}

func (s *MaterialService) ListMaterials() ([]models.Material, error) {
	return s.materialRepo.List()
}
