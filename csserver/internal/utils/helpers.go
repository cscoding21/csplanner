package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/wI2L/jsondiff"
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

// GeneratePassword create a secure password
func GeneratePassword() string {
	u := uuid.NewString()

	return u
}

// GenerateBase64Guid create a uuid that is compressed to 22 characters.  Suitable for ID generateion
func GenerateBase64UUID() string {
	guid := uuid.New()
	encodedGuid := base64.RawURLEncoding.EncodeToString(guid[:])
	return encodedGuid
}

// GetDiffGraph return a list of differences between 2 objects
func GetDiffGraph[T any](source T, target T) string {
	sourceJSON, err := json.Marshal(source)
	if err != nil {
		return ""
	}

	targetJSON, err := json.Marshal(target)
	if err != nil {
		return ""
	}

	diffs, err := jsondiff.CompareJSON(sourceJSON, targetJSON)
	if err != nil {
		return ""
	}

	diffOut, err := json.Marshal(diffs)
	if err != nil {
		return ""
	}

	return string(diffOut)
}
