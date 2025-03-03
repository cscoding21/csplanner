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

func (c *ContentfulProvider) GetContent(ctx context.Context, id string) (*contentful.Entry, error) {
	cma := contentful.NewCMA(c.PAT)

	cma.SetOrganization(c.OrgID)
	cma.Debug = true

	item, err := cma.Entries.Get(c.SpaceID, id)
	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, fmt.Errorf("contentful error for id %s was returned as nil value", id)
	}

	return item, nil
}
