package repositories

import (
	"database/sql"

	"github.com/bnock/nockchat-api-go/internal/models"
)

type MessageRepository struct {
	DB *sql.DB
}

func (mr *MessageRepository) CreateMessage(m models.Message) error {
	_, err := mr.DB.Exec(`
		INSERT INTO messages (
	  		id, 
		  	channel_id, 
		  	sender_id, 
		  	content, 
		  	sent_at, 
		  	created_at, 
		  	updated_at
		) VALUES (?, ?, ?, ?, NOW(), NOW(), NOW())`,
		m.ID,
		m.ChannelID,
		m.SenderID,
		m.Content,
	)

	return err
}

func (mr *MessageRepository) AllByChannelID(channelID string) ([]models.Message, error) {
	rows, err := mr.DB.Query(`
		SELECT 
		    id,
		    channel_id,
		    sender_id,
		    content,
		    sent_at,
		    created_at,
		    deleted_at
		FROM 
		    messages 
		WHERE 
		    channel_id = ? 
		  	AND deleted_at IS NULL
		ORDER BY 
		    messages.sent_at DESC`,
		channelID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message

	for rows.Next() {
		var m models.Message

		if err := rows.Scan(
			&m.ID,
			&m.ChannelID,
			&m.SenderID,
			&m.Content,
			&m.SentAt,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,
		); err != nil {
			return nil, err
		}

		messages = append(messages, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
