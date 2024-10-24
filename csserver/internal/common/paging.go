package common

import (
	"math"
	"strings"
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
	FilterFuzzyLike               FilterOperation = "fl"
	FilterCustom                  FilterOperation = "custom"
)

// NewPagedResults helper method for generating a paged results object
func NewPagedResults[T any](paging Pagination, filters Filters) PagedResults[T] {
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

	filters := Filters{Filters: []Filter{}}

	paging := Pagination{
		ResultsPerPage: &maxInt,
		PageNumber:     &one,
	}

	return NewPagedResults[T](paging, filters)
}

type PagedResults[T any] struct {
	Pagination Pagination
	Filters    Filters
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

type Filters struct {
	Filters []Filter
}

type Filter struct {
	Key        string
	Value      interface{}
	Operation  FilterOperation
	CustomName *string
}

func (filters *Filters) GetFilter(key string) *Filter {
	for _, f := range filters.Filters {
		if f.Key == key {
			return &f
		}
	}

	return nil
}

func (filters *Filters) GetFilterValue(key string) interface{} {
	f := filters.GetFilter(key)
	if f == nil {
		return ""
	}

	return f.Value
}

func (filters *Filters) GetFiltersAsMap() map[string]interface{} {
	out := make(map[string]interface{})

	for _, f := range filters.Filters {
		ky := strings.Replace(f.Key, ".", "_", 1)
		out[ky] = f.Value
	}

	return out
}

func (filters *Filters) AddFilter(filterToAdd Filter) {
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

func (filters *Filters) HasFilters() bool {
	return len(filters.Filters) > 0
}
