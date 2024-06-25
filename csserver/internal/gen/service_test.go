package gen

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	props := GenProps{
		ServiceName: "Test",
		OutputPath:  "/home/jeph/projects/cscoding21/csplanner/csserver/internal/services",
	}

	err := GenService(props)
	if err != nil {
		t.Error(err)
	}
}

func TestGenerateList(t *testing.T) {
	props := GenProps{
		ServiceName: "List",
		OutputPath:  "/home/jeph/projects/cscoding21/csplanner/csserver/internal/services",
	}

	err := GenService(props)
	if err != nil {
		t.Error(err)
	}
}
