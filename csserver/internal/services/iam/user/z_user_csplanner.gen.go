// Code generated . DO NOT EDIT.
// ################################## DO NOT EDIT THIS FILE ######################################
// Common Sense Coding (https://github.com/cscoding21/csgen)

// Generate Date: 2024-07-07 14:02:37.748589763 -0700 PDT m=+0.116629343
// Implementation Name: csplanner
// Developer Note: The contents of this file will be recreated each time its generator is called

// -----------------------------------------------------------------------------------------------

package user

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/interfaces"
	"csserver/internal/marshal"
	"csserver/internal/providers/surreal"

	"github.com/cscoding21/csmap/utils"
)

// ---This is the name of the object in the database
const UserIdentifier = "user"

// UserService is a service for interacting with lists.
type UserService struct {
	DBClient      surreal.DBClient
	ContextHelper interfaces.ContextHelpers
}

// NewUserService creates a new User service.
func NewUserService(
	db surreal.DBClient,
	ch config.ContextHelper) *UserService {

	return &UserService{
		DBClient:      db,
		ContextHelper: &ch,
	}
}

// GetUserByID gets a User by its ID.
func (s *UserService) GetUserByID(ctx context.Context, id string) (*User, error) {
	outData, err := s.DBClient.GetObjectById(id)
	if err != nil {
		return common.HandleReturnWithValue[User](nil, err)
	}

	output, err := marshal.SurrealUnmarshal[User](outData)

	return common.HandleReturnWithValue(
		output,
		err,
	)
}

// FindAllUsers return all User in the system
func (s *UserService) FindAllUsers(ctx context.Context) (common.PagedResults[User], error) {
	pagingResults := common.NewPagedResultsForAllRecords[User]()
	sql := "select * from user where deleted_at is null order by name"

	results, resultCount, err := s.DBClient.FindPagedObjects(sql, pagingResults.Pagination, pagingResults.Filters)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Pagination.TotalResults = &resultCount
	unpacked, err := marshal.SurrealSmartUnmarshal[[]User](results)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Results = common.RefToVal(unpacked)
	return pagingResults, nil
}

// CreateUser creates a new User.
func (s *UserService) CreateUser(ctx context.Context, input *User) (*User, error) {

	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	outData, err := s.DBClient.CreateObject(userID, UserIdentifier, input)
	if err != nil {
		return common.HandleReturnWithValue[User](nil, err)
	}

	outArray, err := marshal.SurrealUnmarshal[[]User](outData)

	outObj := utils.RefToVal(outArray)[0]
	return common.HandleReturnWithValue(&outObj, err)
}

// UpdateUser update an existing User.
func (s *UserService) UpdateUser(ctx context.Context, input *User) (*User, error) {

	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	outData, err := s.DBClient.UpdateObject(userID, input.ID, input)
	if err != nil {
		return common.HandleReturnWithValue[User](nil, err)
	}

	outArray, err := marshal.SurrealUnmarshal[[]User](outData)

	outObj := utils.RefToVal(outArray)[0]
	return common.HandleReturnWithValue(&outObj, err)
}

// DeleteUser deletes a User.
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	list, err := s.GetUserByID(ctx, id)
	if err != nil {
		return common.HandleReturn(err)
	}

	return common.HandleReturn(
		s.DBClient.SoftDeleteObject(userID, list),
	)
}