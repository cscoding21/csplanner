// Code generated . DO NOT EDIT.
// ################################## DO NOT EDIT THIS FILE ######################################
// Common Sense Coding (https://github.com/cscoding21/csgen)

// Generate Date: 2024-07-07 14:02:37.632927303 -0700 PDT m=+0.000966882
// Implementation Name: csplanner
// Developer Note: The contents of this file will be recreated each time its generator is called

// -----------------------------------------------------------------------------------------------

package list

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
const ListIdentifier = "list"

// ListService is a service for interacting with lists.
type ListService struct {
	DBClient      surreal.DBClient
	ContextHelper interfaces.ContextHelpers
}

// NewListService creates a new List service.
func NewListService(
	db surreal.DBClient,
	ch config.ContextHelper) *ListService {

	return &ListService{
		DBClient:      db,
		ContextHelper: &ch,
	}
}

// GetListByID gets a List by its ID.
func (s *ListService) GetListByID(ctx context.Context, id string) (*List, error) {
	outData, err := s.DBClient.GetObjectById(id)
	if err != nil {
		return common.HandleReturnWithValue[List](nil, err)
	}

	output, err := marshal.SurrealUnmarshal[List](outData)

	return common.HandleReturnWithValue(
		output,
		err,
	)
}

// FindAllLists return all List in the system
func (s *ListService) FindAllLists(ctx context.Context) (common.PagedResults[List], error) {
	pagingResults := common.NewPagedResultsForAllRecords[List]()
	sql := "select * from list where deleted_at is null order by name"

	results, resultCount, err := s.DBClient.FindPagedObjects(sql, pagingResults.Pagination, pagingResults.Filters)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Pagination.TotalResults = &resultCount
	unpacked, err := marshal.SurrealSmartUnmarshal[[]List](results)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Results = common.RefToVal(unpacked)
	return pagingResults, nil
}

// CreateList creates a new List.
func (s *ListService) CreateList(ctx context.Context, input *List) (*List, error) {

	val := input.Validate()
	if !val.Pass {
		return common.HandleReturnWithValue[List](nil, val.Error("List validation failed"))
	}

	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	outData, err := s.DBClient.CreateObject(userID, ListIdentifier, input)
	if err != nil {
		return common.HandleReturnWithValue[List](nil, err)
	}

	outArray, err := marshal.SurrealUnmarshal[[]List](outData)

	outObj := utils.RefToVal(outArray)[0]
	return common.HandleReturnWithValue(&outObj, err)
}

// UpdateList update an existing List.
func (s *ListService) UpdateList(ctx context.Context, input *List) (*List, error) {

	val := input.Validate()
	if !val.Pass {
		return common.HandleReturnWithValue[List](nil, val.Error("List validation failed"))
	}

	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	outData, err := s.DBClient.UpdateObject(userID, input.ID, input)
	if err != nil {
		return common.HandleReturnWithValue[List](nil, err)
	}

	outArray, err := marshal.SurrealUnmarshal[[]List](outData)

	outObj := utils.RefToVal(outArray)[0]
	return common.HandleReturnWithValue(&outObj, err)
}

// DeleteList deletes a List.
func (s *ListService) DeleteList(ctx context.Context, id string) error {
	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	list, err := s.GetListByID(ctx, id)
	if err != nil {
		return common.HandleReturn(err)
	}

	return common.HandleReturn(
		s.DBClient.SoftDeleteObject(userID, list),
	)
}
