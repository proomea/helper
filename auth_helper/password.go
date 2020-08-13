package auth_helper

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// Hash password
func (ah *authHelper) CreatePasswordHash(pw string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		log.Println("error while hashing password")
		return "", err
	}
	return string(res), nil

}

// Check password
func (ah *authHelper) AuthPassword(pwStored string, pwInput string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwStored), []byte(pwInput))
	if err != nil {
		log.Println("error while checking password")
		return false
	}
	return true
}
