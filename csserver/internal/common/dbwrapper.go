package common

import (
	"time"
)

// BaseModel wrapper for all Posgtres database tables
type BaseModel[T any] struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"craeted_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy string     `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *string    `json:"deleted_by"`
	Data      T          `json:"data"`
}

// ExtractDataFromBase return a data property from a base result
func ExtractDataFromBase[T any](results []BaseModel[T]) []T {
	out := []T{}

	for _, r := range results {
		out = append(out, r.Data)
	}

	return out
}

// Base contains common columns for all tables.
type ControlFields struct {
	ID string `json:"id,omitempty" csval:"req"`
}
