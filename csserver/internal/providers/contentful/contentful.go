package contentful

import (
	"context"
	"fmt"

	"github.com/contentful-labs/contentful-go"
)

type ContentfulProvider struct {
	OrgID   string
	SpaceID string
	PAT     string
}

const DefaultLanguage = "en-US"

func (c *ContentfulProvider) GetContent(ctx context.Context, id string) (*string, error) {
	cma := contentful.NewCMA(c.PAT)

	cma.SetOrganization(c.OrgID)
	cma.Debug = true

	item, err := cma.Entries.Get(c.SpaceID, id)
	if err != nil {
		return nil, err
	}

	field := item.Fields["title"].(map[string]interface{})
	out := fmt.Sprint(field[DefaultLanguage].(string))

	return &out, nil
}
