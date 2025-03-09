package list

import (
	"context"
	"fmt"

	"csserver/internal/common"
	"csserver/internal/providers/postgres"
)

const (
	ListNameSkills        = "Skills"
	ListNameFundingSource = "Funding Source"
	ListNameValueCategory = "Value Catetgory"
)

// GetList get a list based on the passed in email
func (s *ListService) GetList(ctx context.Context, idOrName string) (*common.BaseModel[List], error) {
	list, err := postgres.GetObject[List](ctx, s.db, ListIdentifier,
		fmt.Sprintf("SELECT * FROM %s WHERE data->>'name' = $1 OR id = $1", ListIdentifier),
		idOrName,
	)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// SaveList higher order function for saving an existing list
func (s *ListService) SaveList(ctx context.Context, list List) (common.UpdateResult[*common.BaseModel[List]], error) {
	id := list.ID

	existingList, err := s.GetList(ctx, id)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[List]](nil, err)
	}

	existingList.Data.Values = list.Values

	return s.UpdateList(ctx, existingList.Data)
}
