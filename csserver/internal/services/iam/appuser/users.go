package appuser

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/providers/postgres"
	"fmt"
)

// GetAppuser return a user based on their ID or email address
func (s *AppuserService) GetAppuser(ctx context.Context, idOrEmail string) (*Appuser, error) {
	user, err := postgres.GetObject[Appuser](ctx, s.db, AppuserIdentifier,
		fmt.Sprintf("SELECT * FROM %s WHERE data->>'email' = $1 OR id = $1", AppuserIdentifier),
		idOrEmail,
	)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("user %s not found in database", idOrEmail)
	}

	return &user.Data, err
}

// AppupsertUserByEmail create or update a user using the email as a key
func (s *AppuserService) UpsertUserByEmail(ctx context.Context, usr Appuser) (common.UpdateResult[*common.BaseModel[Appuser]], error) {
	existingUser, _ := s.GetAppuser(ctx, usr.Email)

	if existingUser == nil {
		resp, err := s.CreateAppuser(ctx, usr)
		return resp, err
	}

	usr.ID = existingUser.ID
	resp, err := s.UpdateAppuser(ctx, usr)
	return resp, err
}
