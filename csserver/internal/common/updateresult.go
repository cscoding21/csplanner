package common

import (
	"github.com/cscoding21/csval/validate"
)

type UpdateResult[T any] struct {
	Success          bool
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

	ur.Success = val.Pass

	return ur
}

// NewFailingUpdateResult return a failing update result and error
func NewFailingUpdateResult[T any](obj *T, err error) (UpdateResult[T], error) {
	return UpdateResult[T]{
		ValidationResult: &validate.ValidationResult{},
		Success:          false,
		Object:           obj,
	}, err

}
