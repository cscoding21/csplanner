//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package notification

import (
	"csserver/internal/common"
	"csserver/internal/services/notification/ntypes"
)

type Notification struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	UserEmail             string                  `json:"user_email"`
	UserName              string                  `json:"user_name"`
	RecipientIsBot        bool                    `json:"recipient_is_bot"`
	Type                  ntypes.NotificationType `json:"type"`
	ContextID             string                  `json:"context_id"`
	Text                  string                  `json:"text"`
	IsRead                bool                    `json:"is_read"`
	InitiatorName         string                  `json:"initiator_name"`
	InitiatorEmail        string                  `json:"initiator_email"`
	InitiatorProfileImage *string                 `json:"initiator_profile_image"`
}
