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

			//---augment map with calculated values
			m["name"] = project.Data.ProjectBasics.Name
			m["comment"] = comment

			// mentionEmbed := quilljs.NewMentionEmbed(wrapper.UserID, "Jeph")
			// delta := quilljs.New(nil).
			// 	InsertEmbed(mentionEmbed,
			// 		map[string]any{
			// 			"bold": true,
			// 		}).
			// 	Insert(" commented on the project ", nil).
			// 	Insert(project.Data.ProjectBasics.Name, map[string]any{"bold": true}).
			// 	Insert("\n", nil).
			// 	Insert(comment, map[string]any{"small": true})

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (comment.comment.created): %s", err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			projectID := m["project_id"].(string)

			return fmt.Sprintf("/project/detail/%s", projectID)
		},
	},
	"resource.resource.created": {
		Subject: "resource.resource.created",
		GetDetail: func(ctx context.Context, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			fmt.Println(m)

			//TOOD: enhance delta to build activity output
			// delta := quilljs.New(nil).
			// 	Insert(m["name"].(string), map[string]any{"bold": true}).
			// 	Insert(" was added to the resource roster.", nil)

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (resource.resource.created): %s", err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			resourceID := m["id"].(string)

			return fmt.Sprintf("/resource/detail/%s", resourceID)
		},
	},
	"resource.role.created": {
		Subject: "resource.role.created",
		GetDetail: func(ctx context.Context, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			fmt.Println(m)

			//TOOD: enhance delta to build activity output
			// delta := quilljs.New(nil).
			// 	Insert("A new role called ", nil).
			// 	Insert(m["name"].(string), map[string]any{"bold": true}).
			// 	Insert(" was created.", nil)

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (resource.role.created): %s", err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			return "/settings#roles"
		},
	},
	"resource.role.updated": {
		Subject: "resource.role.updated",
		GetDetail: func(ctx context.Context, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			fmt.Println(m)

			//TOOD: enhance delta to build activity output
			// delta := quilljs.New(nil).
			// 	Insert("Details of the role ", nil).
			// 	Insert(m["name"].(string), map[string]any{"bold": true}).
			// 	Insert(" were updated.", nil)

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (resource.role.updated): %s", err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			return "/settings#roles"
		},
	},
	"project.project.created": {
		Subject: "project.project.created",
		GetDetail: func(ctx context.Context, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			fmt.Println(m)

			//TOOD: enhance delta to build activity output
			// delta := quilljs.New(nil).
			// 	Insert("A new project,  ", nil).
			// 	Insert(m["name"].(string), map[string]any{"bold": true}).
			// 	Insert(", was added to the portfolio!", nil)

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (project.project.created): %s", err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			projectID := m["id"].(string)

			return fmt.Sprintf("/project/detail/%s", projectID)
		},
	},
	"project.project.updated": {
		Subject: "project.project.updated",
		GetDetail: func(ctx context.Context, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			fmt.Println(m)

			//TOOD: enhance delta to build activity output
			// delta := quilljs.New(nil).
			// 	Insert("Details for the project ", nil).
			// 	Insert(m["name"].(string), map[string]any{"bold": true}).
			// 	Insert(" were updated.", nil)

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (project.project.updated): %s", err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			projectID := m["id"].(string)

			return fmt.Sprintf("/project/detail/%s", projectID)
		},
	},
	"project.state.updated": {
		Subject: "project.state.updated",
		GetDetail: func(ctx context.Context, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			fmt.Println(m)

			//TOOD: enhance delta to build activity output
			// delta := quilljs.New(nil).
			// 	Insert("Project  ", nil).
			// 	Insert(m["name"].(string), map[string]any{"bold": true}).
			// 	Insert(" was moved from ", nil).
			// 	Insert(m["from_state"].(string), map[string]any{"bold": true}).
			// 	Insert(" to ", nil).
			// 	Insert(m["to_state"].(string), map[string]any{"bold": true})

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (project.state.updated): %s", err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			projectID := m["id"].(string)

			return fmt.Sprintf("/project/detail/%s", projectID)
		},
	},
}
