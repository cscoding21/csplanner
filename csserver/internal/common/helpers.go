package common

import (
	"encoding/base64"

	"github.com/google/uuid"
)

// IsOneOf test to see if the input is equal to a list of items
func IsOneOf[T comparable](input T, opts ...T) bool {
	for _, o := range opts {
		if input == o {
			return true
		}
	}

	return false
}

// Coalesce return the first non-nil value in the argument array
func Coalesce[T any](opts ...*T) *T {
	for _, o := range opts {
		if o != nil {
			return o
		}
	}

	return nil
}

// GetDBID create a base64 guid for DB ids
func GetDBID() string {
	uuid := []byte(uuid.New().String())
	return base64.RawURLEncoding.EncodeToString(uuid)
}
