package user

import (
	"context"
	"errors"

	"csserver/internal/common"
	"csserver/internal/config"

	"github.com/Nerzal/gocloak/v13"
)

type UserService struct {
	KCClient      *gocloak.GoCloak
	KCRealm       string
	KCAdminUser   string
	KCAdminPass   string
	ContextHelper config.ContextHelper
}

func NewUserService(
	ch config.ContextHelper,
	client *gocloak.GoCloak,
	realm string,
	adminUser string,
	adminPass string,
) UserService {
	out := UserService{}
	out.ContextHelper = ch
	out.KCClient = client
	out.KCRealm = realm
	out.KCAdminUser = adminUser
	out.KCAdminPass = adminPass

	return out
}

// NewUser create a new Keycloak user
func (s *UserService) NewUser(ctx context.Context, user *User) error {
	token, err := s.getAdminToken(ctx)
	if err != nil {
		return err
	}

	u := gocloak.User{
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
		Email:     &user.Email,
		Enabled:   gocloak.BoolP(true),
		Username:  &user.Email,
	}

	uid, err := s.KCClient.CreateUser(ctx, token.AccessToken, s.KCRealm, u)
	if err != nil {
		return err
	}

	err = s.KCClient.SetPassword(ctx, token.AccessToken, uid, s.KCRealm, user.Password, false)
	return common.HandleReturn(err)

}

// getAdminToken return an token that can be used for
func (s *UserService) getAdminToken(ctx context.Context) (*gocloak.JWT, error) {
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		"master")

	return common.HandleReturnWithValue(token, err)
}

// GetCurrentUser gets the user based on the credentials from the context
func (s *UserService) GetCurrentUser(ctx context.Context) (*User, error) {
	contextUserEmail, ok := ctx.Value(config.UserEmailKey).(string)
	if !ok {
		return common.HandleReturnWithValue[User](nil, errors.New("context does not contain user email"))
	}

	user, err := s.GetUser(ctx, contextUserEmail)
	return common.HandleReturnWithValue(user, err)
}

// GetUser get a user based on the passed in email
func (s *UserService) GetUser(ctx context.Context, idOrEmail string) (*User, error) {
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		"master")
	if err != nil {
		return common.HandleReturnWithValue[User](nil, err)
	}

	params := gocloak.GetUsersParams{
		Email: &idOrEmail,
	}

	users, err := s.KCClient.GetUsers(ctx, token.AccessToken, s.KCRealm, params)
	if err != nil || len(users) == 0 {
		return common.HandleReturnWithValue[User](nil, err)
	}

	u := users[0]

	user := User{
		FirstName: *u.FirstName,
		LastName:  *u.LastName,
		Email:     *u.Email,
		ID:        *u.ID,
	}
	return common.HandleReturnWithValue(&user, err)
}

// GetUser get a user based on the passed in email
func (s *UserService) FindAllUsers(ctx context.Context) (*[]User, error) {
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		"master")
	if err != nil {
		return common.HandleReturnWithValue[[]User](nil, err)
	}

	params := gocloak.GetUsersParams{}

	users, err := s.KCClient.GetUsers(ctx, token.AccessToken, s.KCRealm, params)
	if err != nil || len(users) == 0 {
		return common.HandleReturnWithValue[[]User](nil, err)
	}

	out := []User{}

	for _, u := range users {
		user := User{
			FirstName: *u.FirstName,
			LastName:  *u.LastName,
			Email:     *u.Email,
			ID:        *u.ID,
		}

		out = append(out, user)

	}
	return common.HandleReturnWithValue(&out, err)
}
