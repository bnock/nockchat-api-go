package repositories

import "database/sql"

type Repositories struct {
	ChannelRepository *ChannelRepository
	MessageRepository *MessageRepository
	UserRepository    *UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		ChannelRepository: &ChannelRepository{DB: db},
		MessageRepository: &MessageRepository{DB: db},
		UserRepository:    &UserRepository{DB: db},
	}
}
