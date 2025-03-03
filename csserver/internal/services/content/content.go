package content

import (
	"context"
	"csserver/internal/providers/contentful"
	"fmt"
)

func NewContentService(provider contentful.ContentfulProvider) *ContentService {
	return &ContentService{
		Provider: provider,
	}
}

const DefaultLanguage = "en-US"

type ContentService struct {
	Provider contentful.ContentfulProvider
}

// GetContent retrieve a single bit of content based on its identifier
func (s *ContentService) GetContent(ctx context.Context, id string) (*Content, error) {
	entry, err := s.Provider.GetContent(ctx, id)
	if err != nil {
		return nil, err
	}

	c := Content{
		Title: fmt.Sprint(entry.Fields["title"].(map[string]any)[DefaultLanguage].(string)),
		Token: fmt.Sprint(entry.Fields["token"].(map[string]any)[DefaultLanguage].(string)),
		//Content:          fmt.Sprint(entry.Fields["content"].(map[string]any)[DefaultLanguage].(string)),
	}

	//ShortDescription: fmt.Sprint(entry.Fields["shortDescription"].(map[string]any)[DefaultLanguage].(string)),
	vurl, ok := entry.Fields["videoUrl"]
	if ok {
		c.VideoURL = fmt.Sprint(vurl.(map[string]any)[DefaultLanguage].(string))
	}

	sdes, ok := entry.Fields["shortDescription"]
	if ok {
		c.ShortDescription = fmt.Sprint(sdes.(map[string]any)[DefaultLanguage].(string))
	}

	// cont, ok := entry.Fields["content"]
	// if ok {
	// 	ic, ok := cont.(map[string]any)[DefaultLanguage]
	// 	if ok {
	// 		c.Content = ic.(string)
	// 	}
	// }

	return &c, nil
}
