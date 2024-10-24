package csmap

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/common"

	"github.com/cscoding21/csval/validate"
)

// GetPageAndFilterIDL return pagination and filters objects converted to IDL
func GetPageAndFilterIdl(paging common.Pagination, filters common.Filters) (idl.Pagination, idl.Filters) {
	p := PaginationCommonToIdl(paging)
	f := FiltersCommonToIdl(filters)

	return p, f
}

// GetPageAndFilterModel return pagination and filters objects converted to IDL
func GetPageAndFilterModel(paging idl.InputPagination, filters *idl.InputFilters) (common.Pagination, common.Filters) {
	//---TODO: insert the proper mapping code
	p := common.Pagination{
		PageNumber:     paging.PageNumber,
		ResultsPerPage: paging.ResultsPerPage,
	}
	f := common.Filters{Filters: []common.Filter{}}

	if filters != nil {
		for _, fil := range filters.Filters {
			thisFilter := common.Filter{
				Key:        fil.Key,
				Value:      fil.Value,
				Operation:  common.FilterOperation(fil.Operation),
				CustomName: fil.CustomName,
			}

			f.Filters = append(f.Filters, thisFilter)
		}
	}

	return p, f
}

// GetValidationResultIdl create a validation result IDL from the caval package objects
func GetValidationResultIdl(result validate.ValidationResult) idl.ValidationResult {
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

// GetStatusFromError returns a failing status when an error occurs
func GetStatusFromError(err error) (*idl.Status, error) {
	status := idl.Status{}

	if err != nil {
		status.Success = false
		status.Message = []string{err.Error()}

		return &status, err
	}

	status.Success = true
	return &status, nil
}

// GetStatusFromUpdateResult return an IDL status from an update result
func GetStatusFromUpdateResult[T any](result common.UpdateResult[T]) (*idl.Status, error) {
	status := idl.Status{}

	if result.ValidationResult.Pass {
		status.Success = true

		return &status, nil
	}

	status.Success = false
	if result.ValidationResult != nil {
		status.ValidationResult = common.ValToRef(GetValidationResultIdl(*result.ValidationResult))
	}

	return &status, nil
}
