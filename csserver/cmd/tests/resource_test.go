package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/resource"
	"fmt"
	"testing"
)

func TestFindResources(t *testing.T) {
	ctx := getTestContext()

	service := factory.GetResourceService(ctx)

	filters := common.NewPagedResultsForAllRecords[resource.Resource]()

	filters.Filters.AddFilter(common.Filter{
		Key:       "data->>'name'",
		Value:     "Biggs",
		Operation: common.FilterOperationNotEqual,
	})

	results, err := service.FindResources(ctx, filters.Pagination, filters.Filters)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Total results returned: %v\n", results.Pagination.TotalResults)
	fmt.Printf("Page %v of %v RPP\n", results.Pagination.PageNumber, results.Pagination.ResultsPerPage)

	records := common.ExtractDataFromBase(results.Results)
	for _, r := range records {
		fmt.Printf(" -- %s\n", r.Name)
	}
}
