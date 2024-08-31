package utils

import (
	"testing"
)

func TestGetTimeFromString(t *testing.T) {
	have := "2023-11-21"

	got := GetTimeFromString(have)

	if got.Year() != 2023 || got.Month() != 11 || got.Day() != 21 {
		t.Fatalf(`Time did not resolve correctly : want %s - got "%v"`, have, got)
	}
}
