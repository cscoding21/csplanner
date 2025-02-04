package resource

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/marshal"
	"csserver/internal/utils"
	"fmt"

	"github.com/cscoding21/csval/validate"
)

// ---This is the name of the object in the database
const RoleIdentifier = "role"

// GetRoleByID gets a Role by its ID.
func (s *ResourceService) GetRoleByID(ctx context.Context, id string) (*Role, error) {
	outData, err := s.DBClient.GetObjectById(id)
	if err != nil {
		return common.HandleReturnWithValue[Role](nil, err)
	}

	output, err := marshal.SurrealUnmarshal[Role](outData)

	return common.HandleReturnWithValue(
		output,
		err,
	)
}

// FindAllRoles return all Role in the system
func (s *ResourceService) FindAllRoles(ctx context.Context) (common.PagedResults[Role], error) {
	pagingResults := common.NewPagedResultsForAllRecords[Role]()
	sql := fmt.Sprintf("select * from %s where deleted_at is null order by name", RoleIdentifier)

	results, resultCount, err := s.DBClient.FindPagedObjects(sql, pagingResults.Pagination, pagingResults.Filters)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Pagination.TotalResults = &resultCount
	unpacked, err := marshal.SurrealSmartUnmarshal[[]Role](results)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Results = common.RefToVal(unpacked)
	return pagingResults, nil
}

// FindRoles return a paged list of Role based on filter criteria
func (s *ResourceService) FindRoles(ctx context.Context, paging common.Pagination, filters common.Filters) (common.PagedResults[Role], error) {
	out := common.NewPagedResults[Role](paging, filters)

	whereSql, _ := s.DBClient.BuildWhereClauseFromFilters(&filters)

	sql := fmt.Sprintf("SELECT * FROM %s WHERE true AND deleted_at is null %s ORDER BY name", RoleIdentifier, whereSql)

	rawResults, count, err := s.DBClient.FindPagedObjects(sql, paging, filters)
	if err != nil {
		return out, err
	}

	out.Pagination.TotalResults = &count
	unpacked, err := marshal.SurrealSmartUnmarshal[[]Role](rawResults)
	if err != nil {
		return out, err
	}

	out.Results = common.RefToVal(unpacked)
	return out, nil
}

// CreateRole creates a new Role.
func (s *ResourceService) CreateRole(ctx context.Context, input *Role) (common.UpdateResult[Role], error) {

	val := validate.NewSuccessValidationResult()

	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	outData, err := s.DBClient.CreateObject(userEmail, RoleIdentifier, input)
	if err != nil {
		return common.NewUpdateResult[Role](&val, input), err
	}

	outArray, err := marshal.SurrealUnmarshal[[]Role](outData)
	if err != nil {
		return common.NewUpdateResult[Role](&val, input), err
	}

	outObj := utils.RefToVal(outArray)[0]
	return common.NewUpdateResult[Role](&val, &outObj), nil
}

// UpdateRole update an existing Role.
func (s *ResourceService) UpdateRole(ctx context.Context, input *Role) (common.UpdateResult[Role], error) {

	val := validate.NewSuccessValidationResult()

	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	lastObj, err := s.GetRoleByID(ctx, input.ID)
	if err != nil {
		return common.NewUpdateResult[Role](&val, input), err
	}

	//---ensure the integrity of the creation data
	input.CreatedAt = lastObj.CreatedAt
	input.CreatedBy = lastObj.CreatedBy

	outData, err := s.DBClient.UpdateObject(userEmail, input.ID, input)
	if err != nil {
		return common.NewUpdateResult[Role](&val, input), err
	}

	outObj, err := marshal.SurrealUnmarshal[Role](outData)
	if err != nil {
		return common.NewUpdateResult[Role](&val, input), err
	}

	return common.NewUpdateResult[Role](&val, outObj), nil
}

// UpsertRole create or update a Role
func (s *ResourceService) UpsertRole(ctx context.Context, obj Role) (*common.UpdateResult[Role], error) {
	existingObj, _ := s.GetRoleByID(ctx, obj.ID)

	if existingObj == nil {
		resp, err := s.CreateRole(ctx, &obj)
		return &resp, err
	}

	obj.ID = existingObj.ID
	resp, err := s.UpdateRole(ctx, &obj)
	return &resp, err
}

// DeleteRole deletes a Role.
func (s *ResourceService) DeleteRole(ctx context.Context, id string) error {
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	role, err := s.GetRoleByID(ctx, id)
	if err != nil {
		return common.HandleReturn(err)
	}

	return common.HandleReturn(
		s.DBClient.SoftDeleteObject(userEmail, role),
	)
}
