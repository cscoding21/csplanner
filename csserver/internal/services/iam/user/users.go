package user

import (
	"context"
	"errors"
	"fmt"

	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/marshal"

	"github.com/cscoding21/csmap/utils"
)

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
	userRaw, err := s.DBClient.GetObject(
		fmt.Sprintf("SELECT * OMIT password FROM %s WHERE email = $idOrEmail OR id = $idOrEmail", UserIdentifier),
		map[string]interface{}{"idOrEmail": idOrEmail})

	if err != nil {
		return common.HandleReturnWithValue[User](nil, err)
	}

	userArray, err := marshal.SurrealSmartUnmarshal[[]User](userRaw)

	if len(*userArray) == 0 {
		return common.HandleReturnWithValue[User](nil, errors.New("user not found"))
	}

	user := utils.RefToVal(userArray)[0]
	return common.HandleReturnWithValue[User](&user, err)
}
