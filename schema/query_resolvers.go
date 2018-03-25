package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/shjp/shjp-server/models"
)

func group(p graphql.ResolveParams) (interface{}, error) {
	g := &models.Group{
		ID: p.Args["id"].(string),
	}
	err := g.Find()
	if err != nil {
		return nil, err
	}

	return g, err
}

func groups(_ graphql.ResolveParams) (interface{}, error) {
	gs, err := (&models.Group{}).FindAll()
	return gs, err
}

func members(_ graphql.ResolveParams) (interface{}, error) {
	return (&models.Member{}).FindAll()
}

func events(p graphql.ResolveParams) (interface{}, error) {
	var groupIDs []string
	for _, gid := range p.Args["groupIds"].([]interface{}) {
		groupIDs = append(groupIDs, gid.(string))
	}
	return (&models.Event{}).Find(groupIDs)
}
