//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package comment

import (
	"csserver/internal/common"
)

// type CommentRelationshipType string
type CommentReactionType string

type Comment struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	ProjectID      string                      `json:"project_id"`
	Text           string                      `json:"text"`
	CommentReplies []common.BaseModel[Comment] `json:"replies"`
	Likes          []string                    `json:"likes"`
	Loves          []string                    `json:"loves"`
	Dislikes       []string                    `json:"dislikes"`
	LaughsAt       []string                    `json:"laughs_at"`
	Acknowledges   []string                    `json:"acknowledges"`
	IsEdited       bool                        `json:"is_edited"`
}

type Reaction struct {
	ProjectID string              `json:"project_id"`
	Type      CommentReactionType `json:"type"`
}

type Mention struct {
	UserEmail       string `json:"user_email"`
	UserName        string `json:"user_name"`
	CommentID       string `json:"comment_ie"`
	ProjectID       string `json:"project_id"`
	MentionedByID   string `json:"mentioned_by_id"`
	MentionedByName string `json:"mentioned_by_name"`
	Text            string `json:"text"`
	IsBot           bool   `json:"is_bot"`
}
