package common

import "github.com/cscoding21/csval/validate"

type UpdateResult[T any] struct {
	ValidationResult *validate.ValidationResult
	Object           *T
}

func NewUpdateResult[T any](val *validate.ValidationResult, obj *T) UpdateResult[T] {
	ur := UpdateResult[T]{
		ValidationResult: val,
		Object:           obj,
	}

	return ur
}
