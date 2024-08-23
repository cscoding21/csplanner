package common

import (
	log "github.com/sirupsen/logrus"
)

// HandleReturn passthrough function that logs an error if needed.
func HandleReturn(err error) error {
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// HandleReturnWithValue passthrough function that logs an error and value if needed.
func HandleReturnWithValue[T any](ret *T, err error) (*T, error) {
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return ret, nil
}
