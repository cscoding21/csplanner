//go:generate csval

package user

type User struct {
	ID              string `json:"id,omitempty"`
	FirstName       string `json:"firstName" csval:"req"`
	LastName        string `json:"lastName" csval:"req"`
	Email           string `json:"email" csval:"req,email"`
	Password        string `json:"password" csval:"req"`
	ConfirmPassword string `json:"confirmPassword" csval:"req,equals(Password)"`
	ProfileImage    string `json:"profileImage"`
}
