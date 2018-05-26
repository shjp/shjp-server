package auth

import (
	"fmt"

	"github.com/graphql-go/graphql"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/shjp/shjp-server/models"
	"github.com/shjp/shjp-server/session"
)

var (
	signingKey     = []byte("foobar")
	sessionManager = session.NewManager(int64(0))
)

// Authenticate is the authentication wrapper for GraphQL resolvers
func Authenticate(resolver func(graphql.ResolveParams) (interface{}, error)) func(graphql.ResolveParams) (interface{}, error) {
	return func(p graphql.ResolveParams) (interface{}, error) {
		token := p.Context.Value("token").(string)
		_, ok := sessionManager.Get("user", token)
		if !ok {
			return failAuthentication()
		}

		return resolver(p)
	}
}

func saveToSession(m *models.Member) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &LoginClaim{
		Key: m.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "shjp",
		},
	})

	ss, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Printf("Cannot sign token: %s\n", err)
		return "", errors.Wrap(err, "cannot sign token")
	}

	fmt.Printf("signed string is: %s\n", ss)

	sessionManager.Set("user", ss, m)
	return ss, nil
}

func failAuthentication() (interface{}, error) {
	return nil, errors.New("Authentication failed")
}
