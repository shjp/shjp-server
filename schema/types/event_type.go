package types

import (
	"github.com/graphql-go/graphql"
)

// EventType defines the GraphQL event type
var EventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "event",
	Fields: graphql.Fields{
		"id":                   &graphql.Field{Type: graphql.ID},
		"name":                 &graphql.Field{Type: graphql.String},
		"date":                 &graphql.Field{Type: graphql.DateTime},
		"length":               &graphql.Field{Type: graphql.Int},
		"creator":              &graphql.Field{Type: MemberType},
		"deadline":             &graphql.Field{Type: graphql.DateTime},
		"allow_maybe":          &graphql.Field{Type: graphql.Boolean},
		"description":          &graphql.Field{Type: graphql.String},
		"location":             &graphql.Field{Type: graphql.String},
		"location_description": &graphql.Field{Type: graphql.String},
	},
})
