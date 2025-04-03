package services

import (
	"time"

	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/repositories"
	"github.com/google/uuid"
)

type MessageService struct {
	messageRepository *repositories.MessageRepository
}

func (ms *MessageService) GetMessagesByChannel(c *models.Channel) ([]models.Message, error) {
	messages, err := ms.messageRepository.AllByChannelID(c.ID)
	if err != nil {
		return nil, err
	}

	if messages == nil {
		return make([]models.Message, 0), nil
	}

	return messages, nil
}

func (ms *MessageService) CreateMessage(u *models.User, c *models.Channel, content string) (*models.Message, error) {
	now := time.Now().UTC()

	m := &models.Message{
		ID:        uuid.NewString(),
		ChannelID: c.ID,
		SenderID:  u.ID,
		Content:   content,
		SentAt:    &now,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	err := ms.messageRepository.CreateMessage(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
