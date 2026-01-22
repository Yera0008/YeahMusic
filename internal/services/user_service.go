package services

import "AP3/internal/models"

type UserService struct{}

func (s *UserService) Register(user models.User) error {
	return nil
}

func (s *UserService) Login(email, password string) (models.User, error) {
	return models.User{}, nil
}
