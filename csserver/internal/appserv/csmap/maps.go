package csmap

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/common"
)

// GetIDLPageAndFilters return pagination and filters objects converted to IDL
func GetIDLPageAndFilters(paging common.Pagination, filters common.Filters) (idl.Pagination, idl.Filters) {
	p := PaginationCommonToIdl(paging)
	f := FiltersCommonToIdl(filters)

	return p, f
}
