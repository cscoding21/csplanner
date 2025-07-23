//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package activity

import (
	"csserver/internal/common"
	"fmt"
	"time"
)

type ActivityType string

const (
	ActivityServiceProject  = "project"
	ActivityServiceResource = "resource"
	ActivityServiceComment  = "comment"

	ActivityObjectProject   = "project"
	ActivityObjectState     = "state"
	ActivityObjectFeature   = "feature"
	ActivityObjectTask      = "task"
	ActivityObjectMilestone = "milestone"

	ActivityObjectResource = "resource"
	ActivityObjectRole     = "role"

	ActivityObjectComment  = "comment"
	ActivityObjectReply    = "reply"
	ActivityObjectReaction = "reaction"

	ActivityEventCreated = "created"
	ActivityEventUpdated = "updated"
	ActivityEventDeleted = "deleted"
)

var (
	ProjectProjectCreated   = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectProject, ActivityEventCreated))
	ProjectProjectUpdated   = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectProject, ActivityEventUpdated))
	ProjectStateUpdated     = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectState, ActivityEventUpdated))
	ProjectFeatureCreated   = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectFeature, ActivityEventCreated))
	ProjectFeatureUpdated   = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectFeature, ActivityEventUpdated))
	ProjectTaskCreated      = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectTask, ActivityEventCreated))
	ProjectTaskUpdated      = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectTask, ActivityEventUpdated))
	ProjectMilestoneCreated = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectMilestone, ActivityEventCreated))
	ProjectMilestoneUpdated = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceProject, ActivityObjectMilestone, ActivityEventUpdated))

	ResourceRoleUpdated     = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceResource, ActivityObjectRole, ActivityEventUpdated))
	ResourceRoleCreated     = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceResource, ActivityObjectRole, ActivityEventCreated))
	ResourceResourceCreated = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceResource, ActivityObjectResource, ActivityEventCreated))
	ResourceResourceUpdated = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceResource, ActivityObjectResource, ActivityEventUpdated))

	CommentReplyCreated    = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceComment, ActivityObjectReply, ActivityEventCreated))
	CommentCommentCreated  = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceComment, ActivityObjectComment, ActivityEventCreated))
	CommentCommentUpdated  = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceComment, ActivityObjectComment, ActivityEventUpdated))
	CommentReactionCreated = ActivityType(fmt.Sprintf("%s.%s.%s", ActivityServiceComment, ActivityObjectReaction, ActivityEventCreated))
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
