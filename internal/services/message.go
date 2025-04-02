package services

import (
	"github.com/bnock/nockchat-api-go/internal/models"
	"github.com/bnock/nockchat-api-go/internal/repositories"
)

type MessageService struct {
	messageRepository *repositories.MessageRepository
}

func (ms *MessageService) GetMessagesByChannel(channel *models.Channel) ([]models.Message, error) {
	messages, err := ms.messageRepository.AllByChannel(channel)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
