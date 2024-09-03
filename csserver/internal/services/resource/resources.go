package resource

// FindResources return a paged list of resources based on filter criteria
/*
func (s *ResourceService) FindResources(ctx context.Context, paging common.Pagination, filters common.Filters) (common.PagedResults[Resource], error) {
	out := common.NewPagedResults[Resource](paging, filters)

	whereSql, _ := s.DBClient.BuildWhereClauseFromFilters(&filters)

	sql := fmt.Sprintf(`SELECT *
			FROM resource
			WHERE true
				AND deleted_at is null
				%s
			ORDER BY name
			`, whereSql)

	rawResults, count, err := s.DBClient.FindPagedObjects(sql, paging, filters)
	if err != nil {
		log.Error()
		return out, err
	}

	out.Pagination.TotalResults = &count
	unpacked, err := marshal.SurrealSmartUnmarshal[[]Resource](rawResults)
	if err != nil {
		return out, err
	}

	out.Results = common.RefToVal(unpacked)
	return out, nil
}
*/
