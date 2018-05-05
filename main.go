package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/shjp/shjp-server/auth"
	"github.com/shjp/shjp-server/db"
	"github.com/shjp/shjp-server/schema"
)

var (
	dbUser     = flag.CommandLine.String("user", "shjp", "Postgres username")
	dbName     = flag.CommandLine.String("dbname", "shjp_youth", "Postgres database name")
	dbPassword = flag.CommandLine.String("password", "hellochurch", "Password for the postgres database")
	dbHost     = flag.CommandLine.String("host", "localhost", "Host for Postgres database")
)

func main() {
	schema, err := schema.ConfigSchema()
	if err != nil {
		log.Fatalf("Failed configuring schema: %v", err)
		return
	}

	connConfig := db.ConnConfig{
		User:     *dbUser,
		Password: *dbPassword,
		DBName:   *dbName,
		Host:     *dbHost,
		SSLMode:  "disable"}
	if err = db.Init(connConfig); err != nil {
		log.Printf("Failed initializing DB: %s", err)
		return
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	http.Handle("/login", auth.NewLoginHandler())

	log.Println("Server listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
