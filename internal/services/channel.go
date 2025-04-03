package services

import (
	"errors"
	"slices"
	"time"

	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/repositories"
	"github.com/google/uuid"
)

type ChannelService struct {
	channelRepository *repositories.ChannelRepository
	userRepository    *repositories.UserRepository
}

func (cs *ChannelService) CreateChannel(owner *models.User, name string, memberIDs []string) (*models.Channel, error) {
	if slices.Contains(memberIDs, owner.ID) {
		return nil, errors.New("owner cannot add themself to channel")
	}

	var members []*models.User

	for _, memberID := range memberIDs {
		member, err := cs.userRepository.UserById(memberID)
		if err != nil {
			return nil, errors.New("error retrieving member")
		}

		members = append(members, member)
	}

	now := time.Now().UTC()

	c := &models.Channel{
		ID:        uuid.NewString(),
		OwnerID:   owner.ID,
		Name:      name,
		Members:   members,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := cs.channelRepository.CreateChannel(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (cs *ChannelService) GetChannelByID(id string) (*models.Channel, error) {
	c, err := cs.channelRepository.ChannelById(id)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, errors.New("channel not found")
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

	if channels == nil {
		return make([]*models.Channel, 0), nil
	}

	return channels, nil
}
