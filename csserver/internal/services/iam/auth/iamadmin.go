package auth

import (
	"context"
	"errors"
	"fmt"

	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/events"
	"csserver/internal/services/iam/appuser"

	"github.com/Nerzal/gocloak/v13"

	log "github.com/sirupsen/logrus"
)

type IAMAdminService struct {
	KCClient    *gocloak.GoCloak
	KCAdminUser string
	KCAdminPass string
	PubSub      events.PubSubProvider
	UserService appuser.AppuserService
}

func NewIAMAdminService(
	client *gocloak.GoCloak,
	pubsub events.PubSubProvider,
	adminUser string,
	adminPass string,
	userService appuser.AppuserService,
) IAMAdminService {
	out := IAMAdminService{}
	out.KCClient = client
	out.KCAdminUser = adminUser
	out.KCAdminPass = adminPass
	out.PubSub = pubsub
	out.UserService = userService

	return out
}

// CreateUser create a new Keycloak user
func (s *IAMAdminService) CreateUser(ctx context.Context, realm string, user *appuser.Appuser) (common.UpdateResult[appuser.Appuser], error) {
	val := user.Validate()
	if !val.Pass {
		return common.NewUpdateResult[appuser.Appuser](val, nil), fmt.Errorf("validation failed")
	}

	token, err := s.getAdminToken(ctx)
	if err != nil {
		return common.NewUpdateResult[appuser.Appuser](val, nil), err
	}

	dbu, err := s.UserService.GetAppuser(ctx, user.Email)
	if err != nil {
		log.Errorf("iamadmin getUser: %s", err)
	}
	if dbu != nil {
		user.ID = dbu.ID
	}

	updateResult, err := s.UserService.UpsertAppuser(ctx, appuser.Appuser{
		ControlFields: common.ControlFields{
			ID: user.ID,
		},
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		ProfileImage: user.ProfileImage,
	})
	if err != nil {
		return common.NewUpdateResult[appuser.Appuser](val, nil), err
	}

	dbUser := *updateResult.Object

	if dbUser == nil {
		return common.NewUpdateResult[appuser.Appuser](val, nil), err
	}

	props := make(map[string][]string)

	props["profileImage"] = []string{user.ProfileImage}
	props["dbid"] = []string{dbUser.ID}

	u := gocloak.User{
		FirstName:  &user.FirstName,
		LastName:   &user.LastName,
		Email:      &user.Email,
		Enabled:    gocloak.BoolP(true),
		Username:   &user.Email,
		Attributes: &props,
	}

	existingKCUser, _ := s.GetUser(ctx, realm, user.Email)

	var uid string

	if existingKCUser == nil {
		uid, err = s.KCClient.CreateUser(ctx, token.AccessToken, realm, u)
	} else {
		uid = existingKCUser.ID
		u.ID = &uid
		err = s.KCClient.UpdateUser(ctx, token.AccessToken, realm, u)
	}
	if err != nil {
		return common.NewUpdateResult[appuser.Appuser](val, nil), err
	}

	err = s.KCClient.SetPassword(ctx, token.AccessToken, uid, realm, user.Password, false)
	if err != nil {
		log.Errorf("iamadmin keycloak setPassword: %s", err)
		return common.NewUpdateResult[appuser.Appuser](val, nil), err
	}

	return common.NewUpdateResult[appuser.Appuser](val, nil), err
}

// UpdateUser create a new Keycloak user
func (s *IAMAdminService) UpdateUser(ctx context.Context, realm string, user appuser.Appuser) (common.UpdateResult[*common.BaseModel[appuser.Appuser]], error) {
	val := user.Validate()
	if !val.Pass {
		return common.NewUpdateResult[*common.BaseModel[appuser.Appuser]](val, nil), fmt.Errorf("validation failed")
	}

	token, err := s.getAdminToken(ctx)
	if err != nil {
		return common.NewUpdateResult[*common.BaseModel[appuser.Appuser]](val, nil), err
	}

	dbUser, err := s.UserService.GetAppuser(ctx, user.Email)
	if err != nil {
		return common.NewUpdateResult[*common.BaseModel[appuser.Appuser]](val, nil), err
	}

	dbUser.FirstName = user.FirstName
	dbUser.LastName = user.LastName
	dbUser.ProfileImage = user.ProfileImage
	dbUser.Email = user.Email

	kcUser, err := s.getKCUser(ctx, realm, user.Email)
	if err != nil {
		return common.NewUpdateResult[*common.BaseModel[appuser.Appuser]](val, nil), err
	}

	props := make(map[string][]string)

	props["profileImage"] = []string{user.ProfileImage}
	props["dbid"] = []string{dbUser.ID}

	kcUser.FirstName = &user.FirstName
	kcUser.LastName = &user.LastName
	kcUser.Email = &user.Email
	kcUser.Attributes = &props

	err = s.KCClient.UpdateUser(ctx, token.AccessToken, realm, *kcUser)
	if err != nil {
		return common.NewUpdateResult[*common.BaseModel[appuser.Appuser]](val, nil), err
	}

	return s.UserService.UpdateAppuser(ctx, *dbUser)
}

// getAdminToken return an token that can be used for
func (s *IAMAdminService) getAdminToken(ctx context.Context) (*gocloak.JWT, error) {
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		config.Config.Security.KeycloakMasterRealmName)

	return token, err
}

// GetCurrentUser gets the user based on the credentials from the context
func (s *IAMAdminService) GetCurrentUser(ctx context.Context, realm string) (*appuser.Appuser, error) {
	contextUserEmail, ok := ctx.Value(config.UserEmailKey).(string)
	if !ok {
		return nil, errors.New("context does not contain user email")
	}

	user, err := s.GetUser(ctx, realm, contextUserEmail)
	return user, err
}

// getKCUser retrieve a user record from Keycloak
func (s *IAMAdminService) getKCUser(ctx context.Context, realm string, idOrEmail string) (*gocloak.User, error) {
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		config.Config.Security.KeycloakMasterRealmName)
	if err != nil {
		return nil, err
	}

	params := gocloak.GetUsersParams{
		Email: &idOrEmail,
	}

	users, err := s.KCClient.GetUsers(ctx, token.AccessToken, realm, params)
	if err != nil || len(users) == 0 {
		return nil, err
	}

	u := users[0]

	return u, nil
}

// GetUser get a user based on the passed in email
func (s *IAMAdminService) GetUser(ctx context.Context, realm string, idOrEmail string) (*appuser.Appuser, error) {
	u, err := s.getKCUser(ctx, realm, idOrEmail)

	if u == nil {
		return nil, err
	}

	//---here in case DB values are needed
	_, err = s.UserService.GetAppuser(ctx, *u.Email)
	if err != nil {
		return nil, err
	}

	user := appuser.Appuser{
		ControlFields: common.ControlFields{
			ID: *u.ID,
		},
		FirstName:    *u.FirstName,
		LastName:     *u.LastName,
		Email:        *u.Email,
		ProfileImage: getProfileImage(u.Attributes),
		DBID:         getDBID(u.Attributes),
	}
	return &user, err
}

// GetUser get a user based on the passed in email
func (s *IAMAdminService) FindAllUsers(ctx context.Context, realm string) (common.PagedResults[appuser.Appuser], error) {
	pagingResults := common.NewPagedResultsForAllRecords[appuser.Appuser]()
	token, err := s.KCClient.LoginAdmin(ctx,
		s.KCAdminUser,
		s.KCAdminPass,
		config.Config.Security.KeycloakMasterRealmName)
	if err != nil {
		return pagingResults, err
	}

	params := gocloak.GetUsersParams{}

	users, err := s.KCClient.GetUsers(ctx, token.AccessToken, realm, params)
	if err != nil || len(users) == 0 {
		return pagingResults, err
	}

	out := []appuser.Appuser{}

	for i := range users {
		user := appuser.Appuser{
			ControlFields: common.ControlFields{
				ID: *users[i].ID,
			},
			FirstName:    *users[i].FirstName,
			LastName:     *users[i].LastName,
			Email:        *users[i].Email,
			DBID:         getDBID(users[i].Attributes),
			ProfileImage: getProfileImage(users[i].Attributes),
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
