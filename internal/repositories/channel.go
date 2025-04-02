package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/bnock/nockchat-api-go/internal/models"
)

type ChannelRepository struct {
	DB *sql.DB
}

func (cr *ChannelRepository) ChannelById(id string) (*models.Channel, error) {
	row := cr.DB.QueryRow(`
		SELECT 
		    id, 
		    name, 
		    owner_id, 
		    created_at, 
		    updated_at, 
		    deleted_at 
		FROM 
		    channels 
		WHERE id = ? 
		  AND deleted_at IS NULL`,
		id,
	)

	var c models.Channel

	err := row.Scan(&c.ID, &c.Name, &c.OwnerID, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprintf("Channel with id '%s' not found", id))
		}

		return nil, err
	}

	return &c, nil
}

func (cr *ChannelRepository) MembersByChannelID(id string) ([]*models.User, error) {
	rows, err := cr.DB.Query(`
		SELECT 
		    users.id,
		    users.first_name,
		    users.last_name,
		    users.email,
		    users.password,
		    users.created_at,
		    users.updated_at,
		    users.deleted_at
		FROM 
		    users 
		    JOIN channel_user ON users.id = channel_user.user_id 
		WHERE 
		    channel_user.channel_id = ?
		    AND users.deleted_at IS NULL`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*models.User

	for rows.Next() {
		var m models.User

		if err := rows.Scan(
			&m.ID,
			&m.FirstName,
			&m.LastName,
			&m.Email,
			&m.Password,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.DeletedAt,
		); err != nil {
			return nil, err
		}

		members = append(members, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func (cr *ChannelRepository) ChannelsByUserID(id string) ([]*models.Channel, error) {
	rows, err := cr.DB.Query(`
		SELECT
		    channels.id,
		    channels.name,
			channels.owner_id,
			channels.created_at,
			channels.updated_at,
			channels.deleted_at
		FROM
		    channels
			JOIN channel_user ON channel_user.channel_id = channels.id
		WHERE
		    channels.deleted_at IS NULL
			AND (
			    channels.owner_id = ?
			    OR channel_user.user_id = ?
			)`,
		id,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []*models.Channel

	for rows.Next() {
		var c models.Channel

		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.OwnerID,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.DeletedAt,
		); err != nil {
			return nil, err
		}

		channels = append(channels, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}
