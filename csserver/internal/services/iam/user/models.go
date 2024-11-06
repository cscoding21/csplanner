//go:generate csval

package user

import "csserver/internal/common"

type User struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	FirstName       string `json:"first_name" csval:"req"`
	LastName        string `json:"last_name" csval:"req"`
	Email           string `json:"email" csval:"req,email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	ProfileImage    string `json:"profile_image"`
	DBID            string `json:"dbid"`
}
