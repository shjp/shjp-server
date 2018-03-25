package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/shjp/shjp-server/schema/types"
)

var Subscriptions = graphql.Fields{
	"member": &graphql.Field{
		Type: types.MemberType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
	},
}
