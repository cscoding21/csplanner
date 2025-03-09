package projecttemplate

import (
	"context"
	"csserver/internal/common"

	"github.com/google/uuid"
)

// SaveTemplate create or update a project template
func (s *ProjecttemplateService) SaveTemplate(ctx context.Context, model Projecttemplate) (common.UpdateResult[*common.BaseModel[Projecttemplate]], error) {
	//---make sure all child item have IDs
	for i, p := range model.Phases {
		if len(p.ID) == 0 {
			model.Phases[i].ID = uuid.New().String()
		}

		for j, t := range p.Tasks {
			if len(t.ID) == 0 {
				model.Phases[i].Tasks[j].ID = uuid.New().String()
			}
		}
	}

	return s.UpsertProjecttemplate(ctx, model)
}
