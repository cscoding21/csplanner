package activity

import (
	"context"
	"csserver/internal/events"
	"csserver/internal/services/iam/appuser"
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
		GetDetail: func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
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

			//TOOD: enhance pubsub body to build activity output
			//---augment map with calculated values
			m["name"] = project.Data.ProjectBasics.Name
			m["comment"] = comment

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

			return fmt.Sprintf("/project/detail/%s#collab", projectID)
		},
	},
	"comment.reply.created": {
		Subject: "comment.reply.created",
		GetDetail: func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			payload := m["text"].(string)
			payloadDelta := quilljs.StringToDelta(payload)
			comment := quilljs.DeltaToString(*payloadDelta)
			projectID := m["project_id"].(string)
			parentUserID := m["parent_user_id"].(string)

			project, err := ps.GetProjectByID(ctx, projectID)
			if err != nil {
				log.Errorf("Activity template error GetProjectByID (comment.reply.created): %s", err)
				return ""
			}

			parentUser, err := us.GetAppuser(ctx, parentUserID)
			if err != nil {
				log.Errorf("Activity template error GetAppUserByID (comment.reply.created): %s", err)
				return ""
			}

			//TOOD: enhance pubsub body to build activity output
			//---augment map with calculated values
			m["project_name"] = project.Data.ProjectBasics.Name
			m["comment"] = comment
			m["parent_user_firstname"] = parentUser.FirstName
			m["parent_user_lastname"] = parentUser.LastName

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (comment.reply.created: %s", err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)
			projectID := m["project_id"].(string)

			return fmt.Sprintf("/project/detail/%s#collab", projectID)
		},
	},
	"resource.resource.created": {
		Subject: "resource.resource.created",
		GetDetail: func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

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
		GetDetail: func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

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
		GetDetail: func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

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
		GetDetail: func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

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
		GetDetail: func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

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
		GetDetail: func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

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
