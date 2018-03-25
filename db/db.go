package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DB is a global database connection instance
var DB *sql.DB

// Init initializes the database connection
func Init() error {
	connConfig := `
		user=shjp
		dbname=shjp_youth
		password=hellochurch
		sslmode=disable
	`

	db, err := sql.Open("postgres", connConfig)
	if err != nil {
		return err
	}

	DB = db

	return nil
}

// Tx starts and returns a database transaction
func Tx() (*sql.Tx, error) {
	return DB.Begin()
}
