package types

import (
	"github.com/graphql-go/graphql"
)

// GroupType defines the GraphQL group type
var GroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "group",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"members":     &graphql.Field{Type: graphql.NewList(MemberType)},
	},
})
