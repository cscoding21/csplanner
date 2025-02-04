package utils

import (
	"encoding/json"
	"strings"
)

// RemoveAll removes all occurrences of items from input.
func RemoveAll(input string, items ...string) string {
	out := input
	for _, item := range items {
		out = strings.Replace(out, item, "", -1)
	}

	return out
}

// DeepCopy returns a deep copy of the passed in object provided it is serializable.
func DeepCopy[T any](input T) (*T, error) {
	origJSON, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var clone T
	if err = json.Unmarshal(origJSON, &clone); err != nil {
		return nil, err
	}

	return &clone, nil
}
