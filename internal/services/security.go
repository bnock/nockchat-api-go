package services

import (
	"errors"

	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type SecurityService struct {
	userRepository *repositories.UserRepository
}

func (ss *SecurityService) GetAuthedUser(c echo.Context) (*models.User, error) {
	userId, err := ss.UserIDFromCtx(c)
	if err != nil {
		return nil, err
	}

	user, err := ss.userRepository.UserById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ss *SecurityService) ClaimsFromCtx(c echo.Context) (jwt.MapClaims, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("JWT claims missing or invalid")
	}

	return claims, nil
}

func (ss *SecurityService) UserIDFromCtx(c echo.Context) (string, error) {
	claims, err := ss.ClaimsFromCtx(c)
	if err != nil {
		return "", err
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return "", err
	}

	return sub, nil
}
