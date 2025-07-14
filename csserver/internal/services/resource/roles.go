package resource

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/providers/postgres"
	"fmt"

	"github.com/cscoding21/csval/validate"

	log "github.com/sirupsen/logrus"
)

// ---This is the name of the object in the database
const RoleIdentifier = postgres.TableName("role")

func (s *ResourceService) UpdateAllRoles(ctx context.Context, roles []*Role) (common.UpdateResult[*common.BaseModel[Role]], error) {
	results := []common.UpdateResult[*common.BaseModel[Role]]{}
	toDeleteMap := make(map[string]bool)

	allRoles, err := s.FindAllRoles(ctx)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Role]](nil, err)
	}

	for _, ar := range allRoles.Results {
		toDeleteMap[ar.Data.ID] = true
	}

	for _, r := range roles {
		toDeleteMap[r.ID] = false
		result, err := s.UpsertRole(ctx, *r)
		if err != nil {
			return common.NewFailingUpdateResult[*common.BaseModel[Role]](nil, err)
		}

		results = append(results, result)
	}

	for k, v := range toDeleteMap {
		if v {
			err = s.DeleteRole(ctx, k)
			if err != nil {
				return common.NewFailingUpdateResult[*common.BaseModel[Role]](nil, err)
			}
		} else {
			err = postgres.UndoSoftDelete(ctx, s.db, RoleIdentifier, k)
			if err != nil {
				return common.NewFailingUpdateResult[*common.BaseModel[Role]](nil, err)
			}
		}
	}

	out, err := common.MergeUpdateResults(results)

	return out, err
}

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

// UpsertRoles create or update a Roles
func (s *ResourceService) UpsertRole(ctx context.Context, input Role) (common.UpdateResult[*common.BaseModel[Role]], error) {

	val := validate.NewSuccessValidationResult()

	obj, err := postgres.UpdateObject(ctx, s.db, input, RoleIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	eventType := "updated"

	log.Warnf("%v - %v", obj.CreatedAt, obj.UpdatedAt)
	if obj.CreatedAt.Equal(obj.UpdatedAt) {
		eventType = "created"
	}

	newRole := *obj
	s.pubsub.StreamPublish(ctx,
		string(ResourceIdentifier),
		"role",
		eventType,
		map[string]any{
			"id":   newRole.ID,
			"name": newRole.Data.Name,
		})

	return common.NewUpdateResult(val, &obj), nil
}

// DeleteRoles deletes a Roles.
func (s *ResourceService) DeleteRole(ctx context.Context, id string) error {
	return postgres.SoftDelete(ctx, s.db, RoleIdentifier, id)
}
