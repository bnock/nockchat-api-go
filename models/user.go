package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (um UserModel) All() ([]User, error) {
	rows, err := um.DB.Query("SELECT * FROM users WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		if err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.Password,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.DeletedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (um UserModel) UserById(id string) (*User, error) {
	row := um.DB.QueryRow("SELECT * FROM users WHERE id = ? AND deleted_at IS NULL", id)

	var u User

	if err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(fmt.Sprintf("User with id '%s' not found", id))
		}

		return nil, err
	}

	return &u, nil
}
