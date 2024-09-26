//go:generate csval

package user

import "csserver/internal/common"

type User struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	FirstName       string `json:"firstName" csval:"req"`
	LastName        string `json:"lastName" csval:"req"`
	Email           string `json:"email" csval:"req,email"`
	Password        string `json:"password" csval:"req"`
	ConfirmPassword string `json:"confirmPassword" csval:"req,equals(Password)"`
	ProfileImage    string `json:"profileImage"`
	DBID            string `json:"dbid"`
}
