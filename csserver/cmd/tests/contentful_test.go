package tests

import (
	"csserver/internal/config"
)

func init() {
	config.InitConfig()
}

/*
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

		if !strings.EqualFold(con.Title, expectedValue) {
			t.Errorf("unexpected value returned: expected '%s' - got '%s'", expectedValue, con.Title)
		}

		fmt.Printf("Title: %s\n", con.Title)
		fmt.Printf("Content: %s\n", con.Content)
		fmt.Printf("Short Description: %s\n", con.ShortDescription)
		fmt.Printf("VideoURL: %s\n", con.VideoURL)
		fmt.Printf("Token: %s\n", con.Token)

	}
}
*/
