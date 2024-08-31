//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package comment

import (
	"csserver/internal/common"
)

type Comment struct {
	//---common for all DB objects
	common.ControlFields `json:"control_fields" csval:"validate"`

	//---TODO: add fields here
	ProjectID    string    `json:"project_id"`
	Text         string    `json:"text"`
	Replies      []Comment `json:"replies"`
	Likes        []string  `json:"likes"`
	Loves        []string  `json:"loves"`
	Dislikes     []string  `json:"dislikes"`
	LaughsAt     []string  `json:"laughs_at"`
	Acknowledges []string  `json:"acknowledges"`
	IsEdited     bool      `json:"is_edited"`
}

type Mention struct {
	UserID          string `json:"userId"`
	UserName        string `json:"userName"`
	CommentID       string `json:"commentId"`
	ProjectID       string `json:"projectId"`
	MentionedByID   string `json:"mentionedById"`
	MentionedByName string `json:"mentionedByName"`
	Text            string `json:"text"`
	IsBot           bool   `json:"isBot"`
}
