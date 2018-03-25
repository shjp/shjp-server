package schema

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"

	"github.com/shjp/shjp-server/models"
)

func createGroup(p graphql.ResolveParams) (interface{}, error) {
	group := &models.Group{}

	if name := p.Args["name"]; name != nil {
		group.Name = name.(string)
	}

	if description := p.Args["description"]; description != nil {
		group.Description = description.(string)
	}

	err := group.Create()
	if err != nil {
		log.Println(err)
	}

	return group, err
}

func createRole(p graphql.ResolveParams) (interface{}, error) {
	role := &models.Role{}

	if name := p.Args["name"]; name != nil {
		role.Name = name.(string)
	}

	if group := p.Args["group"]; group != nil {
		role.Group = group.(string)
	}

	err := role.Create()
	if err != nil {
		log.Println(err)
	}

	return role, err
}

func createMember(p graphql.ResolveParams) (interface{}, error) {
	member := &models.Member{}

	if name := p.Args["name"]; name != nil {
		member.Name = name.(string)
	}

	if baptismalName := p.Args["baptismalName"]; baptismalName != nil {
		bnStr := baptismalName.(string)
		member.BaptismalName = &bnStr
	}

	if birthday := p.Args["birthday"]; birthday != nil {
		bStr := birthday.(string)
		member.Birthday = &bStr
	}

	if feastDay := p.Args["feastDay"]; feastDay != nil {
		fdStr := feastDay.(string)
		member.FeastDay = &fdStr
	}

	if lastActive := p.Args["lastActive"]; lastActive != nil {
		laStr := lastActive.(string)
		member.LastActive = &laStr
	}

	if googleID := p.Args["googleId"]; googleID != nil {
		gidStr := googleID.(string)
		member.GoogleID = &gidStr
	}

	if facebookID := p.Args["facebookId"]; facebookID != nil {
		fidStr := facebookID.(string)
		member.FacebookID = &fidStr
	}

	err := member.Create()
	if err != nil {
		log.Println(err)
	}

	return member, err
}

func registerGroupMember(p graphql.ResolveParams) (interface{}, error) {
	member := p.Args["member"]
	if member == nil {
		return nil, errors.New("member parameter is required")
	}

	group := p.Args["group"]
	if group == nil {
		return nil, errors.New("group parameter is required")
	}

	status := p.Args["status"]
	if status == nil {
		return nil, errors.New("status parameter is required")
	}

	m := &models.Member{
		ID: member.(string),
	}
	err := m.AddToGroup(group.(string), status.(string))
	if err != nil {
		log.Println(err)
	}

	return m, err
}

func createEvent(p graphql.ResolveParams) (interface{}, error) {
	e := &models.Event{}

	if name := p.Args["name"]; name != nil {
		e.Name = name.(string)
	}

	if date := p.Args["date"]; date != nil {
		dateStr := date.(string)
		e.Date = &dateStr
	}

	if length := p.Args["length"]; length != nil {
		e.Length = length.(int)
	}

	if creator := p.Args["creator"]; creator != nil {
		creatorStr := creator.(string)
		e.Creator = &creatorStr
	}

	if deadline := p.Args["deadline"]; deadline != nil {
		deadlineStr := deadline.(string)
		e.Deadline = &deadlineStr
	}

	if allowMaybe := p.Args["allow_maybe"]; allowMaybe != nil {
		e.AllowMaybe = allowMaybe.(bool)
	}

	if description := p.Args["description"]; description != nil {
		e.Description = description.(string)
	}

	if location := p.Args["location"]; location != nil {
		locationStr := location.(string)
		e.Location = &locationStr
	}

	if locationDescription := p.Args["location_description"]; locationDescription != nil {
		ldStr := locationDescription.(string)
		e.LocationDescription = &ldStr
	}

	var groupIDs []string
	for _, gid := range p.Args["group_ids"].([]interface{}) {
		groupIDs = append(groupIDs, gid.(string))
	}

	err := e.Create(groupIDs)
	if err != nil {
		log.Printf("Failed creating event: %v\n", err)
	}

	return e, err
}
