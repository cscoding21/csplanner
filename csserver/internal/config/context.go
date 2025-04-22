package config

import (
	"context"
)

type contextKey string

const (
	UserEmailKey contextKey = "userEmail"
	UserIDKey    contextKey = "userID"
	OrgUrlKey    contextKey = "orgUrlKey"
)

// NewContext create a new context
func NewContext() context.Context {
	ctx := context.Background()

	return ctx
}

func getValueFromContext(ctx context.Context, key contextKey) string {
	val := ctx.Value(key)

	if val != nil {
		return val.(string)
	}

	return ""
}

// GetUserIDFromContext get user id from context
func GetUserIDFromContext(ctx context.Context) string {
	return getValueFromContext(ctx, UserIDKey)
}

// GetUserEmailFromContext get user id from context
func GetUserEmailFromContext(ctx context.Context) string {
	return getValueFromContext(ctx, UserEmailKey)
}

// GetOrgUrlKeyFromContext get org url key from context
func GetOrgUrlKeyFromContext(ctx context.Context) string {
	return getValueFromContext(ctx, OrgUrlKey)
}
