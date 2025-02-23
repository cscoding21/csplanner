package projecttemplate

import (
	"context"
	"csserver/internal/common"
)

// SaveTemplate create or update a project template
func (s *ProjecttemplateService) SaveTemplate(ctx context.Context, model Projecttemplate) (*common.UpdateResult[Projecttemplate], error) {

	return s.UpsertProjecttemplate(ctx, model)
}
