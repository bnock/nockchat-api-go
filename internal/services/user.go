package services

import (
	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func (us *UserService) GetUserByID(id string) (*models.User, error) {
	user, err := us.userRepository.UserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := us.userRepository.UserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
