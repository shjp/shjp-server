package types

import (
	"github.com/graphql-go/graphql"
)

// MemberType defines the GraphQL member type
var MemberType = graphql.NewObject(graphql.ObjectConfig{
	Name: "member",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.ID},
		"name":          &graphql.Field{Type: graphql.String},
		"baptismalName": &graphql.Field{Type: graphql.String},
		"birthday":      &graphql.Field{Type: graphql.DateTime},
		"feastDay":      &graphql.Field{Type: graphql.DateTime},
		"groups":        &graphql.Field{Type: graphql.NewList(graphql.String)},
		"created":       &graphql.Field{Type: graphql.DateTime},
		"lastActive":    &graphql.Field{Type: graphql.DateTime},
		"goolgeId":      &graphql.Field{Type: graphql.String},
		"facebookId":    &graphql.Field{Type: graphql.String},
		"roleName":      &graphql.Field{Type: graphql.String},
		"privilege":     &graphql.Field{Type: graphql.Int},
	},
})
