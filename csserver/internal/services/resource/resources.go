package resource

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/marshal"

	log "github.com/sirupsen/logrus"
)

func (s *ResourceService) FindResources(ctx context.Context, paging common.Pagination, filters common.Filters) (common.PagedResults[Resource], error) {
	out := common.NewPagedResults[Resource](paging, filters)

	rawResults, count, err := s.DBClient.FindPagedObjects("", paging, filters)
	if err != nil {
		log.Error()
		return out, err
	}

	out.Pagination.TotalResults = &count
	_, err = marshal.SurrealUnmarshalRaw[[]Resource](rawResults, out.Results)
	if err != nil {
		log.Error()
		return out, err
	}

	return out, nil
}
