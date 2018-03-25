package models

import (
	"github.com/satori/go.uuid"

	"github.com/shjp/shjp-server/db"
)

// Member represents a user
type Member struct {
	// Core fields
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	BaptismalName *string  `json:"baptismalName"`
	Birthday      *string  `json:"birthday"`
	FeastDay      *string  `json:"feastDay"`
	Groups        []string `json:"groups"`
	Created       *string  `json:"created"`
	LastActive    *string  `json:"lastActive"`
	GoogleID      *string  `json:"googleId"`
	FacebookID    *string  `json:"facebookId"`

	// Extra fields
	RoleName  *string `json:"roleName"`
	Privilege *int    `json:"privilege"`
}

// Create inserts a row of member
func (m *Member) Create() error {
	tx, err := db.Tx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	m.ID = id.String()

	_, err = tx.Exec(`
		INSERT INTO members (
			id,
			name,
			baptismal_name,
			birthday,
			feast_day,
			google_id,
			facebook_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		m.ID,
		m.Name,
		m.BaptismalName,
		m.Birthday,
		m.FeastDay,
		m.GoogleID,
		m.FacebookID)

	if err != nil {
		return err
	}

	return tx.Commit()
}

// FindAll returns all the members in DB
func (m *Member) FindAll() ([]*Member, error) {
	tx, err := db.Tx()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var ms []*Member
	rows, err := tx.Query(`
		SELECT
			m.id,
			m.name,
			m.baptismal_name,
			m.birthday,
			m.feast_day,
			m.google_id,
			m.facebook_id,
			r.name,
			r.privilege
		FROM members m
		INNER JOIN groups_members gm ON gm.member_id = m.id
		INNER JOIN roles r ON r.id = gm.role_id`)
	if err != nil {
		return ms, err
	}

	for rows.Next() {
		m := &Member{}
		err = rows.Scan(
			&m.ID,
			&m.Name,
			&m.BaptismalName,
			&m.Birthday,
			&m.FeastDay,
			&m.GoogleID,
			&m.FacebookID,
			&m.RoleName,
			&m.Privilege)
		if err != nil {
			return ms, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

// AddToGroup associates a members with a group
func (m *Member) AddToGroup(groupID string, status string) error {
	tx, err := db.Tx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO groups_members (
			member_id,
			group_id,
			status
		)
		VALUES ($1, $2, $3)
		ON CONFLICT (member_id, group_id)
		DO UPDATE SET status = $4`,
		m.ID,
		groupID,
		status,
		status)

	if err != nil {
		return err
	}

	return tx.Commit()
}
