package list

import (
	"context"
	"errors"
	"fmt"

	"csserver/internal/common"
	"csserver/internal/marshal"

	"github.com/cscoding21/csmap/utils"
)

const (
	ListNameSkills        = "Skills"
	ListNameFundingSource = "Funding Source"
	ListNameValueCategory = "Value Catetgory"
)

// GetList get a list based on the passed in email
func (s *ListService) GetList(ctx context.Context, idOrName string) (*List, error) {
	listRaw, err := s.DBClient.GetObject(
		fmt.Sprintf("SELECT * FROM %s WHERE name = $idOrName OR id = $idOrName", ListIdentifier),
		map[string]interface{}{"idOrName": idOrName})

	if err != nil {
		return common.HandleReturnWithValue[List](nil, err)
	}

	listArray, err := marshal.SurrealSmartUnmarshal[[]List](listRaw)

	if len(*listArray) == 0 {
		return common.HandleReturnWithValue[List](nil, errors.New("list not found"))
	}

	list := utils.RefToVal(listArray)[0]
	return common.HandleReturnWithValue[List](&list, err)
}

// SaveList higher order function for saving an existing list
func (s *ListService) SaveList(ctx context.Context, list List) (common.UpdateResult[List], error) {
	id := list.ID

	existingList, err := s.GetList(ctx, id)
	if err != nil {
		return common.UpdateResult[List]{
			Object: existingList,
		}, err
	}

	existingList.Values = list.Values

	return s.UpdateList(ctx, existingList)
}
