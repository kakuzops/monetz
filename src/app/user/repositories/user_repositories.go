package repositories

import (
	"monetz/src/app/user/dto"
	"monetz/src/config/database"
	"monetz/src/config/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

type CreateUserInput struct {
	Name     string
	Email    string
	Password string
}

func (r *UserRepository) CreateUser(input *CreateUserInput) error {
	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	result := database.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserRepository) GetAllUsers() ([]dto.UserResponse, error) {
	var users []dto.UserResponse

	result := database.DB.Model(&models.User{}).Select("id", "name", "email").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
