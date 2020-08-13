package auth_helper

import (
	uuid "github.com/satori/go.uuid"
)

// Create an uuid
func (ah *authHelper) CreateUUID() string {
	return uuid.NewV1().String()
}
