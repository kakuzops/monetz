package service

import "monetz/src/app/user/repositories"

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(req *CreateUserRequest) error {
	userRepo := repositories.NewUserRepository()

	user := &repositories.CreateUserInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
