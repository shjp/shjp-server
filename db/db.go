package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB is a global database connection instance
var DB *sql.DB

// ConnConfig is the configuration for the database connection
type ConnConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	SSLMode  string
}

func (c ConnConfig) String() string {
	return fmt.Sprintf(`
		user=%s dbname=%s password=%s host=%s sslmode=%s`,
		c.User, c.DBName, c.Password, c.Host, c.SSLMode)
}

// Init initializes the database connection
func Init(config ConnConfig) error {
	db, err := sql.Open("postgres", config.String())
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
