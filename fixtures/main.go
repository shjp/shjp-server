package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shjp/shjp-server/db"
)

var (
	dbUser     = flag.CommandLine.String("user", "shjp", "Postgres username")
	dbName     = flag.CommandLine.String("dbname", "shjp_youth", "Postgres database name")
	dbPassword = flag.CommandLine.String("password", "hellochurch", "Password for the postgres database")
	dbHost     = flag.CommandLine.String("host", "localhost", "Host for Postgres database")
)

// Order matters here..
var files = []string{
	"groups",
	"members",
	"roles",
	"events",
	"announcements",
	"comments",
	"groups_members",
	"groups_events",
	"groups_announcements",
	"members_events",
}

func insert(tx *sql.Tx, table string) error {
	file, err := os.Open(fmt.Sprintf("fixtures/data/%s.csv", table))
	if err != nil {
		log.Fatalf("Error reading file %s: %s", table, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		query := fmt.Sprintf(`
			INSERT INTO %s
			VALUES (%s)`,
			table, scanner.Text())

		log.Printf("query: %s\n\n", query)
		if _, err = tx.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()

	err := db.Init(db.ConnConfig{
		User:     *dbUser,
		Password: *dbPassword,
		DBName:   *dbName,
		Host:     *dbHost,
		SSLMode:  "disable"})
	if err != nil {
		log.Printf("Error initializing database: %s", err)
		os.Exit(1)
	}

	tx, err := db.DB.Begin()
	if err != nil {
		log.Printf("Error beginning transaction: %s", err)
		os.Exit(1)
	}
	defer tx.Rollback()

	for _, name := range files {
		if err = insert(tx, name); err != nil {
			log.Printf("Error executing query: %s", err)
		}
	}

	if err = tx.Commit(); err != nil {
		log.Printf("Error committing the changes: %s", err)
		os.Exit(1)
	}
}
