package services

import (
	"monetz/internal/models"
	"monetz/internal/repositories"
)

type ColorService struct {
	colorRepo *repositories.ColorRepository
}

func NewColorService(colorRepo *repositories.ColorRepository) *ColorService {
	return &ColorService{colorRepo: colorRepo}
}

func (s *ColorService) CreateColor(color *models.Color) error {
	return s.colorRepo.Create(color)
}

func (s *ColorService) ListColors() ([]models.Color, error) {
	return s.colorRepo.List()
}
