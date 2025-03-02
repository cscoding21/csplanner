package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"fmt"
	"strings"
	"testing"
)

func init() {
	config.InitConfig()
}

func TestGetContent(t *testing.T) {
	cf, err := factory.GetContentfulClient()
	if err != nil {
		t.Error(err)
	}

	ctx := getTestContext()

	expectedValue := "Organization Settings"
	contentID := "1TlbGHQaDfbZ1boVEOs4zi"

	name, err := cf.GetContent(ctx, contentID)
	if err != nil {
		t.Error(err)
	}

	if name == nil {
		t.Error("name returned as nil value")
	} else {
		if !strings.EqualFold(*name, expectedValue) {
			t.Errorf("unexpected value returned: expected '%s' - got '%s'", expectedValue, *name)
		}

		fmt.Println(name)
	}
}
