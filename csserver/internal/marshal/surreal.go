package marshal

import (
	"github.com/surrealdb/surrealdb.go"
)

// SurrealUnmarshal warpper for surrealdb.Unmarshal
func SurrealUnmarshal[T any](data interface{}) (*T, error) {
	var o T
	err := surrealdb.Unmarshal(data, &o)

	return &o, err
}

// SurrealUnmarshalRaw warpper for surrealdb.UnmarshalRaw
func SurrealUnmarshalRaw[T any](data interface{}, object interface{}) (*T, error) {
	var o T
	_, err := surrealdb.UnmarshalRaw(data, &o)

	return &o, err
}

// SurrealSmartUnmarshal warpper for surrealdb.SmartUnmarshal
func SurrealSmartUnmarshal[T any](data interface{}) (*T, error) {
	var err error
	out, err := surrealdb.SmartUnmarshal[T](data, err)

	return &out, err
}
