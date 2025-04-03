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

func (ch *ChannelHandler) GetChannels(c echo.Context) error {
	user, err := ch.securityService.GetAuthedUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}

	channels, err := ch.channelService.GetChannelsByUser(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unable to retrieve channels")
	}

	return c.JSON(http.StatusOK, channels)
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

func (ch *ChannelHandler) CreateChannel(c echo.Context) error {
	user, err := ch.securityService.GetAuthedUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthenticated")
	}

	type in struct {
		Name      string   `json:"name"`
		MemberIDs []string `json:"member_ids"`
	}

	var req in
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "unable to parse request body")
	}

	channel, err := ch.channelService.CreateChannel(user, req.Name, req.MemberIDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to create channel")
	}

	return c.JSON(http.StatusCreated, channel)
}

func (ch *ChannelHandler) CreateMessage(c echo.Context) error {
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

	type in struct {
		Content string `json:"content"`
	}

	var req in
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "unable to parse request body")
	}

	message, err := ch.messageService.CreateMessage(user, channel, req.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "unable to create message")
	}

	return c.JSON(http.StatusCreated, message)
}

func (ch *ChannelHandler) isAllowedToViewChannel(user *models.User, channel *models.Channel) (bool, error) {
	isMember := slices.ContainsFunc(channel.Members, func(member *models.User) bool {
		return member.ID == user.ID
	})

	if user.ID != channel.OwnerID && !isMember {
		return false, errors.New("unauthorized to view this channel")
	}

	return true, nil
}
