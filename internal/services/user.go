package services

import (
	"time"

	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (us *UserService) CreateUser(
	email string,
	firstName string,
	lastName string,
	password string,
) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	u := &models.User{
		ID:        uuid.NewString(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  string(hashedPassword),
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := us.userRepository.CreateUser(u); err != nil {
		return nil, err
	}

	return u, nil
}
