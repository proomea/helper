package auth_helper

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type AuthTokenClaimsCustomer struct {
	UserID     string `json:"user_id"`
	CustomerID string `json:"customer_id"`
	RoleID     string `json:"role_id"`
	jwt.StandardClaims
}

type AuthTokenClaimsEmployee struct {
	UserID string `json:"user_id"`
	OrgaID string `json:"orga_id""`
	RoleID string `json:"role_id"`
	jwt.StandardClaims
}

// JWT TOKEN for CUSTOMER LOGIN Endpoint
func (ah *authHelper) CreateTokenCustomer(userID, customerID, roleID string) (string, error) {
	claims := &AuthTokenClaimsCustomer{
		UserID:         userID,
		CustomerID:     customerID,
		RoleID:         roleID,
		StandardClaims: jwt.StandardClaims{},
	}
	ttl := ah.tokenExp * time.Second
	claims.StandardClaims.ExpiresAt = time.Now().UTC().Add(ttl).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(ah.jwtSecKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// JWT TOKEN for EMPLOYEE LOGIN Endpoint
func (ah *authHelper) CreateTokenEmployee(claims *AuthTokenClaimsEmployee) (string, error) {
	ttl := ah.tokenExp * time.Second
	claims.StandardClaims.ExpiresAt = time.Now().UTC().Add(ttl).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(ah.jwtSecKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Decode token for customer endpoints
// bearer token
func (ah *authHelper) DecodeTokenCustomer(token string) (string, string, string, error) {
	claims := &AuthTokenClaimsCustomer{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ah.jwtSecKey), nil
	})

	if claims, ok := t.Claims.(*AuthTokenClaimsCustomer); ok && t.Valid {
		//claims.UserClaims.OrgaID
		log.Printf("%v %v %v %v", claims.CustomerID, claims.UserID, claims.RoleID, claims.StandardClaims.ExpiresAt)
	} else {
		log.Println(err)
		return "", "", "", err
	}

	return claims.CustomerID, claims.UserID, claims.RoleID, nil
}

// Decode token for employee endpoints
// bearer token
func (ah *authHelper) DecodeTokenEmployee(token string) (string, string, string, error) {
	claims := &AuthTokenClaimsEmployee{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ah.jwtSecKey), nil
	})

	if claims, ok := t.Claims.(*AuthTokenClaimsEmployee); ok && t.Valid {
		//claims.UserClaims.OrgaID
		log.Printf("%v %v %v %v", claims.UserID, claims.UserID, claims.RoleID, claims.StandardClaims.ExpiresAt)
	} else {
		log.Println(err)
		return "", "", "", err
	}

	return claims.UserID, claims.UserID, claims.RoleID, nil
}
