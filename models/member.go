package models

import (
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/shjp/shjp-server/constant"
	"github.com/shjp/shjp-server/db"
)

// Member represents a user
type Member struct {
	// Core fields
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	BaptismalName *string              `json:"baptismalName"`
	Birthday      *string              `json:"birthday"`
	FeastDay      *string              `json:"feastDay"`
	Groups        []string             `json:"groups"`
	Created       *string              `json:"created"`
	LastActive    *string              `json:"lastActive"`
	AccountType   constant.AccountType `json:"accountType"`
	AccountHash   string               `json:"-"`

	// Extra fields
	RoleName    *string            `json:"roleName"`
	Privilege   *int               `json:"privilege"`
	Permissions []GroupPermissions `json:"groupPermissions"`
}

// NewMember is a public constructor for member struct
func NewMember() Member {
	m := Member{}
	m.AccountType = constant.Undefined
	return m
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
			account_type,
			account_hash
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		m.ID,
		m.Name,
		m.BaptismalName,
		m.Birthday,
		m.FeastDay,
		m.AccountType,
		m.AccountHash)

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
			m.account_type,
			m.account_hash
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
			&m.AccountType,
			&m.AccountHash,
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

// FindMe finds the member with AccountType and AccountHash
func (m *Member) FindMe() error {
	if m.AccountType == constant.Undefined {
		return errors.New("FindMe expects a well defined account type")
	}

	if m.AccountHash == "" {
		return errors.New("FindMe expects account hash")
	}

	tx, err := db.Tx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err = tx.QueryRow(`
		SELECT (
			id,
			name,
			baptismal_name,
			birthday,
			feast_day,
			role
		)
		FROM members
		WHERE account_type = $1
		AND account_hash = $2`).Scan(
		&m.ID,
		&m.Name,
		&m.BaptismalName,
		&m.Birthday,
		&m.FeastDay); err != nil {
		return errors.Wrap(err, "Failed querying me")
	}

	return nil
}

func (m *Member) PopulatePermissions() error {
	tx, err := db.Tx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	rows, err := tx.Query(`
		SELECT r.name, r.privilege, r.group_id
		FROM roles r
		INNER JOIN groups_members gm ON gm.role_id = r.id
		AND gm.member_id = $1`,
		m.ID)
	if err != nil {
		return errors.Wrap(err, "failed querying permissions")
	}

	for rows.Next() {
		var roleName string
		var groupName string
		var privilege Privilege
		if err = rows.Scan(
			&roleName,
			&privilege,
			&groupName); err != nil {
			return errors.Wrap(err, "Failed scanning role row")
		}
		m.Permissions = append(m.Permissions, GroupPermissions{
			GroupName:   groupName,
			RoleName:    roleName,
			Permissions: privilege.Expand(),
		})
	}

	return nil
}
