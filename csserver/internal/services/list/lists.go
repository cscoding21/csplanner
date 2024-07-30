package list

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/marshal"
	"errors"
	"fmt"

	"github.com/cscoding21/csmap/utils"
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
