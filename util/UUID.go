package util

import (
	uuid "github.com/satori/go.uuid"
)

// IsValidUUID ID
func IsValidUUID(id uuid.UUID) bool {
	if id == uuid.Nil {
		return false
	}
	return true
}

// GenerateUUID Return New UUID
func GenerateUUID() uuid.UUID {
	return uuid.NewV4()
}
