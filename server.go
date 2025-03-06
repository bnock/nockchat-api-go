package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load .env
	if e := godotenv.Load(); e != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a server
	e := echo.New()

	// Define routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []string{"Hello, World!"})
	})

	// Run the server
	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
