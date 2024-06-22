package interfaces

import (
	"context"
)

type ContextHelpers interface {
	GetUserIDFromContext(ctx context.Context) string
	GetUserEmailFromContext(ctx context.Context) string
}
