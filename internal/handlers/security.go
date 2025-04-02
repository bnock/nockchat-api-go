package handlers

import (
	"net/http"
	"os"

	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type SecurityHandler struct {
	userService *services.UserService
}

func NewSecurityHandler(services *services.Services) *SecurityHandler {
	return &SecurityHandler{userService: services.UserService}
}

func (sh *SecurityHandler) Login(c echo.Context) error {
	type in struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req in
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "unable to parse request body")
	}

	user, err := sh.userService.GetUserByEmail(req.Email)
	if err != nil {
		return c.String(http.StatusUnauthorized, "unauthorized")
	}

	if !sh.isCorrectPassword(*user, req.Password) {
		return c.String(http.StatusUnauthorized, "unauthorized")
	}

	token, err := sh.jwtFromUser(*user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "unable to issue token")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (sh *SecurityHandler) isCorrectPassword(u models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (sh *SecurityHandler) jwtFromUser(u models.User) (string, error) {
	secret := os.Getenv("APP_SECRET")

	claims := &jwt.RegisteredClaims{
		Subject: u.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
