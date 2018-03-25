package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/shjp/shjp-server/auth"
	"github.com/shjp/shjp-server/db"
	"github.com/shjp/shjp-server/schema"
)

func main() {
	schema, err := schema.ConfigSchema()
	if err != nil {
		log.Fatalf("Failed configuring schema: %v", err)
		return
	}

	if err = db.Init(); err != nil {
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
	http.ListenAndServe(":8080", nil)
}
