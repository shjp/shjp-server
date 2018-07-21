package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/shjp/shjp-server/auth"
	"github.com/shjp/shjp-server/schema/types"
)

var Mutations = graphql.Fields{

	"login": &graphql.Field{
		Type:    types.UserSessionType,
		Resolve: login,
		Args: graphql.FieldConfigArgument{
			"accountId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"clientId": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"accountType": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"accountSecret": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"profileImage": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"nickname": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
	},

	"createGroup": &graphql.Field{
		Type:    types.GroupType,
		Resolve: createGroup,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"imageUri": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
	},

	"createRole": &graphql.Field{
		Type:    types.RoleType,
		Resolve: auth.Authenticate(createRole),
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"group": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"privilege": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
	},

	"createMember": &graphql.Field{
		Type:    types.MemberType,
		Resolve: createMember,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"birthday": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"baptismalName": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"feastDay": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"groups": &graphql.ArgumentConfig{
				Type: graphql.NewList(graphql.String),
			},
			"googleId": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"facebookId": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"kakaoId": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"accountSecret": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},

	"registerGroupMember": &graphql.Field{
		Type:    types.MemberType,
		Resolve: registerGroupMember,
		Args: graphql.FieldConfigArgument{
			"member": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"group": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},

	"createEvent": &graphql.Field{
		Type:    types.EventType,
		Resolve: createEvent,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"date": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"length": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"creator": &graphql.ArgumentConfig{
				Type: graphql.ID,
			},
			"deadline": &graphql.ArgumentConfig{
				Type: graphql.DateTime,
			},
			"allow_maybe": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"location": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"location_description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"group_ids": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.NewList(graphql.ID)),
			},
		},
	},
}
