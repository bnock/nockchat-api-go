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
	userHandler     UserHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		channelHandler:  *NewChannelHandler(services),
		securityHandler: *NewSecurityHandler(services),
		userHandler:     *NewUserHandler(services),
	}
}

func (sh *Handlers) RegisterRoutes(e *echo.Echo) {
	e.POST("/register", sh.userHandler.Register).Name = "register"
	e.POST("/login", sh.securityHandler.Login).Name = "login"

	// Protected
	protected := e.Group("")
	protected.Use(echojwt.JWT([]byte(os.Getenv("APP_SECRET"))))

	protected.GET("/channels/:channel/messages", sh.channelHandler.GetChannelMessages).Name = "channelMessages"
	protected.GET("/channels/:channel", sh.channelHandler.GetChannel).Name = "channel"
	protected.GET("/channels", sh.channelHandler.GetChannels).Name = "channels"
	protected.POST("/channels", sh.channelHandler.CreateChannel).Name = "createChannel"
}
