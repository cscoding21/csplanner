package common

import (
	"github.com/cscoding21/csval/validate"
)

type UpdateResult[T any] struct {
	ValidationResult *validate.ValidationResult
	Object           *T
}

// NewUpdateResult return a result of an object update with validation information
func NewUpdateResult[T any](val *validate.ValidationResult, obj *T) UpdateResult[T] {
	ur := UpdateResult[T]{
		ValidationResult: val,
	}

	if obj != nil {
		ur.Object = obj
	}

	return ur
}
