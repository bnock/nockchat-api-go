package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ClaimsFromCtx(c echo.Context) (jwt.MapClaims, error) {
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

func UserIDFromCtx(c echo.Context) (string, error) {
	claims, err := ClaimsFromCtx(c)
	if err != nil {
		return "", err
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return "", err
	}

	return sub, nil
}
