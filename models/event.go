package models

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/satori/go.uuid"

	"github.com/shjp/shjp-server/db"
)

// Event model
type Event struct {
	// Core fields
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	Date                *string `json:"date"`
	Length              int     `json:"length"`
	Creator             *string `json:"creator"`
	Deadline            *string `json:"deadline"`
	AllowMaybe          bool    `json:"allow_maybe"`
	Description         string  `json:"description"`
	Location            *string `json:"location"`
	LocationDescription *string `json:"location_description"`
}

// Create insert event row into the events table
func (e *Event) Create(groupIDs []string) error {
	tx, err := db.Tx()
	if err != nil {
		return errors.Wrap(err, "Failed initializing transaction")
	}

	defer tx.Rollback()

	id, err := uuid.NewV4()
	if err != nil {
		return errors.Wrap(err, "Failed generating UUID for event")
	}

	e.ID = id.String()

	_, err = tx.Exec(`
		INSERT INTO events (
			id,
			name,
			date,
			length,
			creator,
			deadline,
			allow_maybe,
			description,
			location,
			location_description
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		e.ID,
		e.Name,
		e.Date,
		e.Length,
		e.Creator,
		e.Deadline,
		e.AllowMaybe,
		e.Description,
		e.Location,
		e.LocationDescription)
	if err != nil {
		return errors.Wrap(err, "Failed executing insert query into events")
	}

	rows := make([]string, len(groupIDs))
	for i, gid := range groupIDs {
		rows[i] = fmt.Sprintf("('%s', '%s')", gid, e.ID)
	}
	log.Printf("rows = %v", rows)
	_, err = tx.Exec(fmt.Sprintf(`
		INSERT INTO groups_events (
			group_id,
			event_id
		) VALUES %s`,
		strings.Join(rows, ", ")))
	if err != nil {
		return errors.Wrap(err, "Failed executing insert query into groups_events")
	}

	return tx.Commit()
}

// Find finds all the events associated with the given group IDs
// If group IDs are not given, it returns all the events
func (e *Event) Find(groupIDs []string) ([]*Event, error) {
	tx, err := db.Tx()
	if err != nil {
		return nil, errors.Wrap(err, "Failed initializing transaction")
	}

	defer tx.Rollback()

	query := `
		SELECT
			id,
			name,
			date,
			length,
			creator,
			deadline,
			allow_maybe,
			description,
			location,
			location_description
		FROM events e`
	if groupIDs != nil {
		var stringified []string
		for _, gid := range groupIDs {
			stringified = append(stringified, fmt.Sprintf("'%s'", gid))
		}
		query += fmt.Sprintf(`
			INNER JOIN groups_events ge ON e.id = ge.event_id
			WHERE ge.group_id IN (%s)`,
			strings.Join(stringified, ", "))
	}

	rows, err := tx.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "Failed finding event rows")
	}

	var es []*Event
	for rows.Next() {
		e := &Event{}
		err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Date,
			&e.Length,
			&e.Creator,
			&e.Deadline,
			&e.AllowMaybe,
			&e.Description,
			&e.Location,
			&e.LocationDescription)
		if err != nil {
			return nil, errors.Wrap(err, "Failed reading a row to event object")
		}

		es = append(es, e)
	}

	return es, nil
}
