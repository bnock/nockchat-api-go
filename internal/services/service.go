package services

import "github.com/bnock/nockchat-api-go/internal/repositories"

type Services struct {
	ChannelService  *ChannelService
	MessageService  *MessageService
	UserService     *UserService
	SecurityService *SecurityService
}

func NewServices(repositories *repositories.Repositories) *Services {
	return &Services{
		ChannelService: &ChannelService{
			channelRepository: repositories.ChannelRepository,
			userRepository:    repositories.UserRepository,
		},
		MessageService:  &MessageService{messageRepository: repositories.MessageRepository},
		UserService:     &UserService{userRepository: repositories.UserRepository},
		SecurityService: &SecurityService{userRepository: repositories.UserRepository},
	}
}
