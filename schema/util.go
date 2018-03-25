package schema

import (
	"log"

	"github.com/graphql-go/graphql"
)

// PickFields returns a subset of the given table picking only the given names as keys
func PickFields(table map[string]*graphql.Field, names []string) graphql.Fields {
	if len(names) == 0 {
		return table
	}

	var fields graphql.Fields
	for _, name := range names {
		t, ok := table[name]
		if !ok {
			log.Fatalf("The field %s does not exist!", name)
			continue
		}
		fields[name] = t
	}
	return fields
}
