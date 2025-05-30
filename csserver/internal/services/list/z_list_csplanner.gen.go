// Code generated . DO NOT EDIT.
// ################################## DO NOT EDIT THIS FILE ######################################
// Common Sense Coding (https://github.com/cscoding21/csgen)

// Generate Date: 2025-05-22 08:42:38.807676173 -0700 PDT m=+0.382534896
// Implementation Name: csplanner
// Developer Note: The contents of this file will be recreated each time its generator is called

// -----------------------------------------------------------------------------------------------

package list

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/events"
	"csserver/internal/providers/postgres"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ---This is the name of the object in the database
const ListIdentifier = postgres.TableName("list")

// ListService is a service for interacting with lists.
type ListService struct {
	db     *pgxpool.Pool
	pubsub events.PubSubProvider
}

// NewListService creates a new List service.
func NewListService(
	db *pgxpool.Pool,
	ps events.PubSubProvider) *ListService {

	return &ListService{
		db:     db,
		pubsub: ps,
	}
}

// GetListByID gets a List by its ID.
func (s *ListService) GetListByID(ctx context.Context, id string) (*common.BaseModel[List], error) {
	return postgres.GetObjectByID[List](ctx, s.db, ListIdentifier, id)
}

// FindAllLists return all List in the system
func (s *ListService) FindAllLists(ctx context.Context) (common.PagedResults[common.BaseModel[List]], error) {
	return postgres.FindAllObjects[List](ctx, s.db, ListIdentifier)
}

// FindList return a paged list of List based on filter criteria
func (s *ListService) FindLists(ctx context.Context, paging common.Pagination, filters common.Filters) (common.PagedResults[common.BaseModel[List]], error) {
	out := common.NewPagedResults[List](paging, filters)

	whereSql, params := postgres.BuildWhereClauseFromFilters(&filters)

	sql := fmt.Sprintf("SELECT * FROM %s WHERE true AND deleted_at is null %s ORDER BY created_at DESC", ListIdentifier, whereSql)

	return postgres.FindPagedObjects[List](ctx, s.db, sql, out.Pagination, out.Filters, params)
}

// CreateList creates a new List.
func (s *ListService) CreateList(ctx context.Context, input List) (common.UpdateResult[*common.BaseModel[List]], error) {

	val := input.Validate()
	if !val.Pass {
		return common.NewUpdateResult[*common.BaseModel[List]](val, nil), fmt.Errorf("validation failed")
	}

	obj, err := postgres.UpdateObject(ctx, s.db, input, ListIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

// UpdateList update an existing List.
func (s *ListService) UpdateList(ctx context.Context, input List) (common.UpdateResult[*common.BaseModel[List]], error) {

	val := input.Validate()
	if !val.Pass {
		return common.NewUpdateResult[*common.BaseModel[List]](val, nil), fmt.Errorf("validation failed")
	}

	obj, err := postgres.UpdateObject(ctx, s.db, input, ListIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

// UpsertList create or update a List
func (s *ListService) UpsertList(ctx context.Context, input List) (common.UpdateResult[*common.BaseModel[List]], error) {

	val := input.Validate()
	if !val.Pass {
		return common.NewUpdateResult[*common.BaseModel[List]](val, nil), fmt.Errorf("validation failed")
	}

	obj, err := postgres.UpdateObject(ctx, s.db, input, ListIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

// DeleteList deletes a List.
func (s *ListService) DeleteList(ctx context.Context, id string) error {
	return postgres.SoftDelete(ctx, s.db, ListIdentifier, id)
}
