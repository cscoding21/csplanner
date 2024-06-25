package gen

import (
	"testing"
)

func TestGenServices(t *testing.T) {
	err := GenServices("/home/jeph/projects/cscoding21/csplanner/csserver/cmd/cstools/.cstools.yaml")
	if err != nil {
		t.Error(err)
	}
}
