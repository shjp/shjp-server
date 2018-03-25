package models

import (
	"github.com/satori/go.uuid"

	"github.com/shjp/shjp-server/db"
)

// Role is the user role associated with the group
type Role struct {
	// Core fields
	ID        string `json:"id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Privilege int8   `json:"privilege"`
}

// Create inserts a row in the roles table
func (r *Role) Create() error {
	tx, err := db.Tx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	r.ID = id.String()

	_, err = tx.Exec(`
		INSERT INTO roles (
			id,
			name,
			group_id,
			privilege
		) VALUES ($1, $2, $3, $4)`,
		r.ID,
		r.Name,
		r.Group,
		r.Privilege)

	if err != nil {
		return err
	}

	return tx.Commit()
}
