package auth

import (
	"context"
	"errors"
	"fmt"

	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/providers/nats"
	userService "csserver/internal/services/iam/user"

	"github.com/Nerzal/gocloak/v13"

	log "github.com/sirupsen/logrus"
)

type IAMAdminService struct {
	KCClient      *gocloak.GoCloak
	KCRealm       string
	KCAdminUser   string
	KCAdminPass   string
	ContextHelper config.ContextHelper
	PubSub        nats.PubSubProvider
	UserService   userService.UserService
}

func NewIAMAdminService(
	ch config.ContextHelper,
	client *gocloak.GoCloak,
	pubsub nats.PubSubProvider,
	realm string,
	adminUser string,
	adminPass string,
	userService userService.UserService,
) IAMAdminService {
	out := IAMAdminService{}
	out.ContextHelper = ch
	out.KCClient = client
	out.KCRealm = realm
	out.KCAdminUser = adminUser
	out.KCAdminPass = adminPass
	out.PubSub = pubsub
	out.UserService = userService

	return out
}

// CreateUser create a new Keycloak user
func (s *IAMAdminService) CreateUser(ctx context.Context, user *userService.User) (common.UpdateResult[userService.User], error) {
	val := user.Validate()
	if !val.Pass {
		return common.NewUpdateResult[userService.User](&val, nil), fmt.Errorf("validation failed")
	}

	token, err := s.getAdminToken(ctx)
	if err != nil {
		return common.NewUpdateResult[userService.User](&val, nil), err
	}

	dbUser, err := s.UserService.UpsertUser(ctx, *user)
	if err != nil {
		log.Errorf("iamadmin createUser: %s", err)
		return common.NewUpdateResult[userService.User](&val, nil), err
	}

	if dbUser == nil {
		log.Error("iamadmin upsertUser returned nil")
		return common.NewUpdateResult[userService.User](&val, nil), err
	}

	props := make(map[string][]string)

	props["profileImage"] = []string{user.ProfileImage}
	props["dbid"] = []string{dbUser.Object.ID}

	u := gocloak.User{
		FirstName:  &user.FirstName,
		LastName:   &user.LastName,
		Email:      &user.Email,
		Enabled:    gocloak.BoolP(true),
		Username:   &user.Email,
		Attributes: &props,
	}

	uid, err := s.KCClient.CreateUser(ctx, token.AccessToken, s.KCRealm, u)
	if err != nil {
		return common.NewUpdateResult[userService.User](&val, nil), err
	}

	err = s.KCClient.SetPassword(ctx, token.AccessToken, uid, s.KCRealm, user.Password, false)
	if err != nil {
		return common.NewUpdateResult[userService.User](&val, nil), err
	}

	return common.NewUpdateResult[userService.User](&val, nil), err
}

// UpdateUser create a new Keycloak user
func (s *IAMAdminService) UpdateUser(ctx context.Context, user *userService.User) (common.UpdateResult[userService.User], error) {
	val := user.Validate()
	if !val.Pass {
		return common.NewUpdateResult[userService.User](&val, nil), fmt.Errorf("validation failed")
	}

	token, err := s.getAdminToken(ctx)
	if err != nil {
		return common.NewUpdateResult[userService.User](&val, nil), err
	}

	props := make(map[string][]string)

	props["profileImage"] = []string{user.ProfileImage}

	u := gocloak.User{
		FirstName:  &user.FirstName,
		LastName:   &user.LastName,
		Email:      &user.Email,
		Enabled:    gocloak.BoolP(true),
		Username:   &user.Email,
		Attributes: &props,
	}

	err = s.KCClient.UpdateUser(ctx, token.AccessToken, s.KCRealm, u)

	return common.NewUpdateResult[userService.User](&val, nil), err
}

// getAdminToken return an token that can be used for
func (s *IAMAdminService) getAdminToken(ctx context.Context) (*gocloak.JWT, error) {
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		"master")

	return common.HandleReturnWithValue(token, err)
}

// GetCurrentUser gets the user based on the credentials from the context
func (s *IAMAdminService) GetCurrentUser(ctx context.Context) (*userService.User, error) {
	contextUserEmail, ok := ctx.Value(config.UserEmailKey).(string)
	if !ok {
		return common.HandleReturnWithValue[userService.User](nil, errors.New("context does not contain user email"))
	}

	user, err := s.GetUser(ctx, contextUserEmail)
	return common.HandleReturnWithValue(user, err)
}

// GetUser get a user based on the passed in email
func (s *IAMAdminService) GetUser(ctx context.Context, idOrEmail string) (*userService.User, error) {
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		"master")
	if err != nil {
		return common.HandleReturnWithValue[userService.User](nil, err)
	}

	params := gocloak.GetUsersParams{
		Email: &idOrEmail,
	}

	users, err := s.KCClient.GetUsers(ctx, token.AccessToken, s.KCRealm, params)
	if err != nil || len(users) == 0 {
		return common.HandleReturnWithValue[userService.User](nil, err)
	}

	u := users[0]

	user := userService.User{
		FirstName:    *u.FirstName,
		LastName:     *u.LastName,
		Email:        *u.Email,
		ProfileImage: getProfileImage(u.Attributes),
		DBID:         getDBID(u.Attributes),
	}
	return common.HandleReturnWithValue(&user, err)
}

// GetUser get a user based on the passed in email
func (s *IAMAdminService) FindAllUsers(ctx context.Context) (common.PagedResults[userService.User], error) {
	pagingResults := common.NewPagedResultsForAllRecords[userService.User]()
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		"master")
	if err != nil {
		return pagingResults, err
	}

	params := gocloak.GetUsersParams{}

	users, err := s.KCClient.GetUsers(ctx, token.AccessToken, s.KCRealm, params)
	if err != nil || len(users) == 0 {
		return pagingResults, err
	}

	out := []userService.User{}

	for _, u := range users {
		user := userService.User{
			FirstName:    *u.FirstName,
			LastName:     *u.LastName,
			Email:        *u.Email,
			DBID:         getDBID(u.Attributes),
			ProfileImage: getProfileImage(u.Attributes),
		}

		out = append(out, user)
	}

	pagingResults.Results = out
	resultCount := len(pagingResults.Results)
	pagingResults.Pagination.TotalResults = &resultCount

	return pagingResults, nil
}

// getProfileImage get a profile image from the Keycloak attributes
func getProfileImage(attrs *map[string][]string) string {
	profileImg := ""

	if attrs != nil {
		attrs := *attrs
		st := attrs["profileImage"]
		if len(st) > 0 {
			profileImg = st[0]
		}
	}

	return profileImg
}

// getDBID get a profile image from the Keycloak attributes
func getDBID(attrs *map[string][]string) string {
	dbid := ""

	if attrs != nil {
		attrs := *attrs
		st := attrs["dbid"]
		if len(st) > 0 {
			dbid = st[0]
		}
	}

	return dbid
}