package handlers

import (
	"errors"
	"net/http"
	"slices"

	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/services"
	"github.com/labstack/echo/v4"
)

type ChannelHandler struct {
	channelService  *services.ChannelService
	messageService  *services.MessageService
	userService     *services.UserService
	securityService *services.SecurityService
}

func NewChannelHandler(services *services.Services) *ChannelHandler {
	return &ChannelHandler{
		channelService:  services.ChannelService,
		messageService:  services.MessageService,
		userService:     services.UserService,
		securityService: services.SecurityService,
	}
}

func (ch *ChannelHandler) GetChannel(c echo.Context) error {
	user, err := ch.securityService.GetAuthedUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}

	channel, err := ch.channelService.GetChannelByID(c.Param("channel"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "channel not found")
	}

	isAllowed, _ := ch.isAllowedToViewChannel(user, channel)
	if !isAllowed {
		return c.JSON(http.StatusForbidden, "unauthorized")
	}

	return c.JSON(http.StatusOK, channel)
}

func (ch *ChannelHandler) GetChannelMessages(c echo.Context) error {
	user, err := ch.securityService.GetAuthedUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}

	channel, err := ch.channelService.GetChannelByID(c.Param("channel"))
	if err != nil {
		return c.JSON(http.StatusNotFound, "channel not found")
	}

	isAllowed, _ := ch.isAllowedToViewChannel(user, channel)
	if !isAllowed {
		return c.JSON(http.StatusForbidden, "unauthorized")
	}

	messages, err := ch.messageService.GetMessagesByChannel(channel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to retrieve messages")
	}

	return c.JSON(http.StatusOK, messages)
}

func (ch *ChannelHandler) isAllowedToViewChannel(user *models.User, channel *models.Channel) (bool, error) {
	isMember := slices.ContainsFunc(channel.Members, func(member models.User) bool {
		return member.ID == user.ID
	})

	if user.ID != channel.OwnerID && !isMember {
		return false, errors.New("unauthorized to view this channel")
	}

	return true, nil
}
