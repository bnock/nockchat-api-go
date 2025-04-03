package services

import (
	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/repositories"
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
