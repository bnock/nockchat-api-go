package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/bnock/nockchat-api-go/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (ur *UserRepository) All() ([]models.User, error) {
	rows, err := ur.DB.Query(`
		SELECT 
		    id,
		    first_name,
		    last_name,
		    email,
		    password,
		    created_at,
		    updated_at,
		    deleted_at
		FROM 
		    users 
		WHERE 
		    deleted_at IS NULL`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User

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

func (ur *UserRepository) UserById(id string) (*models.User, error) {
	row := ur.DB.QueryRow(`
		SELECT 
		    id,
		    first_name,
		    last_name,
		    email,
		    password,
		    created_at,
		    updated_at,
		    deleted_at
		FROM 
		    users 
		WHERE 
		    id = ? 
		  	AND deleted_at IS NULL`,
		id,
	)

	var u models.User

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

func (ur *UserRepository) UserByEmail(email string) (*models.User, error) {
	row := ur.DB.QueryRow(`
		SELECT 
		    id,
		    first_name,
		    last_name,
		    email,
		    password,
		    created_at,
		    updated_at,
		    deleted_at
		FROM 
		    users 
		WHERE 
		    email LIKE ? 
		  	AND deleted_at IS NULL`,
		strings.ToLower(email),
	)

	var u models.User

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
			return nil, errors.New(fmt.Sprintf("User with email '%s' not found", email))
		}

		return nil, err
	}

	return &u, nil
}
