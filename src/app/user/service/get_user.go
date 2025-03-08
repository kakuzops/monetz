package service

import (
	"monetz/src/app/user/dto"
	"monetz/src/app/user/repositories"
)

type GetUserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers() ([]dto.UserResponse, error) {
	userRepo := repositories.NewUserRepository()

	users, err := userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
