package activity

import (
	"context"
	"csserver/internal/events"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"csserver/internal/utils/quilljs"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

var templateMap = map[string]ActivityTemplate{
	"comment.comment.created": {
		Subject: "comment.comment.created",
		GetDetail: func(ctx context.Context, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			payload := m["text"].(string)
			payloadDelta := quilljs.StringToDelta(payload)
			comment := quilljs.DeltaToString(*payloadDelta)
			projectID := m["project_id"].(string)

			project, err := ps.GetProjectByID(ctx, projectID)
			if err != nil {
				log.Errorf("Activity template error (comment.comment.created): %s", err)
				return ""
			}

			mentionEmbed := quilljs.NewMentionEmbed(wrapper.UserID, "Jeph")
			delta := quilljs.New(nil).
				InsertEmbed(mentionEmbed,
					map[string]any{
						"bold": true,
					}).
				Insert(" commented on the project ", nil).
				Insert(project.Data.ProjectBasics.Name, map[string]any{"bold": true}).
				Insert("\n", nil).
				Insert(comment, map[string]any{"blockquote": true})

			out, err := json.Marshal(delta)
			if err != nil {
				log.Errorf("Activity template error (comment.comment.created): %s", err)
				return ""
			}

			return string(out)
		},
	},
	"resource.resource.created": {
		Subject: "resource.resource.created",
		GetDetail: func(ctx context.Context, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			fmt.Println(m)

			//TOOD: enhance delta to build activity output
			delta := quilljs.New(nil)

			out, err := json.Marshal(delta)
			if err != nil {
				log.Errorf("Activity template error (resource.resource.created): %s", err)
				return ""
			}

			return string(out)
		},
	},
}
