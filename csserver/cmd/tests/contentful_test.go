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
	cf := factory.GetContentService()

	ctx := getTestContext()

	expectedValue := "Organization Settings"
	contentID := "1TlbGHQaDfbZ1boVEOs4zi"

	con, err := cf.GetContent(ctx, contentID)
	if err != nil {
		t.Error(err)
	}

	if con == nil {
		t.Error("name returned as nil value")
	} else {
		name := con.(*string)

		if !strings.EqualFold(*name, expectedValue) {
			t.Errorf("unexpected value returned: expected '%s' - got '%s'", expectedValue, *name)
		}

		fmt.Println(name)
	}
}
