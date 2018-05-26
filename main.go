package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/shjp/shjp-server/auth"
	"github.com/shjp/shjp-server/db"
	"github.com/shjp/shjp-server/schema"
)

var (
	dbUser     = flag.CommandLine.String("user", "shjp", "Postgres username")
	dbName     = flag.CommandLine.String("dbname", "shjp_youth", "Postgres database name")
	dbPassword = flag.CommandLine.String("password", "hellochurch", "Password for the postgres database")
	dbHost     = flag.CommandLine.String("host", "localhost", "Host for Postgres database")

	Schema graphql.Schema
)

func main() {
	flag.Parse()

	var err error
	Schema, err = schema.ConfigSchema()
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

	// Uncomment the below to test with GraphQL interface
	//h := handler.New(&handler.Config{Schema: &schema, Pretty: true, GraphiQL: true})
	//http.Handle("/graphql", h)

	http.HandleFunc("/graphql", graphqlHandler)

	// TODO: move this to the GraphQL mutation type
	http.HandleFunc("/login/kakao", auth.HandleKakaoLogin)

	log.Println("Server listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("token")
	log.Printf("url = %+v", r.URL)
	log.Printf("query = %+v", r.URL.Query())

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: r.URL.Query().Get("query"),
		Context:       context.WithValue(context.Background(), "token", authToken),
	})
	if len(result.Errors) > 0 {
		log.Printf("graphql errors: %v", result.Errors)
		return
	}

	json.NewEncoder(w).Encode(result)
}
