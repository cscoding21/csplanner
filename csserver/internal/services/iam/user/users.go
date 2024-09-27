package user

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/marshal"
	"csserver/internal/utils"
	"fmt"
)

// GetUser return a user based on their ID or email address
func (s *UserService) GetUser(idOrEmail string) (*User, error) {
	userRaw, err := s.DBClient.GetObject(
		fmt.Sprintf("SELECT * FROM %s WHERE email = $idOrEmail OR id = $idOrEmail", UserIdentifier),
		map[string]interface{}{"idOrEmail": idOrEmail})

	if err != nil {
		return common.HandleReturnWithValue[User](nil, err)
	}

	userArray, err := marshal.SurrealSmartUnmarshal[[]User](userRaw)
	if len(*userArray) == 0 || err != nil {
		return common.HandleReturnWithValue[User](nil, fmt.Errorf("user not found: %s", err))
	}

	u := utils.RefToVal(userArray)[0]
	return common.HandleReturnWithValue(&u, err)
}

// UpsertUserByEmail create or update a user using the email as a key
func (s *UserService) UpsertUserByEmail(ctx context.Context, usr User) (*common.UpdateResult[User], error) {
	existingUser, _ := s.GetUser(usr.Email)

	if existingUser == nil {
		resp, err := s.CreateUser(ctx, &usr)
		return &resp, err
	}

	usr.ID = existingUser.ID
	resp, err := s.UpdateUser(ctx, &usr)
	return &resp, err
}
