package types

import (
	"github.com/graphql-go/graphql"
)

// RoleType defines the GraphQL role type
var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "role",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.ID},
		"group":     &graphql.Field{Type: graphql.ID},
		"name":      &graphql.Field{Type: graphql.String},
		"privilege": &graphql.Field{Type: graphql.Int},
	},
})
