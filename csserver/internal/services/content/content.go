package content

import (
	"context"
	"csserver/internal/providers/contentful"
)

func NewContentService(provider contentful.ContentfulProvider) *ContentService {
	return &ContentService{
		Provider: provider,
	}
}

type ContentService struct {
	Provider contentful.ContentfulProvider
}

// GetContent retrieve a single bit of content based on its identifier
func (s *ContentService) GetContent(ctx context.Context, id string) (interface{}, error) {
	return s.Provider.GetContent(ctx, id)
}
