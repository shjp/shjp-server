package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/shjp/shjp-server/schema/types"
)

var Queries = graphql.Fields{

	"group": &graphql.Field{
		Type: types.GroupType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
		Resolve: group,
	},

	"groups": &graphql.Field{
		Type:    graphql.NewList(types.GroupType),
		Resolve: groups,
	},

	"member": &graphql.Field{
		Type: types.MemberType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
		},
	},

	"members": &graphql.Field{
		Type:    graphql.NewList(types.MemberType),
		Resolve: members,
	},

	"events": &graphql.Field{
		Type: graphql.NewList(types.EventType),
		Args: graphql.FieldConfigArgument{
			"groupIds": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.ID),
			},
		},
		Resolve: events,
	},
}
