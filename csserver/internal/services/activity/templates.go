package activity

import (
	"context"
	"csserver/internal/events"
	"csserver/internal/services/comment"
	"csserver/internal/services/iam/appuser"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"csserver/internal/utils/quilljs"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

var templateMap = map[ActivityType]ActivityTemplate{
	CommentCommentCreated: {
		Subject: CommentCommentCreated,
		GetDetail: func(ctx context.Context, cs comment.CommentService, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			payload := m["text"].(string)
			payloadDelta := quilljs.StringToDelta(payload)
			comment := quilljs.DeltaToString(*payloadDelta)
			projectID := m["project_id"].(string)

			project, err := ps.GetProjectByID(ctx, projectID)
			if err != nil {
				log.Errorf("Activity template error (%s): %s", CommentCommentCreated, err)
				return ""
			}

			//TOOD: enhance pubsub body to build activity output
			//---augment map with calculated values
			m["name"] = project.Data.ProjectBasics.Name
			m["comment"] = comment

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (%s): %s", CommentCommentCreated, err)
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
	CommentReplyCreated: {
		Subject: CommentReplyCreated,
		GetDetail: func(ctx context.Context, cs comment.CommentService, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			payload := m["text"].(string)
			payloadDelta := quilljs.StringToDelta(payload)
			comment := quilljs.DeltaToString(*payloadDelta)
			projectID := m["project_id"].(string)
			parentUserID := m["parent_user_id"].(string)

			project, err := ps.GetProjectByID(ctx, projectID)
			if err != nil {
				log.Errorf("Activity template error GetProjectByID (%s): %s", CommentReplyCreated, err)
				return ""
			}

			parentUser, err := us.GetAppuser(ctx, parentUserID)
			if err != nil {
				log.Errorf("Activity template error GetAppUserByID (%s): %s", CommentReplyCreated, err)
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
				log.Errorf("Activity template error (%s: %s", CommentReplyCreated, err)
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
	ResourceResourceCreated: {
		Subject: ResourceResourceCreated,
		GetDetail: func(ctx context.Context, cs comment.CommentService, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (%s): %s", ResourceResourceCreated, err)
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
	ResourceRoleCreated: {
		Subject: ResourceRoleCreated,
		GetDetail: func(ctx context.Context, cs comment.CommentService, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (%s): %s", ResourceRoleCreated, err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			return "/settings#roles"
		},
	},
	ResourceRoleUpdated: {
		Subject: ResourceRoleUpdated,
		GetDetail: func(ctx context.Context, cs comment.CommentService, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (%s): %s", ResourceRoleUpdated, err)
				return ""
			}

			return string(out)
		},
		GetLink: func(wrapper events.MessageWrapper) string {
			return "/settings#roles"
		},
	},
	ProjectProjectCreated: {
		Subject: ProjectProjectCreated,
		GetDetail: func(ctx context.Context, cs comment.CommentService, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (%s): %s", ProjectProjectCreated, err)
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
	ProjectProjectUpdated: {
		Subject: ProjectProjectUpdated,
		GetDetail: func(ctx context.Context, cs comment.CommentService, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (%s): %s", ProjectProjectUpdated, err)
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
	ProjectStateUpdated: {
		Subject: ProjectStateUpdated,
		GetDetail: func(ctx context.Context, cs comment.CommentService, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string {
			m := wrapper.Body.(map[string]any)

			//TOOD: enhance pubsub body to build activity output

			out, err := json.Marshal(m)
			if err != nil {
				log.Errorf("Activity template error (%s): %s", ProjectStateUpdated, err)
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
