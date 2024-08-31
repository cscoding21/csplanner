package utils

import (
	"testing"
)

func TestRefFromValue(t *testing.T) {
	val := 1
	ref := ValToRef(val)

	if ref == nil {
		t.Fatalf(`ref value should be address to "%v"`, val)
	}

	if *ref != 1 {
		t.Fatalf(`ref value did not resolve to correct value "%v"`, *ref)
	}
}
