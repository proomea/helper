package auth_helper

import (
	"time"
)

type authHelper struct {
	jwtSecKey string
	tokenExp  time.Duration
}

type AuthHelperInterface interface {
	AuthPassword(pwStored string, pwInput string) bool
	CreatePasswordHash(pw string) (string, error)
	CreateToken(claims *AuthTokenClaims) (string, error)
	DecodeToken(token string) (string, string, string, error)
	CreateUUID() string
}

func NewAuthLib(jwtSecKey string, tokenExp time.Duration) AuthHelperInterface {
	return &authHelper{
		jwtSecKey: jwtSecKey,
		tokenExp:  tokenExp,
	}
}
