package common

import (
	"csserver/internal/interfaces"
)

// HandleReturn passthrough function that logs an error if needed.
func HandleReturn(
	l interfaces.Logger,
	err error) error {

	if err != nil {
		l.Error(err)
		return err
	}

	return nil
}

// HandleReturnWithValue passthrough function that logs an error and value if needed.
func HandleReturnWithValue[T any](l interfaces.Logger, ret *T, err error) (*T, error) {
	if err != nil {
		l.Error(err)
		return nil, err
	}

	return ret, nil
}
