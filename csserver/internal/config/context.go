package config

import (
	"context"
)

type contextKey string
type ContextHelper struct{}

const (
	UserEmailKey contextKey = "userEmail"
	UserIDKey    contextKey = "userID"
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
func (h *ContextHelper) GetUserIDFromContext(ctx context.Context) string {
	return getValueFromContext(ctx, UserIDKey)
}

// GetUserEmailFromContext get user id from context
func (h *ContextHelper) GetUserEmailFromContext(ctx context.Context) string {
	return getValueFromContext(ctx, UserEmailKey)
}
