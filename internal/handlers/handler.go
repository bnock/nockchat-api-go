package handlers

import (
	"os"

	"github.com/bnock/nockchat-api-go/internal/services"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	channelHandler  ChannelHandler
	securityHandler SecurityHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		channelHandler:  *NewChannelHandler(services),
		securityHandler: *NewSecurityHandler(services),
	}
}

func (sh Handlers) RegisterRoutes(e *echo.Echo) {
	e.POST("/login", sh.securityHandler.Login).Name = "login"

	// Protected
	protected := e.Group("")
	protected.Use(echojwt.JWT([]byte(os.Getenv("APP_SECRET"))))

	protected.GET("/channels/:channel/messages", sh.channelHandler.GetChannelMessages).Name = "channelMessages"
	protected.GET("/channels/:channel", sh.channelHandler.GetChannel).Name = "channel"
}
