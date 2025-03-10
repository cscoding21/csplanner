package resource

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/providers/postgres"
	"fmt"

	"github.com/cscoding21/csval/validate"
)

// ---This is the name of the object in the database
const RoleIdentifier = postgres.TableName("role")

// GetRoleByID gets a Role by its ID.
func (s *ResourceService) GetRoleByID(ctx context.Context, id string) (*common.BaseModel[Role], error) {
	return postgres.GetObjectByID[Role](ctx, s.db, RoleIdentifier, id)
}

// FindAllRoles return all Roles in the system
func (s *ResourceService) FindAllRoles(ctx context.Context) (common.PagedResults[common.BaseModel[Role]], error) {
	return postgres.FindAllObjects[Role](ctx, s.db, RoleIdentifier)
}

// FindRoles return a paged list of Roles based on filter criteria
func (s *ResourceService) FindRoles(ctx context.Context, paging common.Pagination, filters common.Filters) (common.PagedResults[common.BaseModel[Role]], error) {
	out := common.NewPagedResults[Role](paging, filters)

	whereSql, params := postgres.BuildWhereClauseFromFilters(&filters)

	sql := fmt.Sprintf("SELECT * FROM %s WHERE true AND deleted_at is null %s ORDER BY name", RoleIdentifier, whereSql)

	return postgres.FindPagedObjects[Role](ctx, s.db, sql, out.Pagination, out.Filters, params)
}

// CreateRoles creates a new Roles.
func (s *ResourceService) CreateRole(ctx context.Context, input Role) (common.UpdateResult[*common.BaseModel[Role]], error) {

	val := validate.NewSuccessValidationResult()

	obj, err := postgres.UpdateObject(ctx, s.db, input, RoleIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

// UpdateRoles update an existing Roles.
func (s *ResourceService) UpdateRole(ctx context.Context, input Role) (common.UpdateResult[*common.BaseModel[Role]], error) {

	val := validate.NewSuccessValidationResult()

	obj, err := postgres.UpdateObject(ctx, s.db, input, RoleIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

// UpsertRoles create or update a Roles
func (s *ResourceService) UpsertRole(ctx context.Context, input Role) (common.UpdateResult[*common.BaseModel[Role]], error) {

	val := validate.NewSuccessValidationResult()

	obj, err := postgres.UpdateObject(ctx, s.db, input, RoleIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

// DeleteRoles deletes a Roles.
func (s *ResourceService) DeleteRole(ctx context.Context, id string) error {
	return postgres.SoftDelete(ctx, s.db, RoleIdentifier, id)
}
