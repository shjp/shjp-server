package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/shjp/shjp-server/db"
	"github.com/shjp/shjp-server/schema"
	"github.com/shjp/shjp-server/utils"
)

var (
	dbUser     = flag.CommandLine.String("user", "shjp", "Postgres username")
	dbName     = flag.CommandLine.String("dbname", "shjp_youth", "Postgres database name")
	dbPassword = flag.CommandLine.String("password", "hellochurch", "Password for the postgres database")
	dbHost     = flag.CommandLine.String("host", "localhost", "Host for Postgres database")

	Schema graphql.Schema
)

type GraphQLPostBody struct {
	Query string `json:"query"`
}

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

	// Below handler is for dev purpose
	interactiveGqHandler := handler.New(&handler.Config{Schema: &Schema, Pretty: true, GraphiQL: true})
	http.Handle("/console", interactiveGqHandler)

	http.HandleFunc("/graphql", graphqlHandler)

	log.Println("Server listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("token")
	log.Printf("url = %+v", r.URL)

	var requestString string
	if r.Method == http.MethodGet {
		requestString = r.URL.Query().Get("query")
	} else if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			utils.SendErrorResponse(w, errors.New("Cannot read POST body"), 400)
			return
		}
		var queryWrapper GraphQLPostBody
		if err = json.Unmarshal(body, &queryWrapper); err != nil {
			utils.SendErrorResponse(w, errors.New("Failed parsing POST body. Did you forget query property?"), 400)
			return
		}
		requestString = queryWrapper.Query
	} else {
		utils.SendErrorResponse(w, errors.New("Only GET and POST allowed for GraphQL request"), 400)
		return
	}
	log.Printf("query = %s", requestString)

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: requestString,
		Context:       context.WithValue(context.Background(), "token", authToken),
	})
	if len(result.Errors) > 0 {
		log.Printf("graphql errors: %v", result.Errors)
		return
	}

	fmt.Printf("data = %+v\n", result.Data)
	fmt.Printf("errors = %+v\n", result.Errors)

	respJSON, err := json.Marshal(result)
	if err != nil {
		log.Printf("Error marshaling result: %s", err)
		utils.SendErrorResponse(w, err, 500)
		return
	}

	utils.SendResponse(w, string(respJSON), 200)
}
