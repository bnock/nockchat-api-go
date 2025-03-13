package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Channel struct {
	ID        string
	Name      string
	OwnerID   string
	MemberIDs []string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type ChannelModel struct {
	DB *sql.DB
}

func (cm ChannelModel) ChannelById(id string) (*Channel, error) {
	row := cm.DB.QueryRow(`
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

	var c Channel

	err := row.Scan(&c.ID, &c.Name, &c.OwnerID, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprintf("Channel with id '%s' not found", id))
		}

		return nil, err
	}

	rows, err := cm.DB.Query("SELECT user_id FROM channel_user WHERE channel_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memberIDs []string

	for rows.Next() {
		var memberID string

		if err := rows.Scan(&memberID); err != nil {
			return nil, err
		}

		memberIDs = append(memberIDs, memberID)
	}

	c.MemberIDs = memberIDs

	return &c, nil
}
