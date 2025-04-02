package server

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bnock/nockchat-api-go/internal/handlers"
	"github.com/bnock/nockchat-api-go/internal/repositories"
	"github.com/bnock/nockchat-api-go/internal/services"
	"github.com/labstack/echo/v4"
)

type Option func(*Server)

type Server struct {
	DB   *sql.DB
	Echo *echo.Echo
}

func NewServer(options ...Option) *Server {
	server := &Server{
		DB:   nil,
		Echo: echo.New(),
	}

	for _, option := range options {
		option(server)
	}

	r := repositories.NewRepositories(server.DB)
	s := services.NewServices(r)
	h := handlers.NewHandlers(s)
	h.RegisterRoutes(server.Echo)

	return server
}

func WithDB(db *sql.DB) Option {
	return func(server *Server) {
		server.DB = db
	}
}

func (s Server) Run() {
	s.Echo.Logger.Fatal(s.Echo.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
