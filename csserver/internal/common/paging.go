package common

import (
	"math"
)

type FilterOperation string

const (
	FilterOperationEqual          FilterOperation = "eq"
	FilterOperationNotEqual       FilterOperation = "ne"
	FilterOperationGreater        FilterOperation = "gt"
	FilterOperationLess           FilterOperation = "lt"
	FilterOperationGreaterOrEqual FilterOperation = "ge"
	FilterOperationLessOrEqual    FilterOperation = "le"
	FilterOperationIn             FilterOperation = "in"
	FilterOperationNotIn          FilterOperation = "nin"
)

// NewPagedResults helper method for generating a paged results object
func NewPagedResults[T any](paging Pagination, filters QueryFilters) PagedResults[T] {
	out := PagedResults[T]{
		Pagination: paging,
		Filters:    filters,
		Results:    []T{},
	}

	return out
}

// NewPagedResultsForAllRecords setup a paged results to return all records of a type
func NewPagedResultsForAllRecords[T any]() PagedResults[T] {
	maxInt := math.MaxInt32
	one := 1

	filters := QueryFilters{Filters: []QueryFilter{}}

	paging := Pagination{
		ResultsPerPage: &maxInt,
		PageNumber:     &one,
	}

	return NewPagedResults[T](paging, filters)
}

type PagedResults[T any] struct {
	Pagination Pagination
	Filters    QueryFilters
	Results    []T
}

type Pagination struct {
	ResultsPerPage *int
	After          *string
	PageNumber     *int
	TotalResults   *int
}

// GetOffset calculate and return the page offset based on page number and rpp
func (p *Pagination) GetOffset() int {
	if p.PageNumber == nil || *p.PageNumber < 1 {
		return 1
	}

	pn := *p.PageNumber
	rpp := *p.ResultsPerPage

	return (pn - 1) * rpp
}

type QueryFilters struct {
	Filters []QueryFilter
}

type QueryFilter struct {
	Key       string
	Value     interface{}
	Operation FilterOperation
}

func (filters *QueryFilters) GetFilter(key string) *QueryFilter {
	for _, f := range filters.Filters {
		if f.Key == key {
			return &f
		}
	}

	return nil
}

func (filters *QueryFilters) GetFilterValue(key string) interface{} {
	f := filters.GetFilter(key)
	if f == nil {
		return ""
	}

	return f.Value
}

func (filters *QueryFilters) GetFiltersAsMap() map[string]interface{} {
	out := make(map[string]interface{})

	for _, f := range filters.Filters {
		out[f.Key] = f.Value
	}

	return out
}

func (filters *QueryFilters) AddFilter(filterToAdd QueryFilter) {
	existingFilter := filters.GetFilter(filterToAdd.Key)

	if existingFilter == nil {
		filters.Filters = append(filters.Filters, filterToAdd)
		return
	}

	for i, f := range filters.Filters {
		if f.Key == filterToAdd.Key {
			filters.Filters[i] = filterToAdd
			return
		}
	}
}

func (filters *QueryFilters) HasFilters() bool {
	return len(filters.Filters) > 0
}
