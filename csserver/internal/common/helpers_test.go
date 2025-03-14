package common

import (
	"fmt"
	"testing"
)

func TestGetDBID(t *testing.T) {
	id := GetDBID()

	if len(id) == 0 {
		t.Errorf("id did not return a valid result: %s", id)
	}

	fmt.Println(id)
}
