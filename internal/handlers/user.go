package handlers

import (
	"net/http"

	"github.com/bnock/nockchat-api-go/internal/services"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(services *services.Services) *UserHandler {
	return &UserHandler{userService: services.UserService}
}

func (uh *UserHandler) Register(c echo.Context) error {
	type in struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	var req in
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "unable to parse request body")
	}

	user, err := uh.userService.CreateUser(req.Email, req.FirstName, req.LastName, req.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
		//return c.String(http.StatusInternalServerError, "unable to create user")
	}

	return c.JSON(http.StatusCreated, user)
}
