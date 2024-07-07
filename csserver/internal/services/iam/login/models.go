//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package login

import (
	"csserver/internal/common"
)

type Login struct {
	//---common for all DB objects
	common.ControlFields

	//---TODO: add fields here
	UserID   string `json:"user_id" csval:"req"`
	Username string `json:"username" csval:"req"`
	Password string `json:"password" csval:"req"`
}
