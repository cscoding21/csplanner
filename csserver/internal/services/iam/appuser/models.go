//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package appuser

import (
	"csserver/internal/common"
)

type Appuser struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	FirstName       string `json:"first_name" csval:"req"`
	LastName        string `json:"last_name" csval:"req"`
	Email           string `json:"email" csval:"req,email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	ProfileImage    string `json:"profile_image"`
	DBID            string `json:"dbid"`
}
