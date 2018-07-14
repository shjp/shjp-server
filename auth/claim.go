package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

// LoginClaim is JWT claim used for user auth
type LoginClaim struct {
	Key string `json:"key"`
	jwt.StandardClaims
}

// Valid defines whether the LoginClaim is valid
func (c *LoginClaim) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return errors.Wrap(err, "invalid standard claims")
	}
	return nil
}
