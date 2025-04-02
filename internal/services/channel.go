package services

import (
	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/repositories"
)

type ChannelService struct {
	channelRepository *repositories.ChannelRepository
}

func (cs *ChannelService) GetChannelByID(id string) (*models.Channel, error) {
	c, err := cs.channelRepository.ChannelById(id)
	if err != nil {
		return nil, err
	}

	members, err := cs.channelRepository.MembersByChannelID(id)
	if err != nil {
		return nil, err
	}
	c.Members = members

	return c, nil
}

func (cs *ChannelService) GetChannelsByUser(u *models.User) ([]*models.Channel, error) {
	channels, err := cs.channelRepository.ChannelsByUserID(u.ID)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
