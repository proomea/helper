package auth_helper

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type AuthTokenClaims struct {
	UserID string `json:"user_id"`
	OrgaID string `json:"orga_id""`
	RoleID string `json:"role_id"`
	jwt.StandardClaims
}

// JWT TOKEN creator for LOGIN Endpoint
func (ah *authHelper) CreateToken(claims *AuthTokenClaims) (string, error) {
	ttl := ah.tokenExp * time.Second
	claims.StandardClaims.ExpiresAt = time.Now().UTC().Add(ttl).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(ah.jwtSecKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Decode bearer token
// Extract userid, orgaid and roleid out of token claims
func (ah *authHelper) DecodeToken(token string) (string, string, string, error) {
	claims := &AuthTokenClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ah.jwtSecKey), nil
	})

	if claims, ok := t.Claims.(*AuthTokenClaims); ok && t.Valid {
		//claims.UserClaims.OrgaID
		log.Printf("%v %v %v %v", claims.OrgaID, claims.UserID, claims.RoleID, claims.StandardClaims.ExpiresAt)
	} else {
		log.Println(err)
		return "", "", "", err
	}

	return claims.OrgaID, claims.UserID, claims.RoleID, nil
}
