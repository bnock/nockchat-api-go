package server

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

type routeRegistrar interface {
	RegisterRoutes(e *echo.Echo)
}

type Option func(*Server)

type Server struct {
	Echo *echo.Echo
}

func (s Server) Run() {
	s.Echo.Logger.Fatal(s.Echo.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}

func NewServer(options ...Option) *Server {
	server := &Server{
		Echo: echo.New(),
	}

	for _, option := range options {
		option(server)
	}

	return server
}

func WithRoutes(registrar routeRegistrar) Option {
	return func(server *Server) {
		registrar.RegisterRoutes(server.Echo)
	}
}
