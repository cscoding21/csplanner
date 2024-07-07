//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package user

import (
	"csserver/internal/common"
)

type User struct {
	//---common for all DB objects
	common.ControlFields

	//---TODO: add fields here
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
