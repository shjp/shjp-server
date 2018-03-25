package schema

import (
	"github.com/graphql-go/graphql"
)

// ConfigSchema returns the root level GraphQL schema instance
func ConfigSchema() (graphql.Schema, error) {
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Query",
			Fields: Queries,
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Mutation",
			Fields: Mutations,
		}),
		Subscription: graphql.NewObject(graphql.ObjectConfig{
			Name:   "Subscription",
			Fields: Subscriptions,
		}),
	}

	return graphql.NewSchema(schemaConfig)
}
