package common

import (
	"fmt"

	"github.com/cscoding21/csval/validate"
)

type UpdateResult[T any] struct {
	Success          bool
	ValidationResult validate.ValidationResult
	Object           *T
}

func UpwrapFromUpdateResult[T any](r UpdateResult[*BaseModel[T]]) *T {
	if r.Object != nil {
		obj := *r.Object
		return &obj.Data
	}

	return nil
}

// NewUpdateResult return a result of an object update with validation information
func NewUpdateResult[T any](val validate.ValidationResult, obj *T) UpdateResult[T] {
	ur := UpdateResult[T]{
		ValidationResult: val,
	}

	if obj != nil {
		ur.Object = obj
	}

	ur.Success = val.Pass

	return ur
}

// NewSuccessUpdateResult return a new successful result while handling the validation
func NewSuccessUpdateResult[T any](obj T) (UpdateResult[T], error) {
	val := validate.NewSuccessValidationResult()

	return NewUpdateResult(val, &obj), nil
}

// NewFailingUpdateResult return a failing update result and error
func NewFailingUpdateResult[T any](obj *T, err error) (UpdateResult[T], error) {
	return UpdateResult[T]{
		ValidationResult: validate.ValidationResult{},
		Success:          false,
		Object:           obj,
	}, err
}

// NewFailingUpdateResultWithVal return a hydrated update results
func NewFailingUpdateResultWithVal[T any](obj *T, val validate.ValidationResult) (UpdateResult[T], error) {
	return UpdateResult[T]{
		ValidationResult: val,
		Success:          false,
		Object:           obj,
	}, fmt.Errorf("validation failed for object")
}

// MergeUpdateResults
func MergeUpdateResults[T any](results []UpdateResult[T]) (UpdateResult[T], error) {
	if len(results) == 0 {
		return NewFailingUpdateResult[T](nil, fmt.Errorf("no results to aggregate"))
	}

	result, err := NewSuccessUpdateResult(*results[0].Object)
	if err != nil {
		return NewFailingUpdateResult[T](nil, err)
	}

	for _, r := range results {
		if !r.Success {
			result.Success = false
		}

		result.ValidationResult.Append(r.ValidationResult)
	}

	return result, nil
}
