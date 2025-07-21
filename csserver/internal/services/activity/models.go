//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package activity

import (
	"csserver/internal/common"
	"time"
)

type ActivityType string

const (
	ProjectProjectCreated   = "project.project.created"
	ProjectProjectUpdated   = "project.project.updated"
	ProjectStateUpdated     = "project.state.updated"
	ProjectFeatureCreated   = "project.feature.created"
	ProjectFeatureUpdated   = "project.feature.updated"
	ProjectTaskCreated      = "project.task.created"
	ProjectTaskUpdated      = "project.task.updated"
	ProjectMilestoneCreated = "project.milestone.created"
	ProjectMilestoneUpdated = "project.milestone.updated"
	ResourceRoleUpdated     = "resource.role.updated"
	ResourceRoleCreated     = "resource.role.created"
	ResourceResourceCreated = "resource.resource.created"
	CommentReplyCreated     = "comment.reply.created"
	CommentCommentCreated   = "comment.comment.created"
	CommentReactionCreated  = "comment.reaction.created"
)

type Activity struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	Detail       string       `json:"detail"`
	Context      ActivityType `json:"context"`
	Link         string       `json:"link"`
	UserEmail    string       `json:"user_email"`
	ActivityDate time.Time    `json:"activity_date"`
}
