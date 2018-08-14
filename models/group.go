package models

import (
	"fmt"

	"github.com/satori/go.uuid"

	"github.com/shjp/shjp-server/db"
)

// Group is a group model
type Group struct {
	// Core fields
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ImageURI    interface{} `json:"imageUri"` // nullable string

	// Extra fields
	Members []*Member `json:"members"`
}

// Create inserts a row for group
func (g *Group) Create() error {
	tx, err := db.Tx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	g.ID = id.String()

	_, err = tx.Exec(`
		INSERT INTO groups (
			id,
			name,
			description,
			image_uri
		) VALUES ($1, $2, $3, $4)`,
		g.ID,
		g.Name,
		g.Description,
		g.ImageURI)

	if err != nil {
		return err
	}

	return tx.Commit()
}

// Find returns a group matching the given ID
func (g *Group) Find() error {
	tx, err := db.Tx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err = tx.
		QueryRow(`
			SELECT
				name,
				description,
				image_uri
			FROM groups
			WHERE id = $1`, g.ID).
		Scan(
			&g.Name,
			&g.Description,
			&g.ImageURI); err != nil {
		return err
	}

	rows, err := tx.Query(`
		SELECT
			m.id,
			m.name,
			m.baptismal_name,
			m.birthday,
			m.feast_day,
			r.name,
			r.privilege
		FROM members m
		INNER JOIN groups_members gm ON m.id = gm.member_id
		INNER JOIN roles r ON gm.role_id = r.id
		INNER JOIN groups g ON g.id = gm.group_id
		WHERE g.id = $1`,
		g.ID)
	if err != nil {
		return err
	}

	for rows.Next() {
		m := &Member{}
		err = rows.Scan(
			&m.ID,
			&m.Name,
			&m.BaptismalName,
			&m.Birthday,
			&m.FeastDay,
			&m.RoleName,
			&m.Privilege)
		if err != nil {
			return err
		}
		fmt.Printf("%+v", m)
		g.Members = append(g.Members, m)
	}

	return nil
}

// FindAll returns all the groups
func (g *Group) FindAll() ([]*Group, error) {
	tx, err := db.Tx()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var gs []*Group
	rows, err := tx.Query(`
		SELECT
			id,
			name,
			description,
			image_uri
		FROM groups`)
	if err != nil {
		return gs, err
	}

	for rows.Next() {
		g := &Group{}
		err = rows.Scan(
			&g.ID,
			&g.Name,
			&g.Description,
			&g.ImageURI)
		if err != nil {
			return gs, err
		}

		gs = append(gs, g)
	}

	return gs, nil
}
