package types

import (
	"github.com/graphql-go/graphql"
)

// UserSessionType defines the GraphQL user session type
var UserSessionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "userSession",
	Fields: graphql.Fields{
		"token": &graphql.Field{Type: graphql.String},
	},
})
