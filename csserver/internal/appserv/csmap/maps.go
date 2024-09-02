package csmap

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/common"

	"github.com/cscoding21/csval/validate"
)

// GetPageAndFilterIDL return pagination and filters objects converted to IDL
func GetPageAndFilterIDL(paging common.Pagination, filters common.Filters) (idl.Pagination, idl.Filters) {
	p := PaginationCommonToIdl(paging)
	f := FiltersCommonToIdl(filters)

	return p, f
}

// GetPageAndFilterModel return pagination and filters objects converted to IDL
func GetPageAndFilterModel(paging idl.Pagination, filters idl.Filters) (common.Pagination, common.Filters) {
	//---TODO: insert the proper mapping code
	p := common.Pagination{}
	f := common.Filters{Filters: []common.Filter{}}

	return p, f
}

// GetValidationResultIDL create a validation result IDL from the caval package objects
func GetValidationResultIDL(result validate.ValidationResult) idl.ValidationResult {
	out := idl.ValidationResult{
		Pass:     result.Pass,
		Messages: []*idl.ValidationMessage{},
	}

	if len(result.Messages) > 0 {
		for _, m := range result.Messages {
			msg := idl.ValidationMessage{Field: m.Field, Message: m.Message}
			out.Messages = append(out.Messages, &msg)
		}
	}

	return out
}
