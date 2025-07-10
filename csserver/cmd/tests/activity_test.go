package tests

import (
	"csserver/internal/appserv/csmap"
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/common"
	"csserver/internal/services/activity"
	"fmt"
	"testing"
)

func TestFindActivity(t *testing.T) {
	ctx := getTestContext()
	service := factory.GetActivityService(ctx)

	filters := common.NewPagedResultsForAllRecords[activity.Activity]()

	comments, err := service.FindActivitys(ctx, filters.Pagination, filters.Filters)
	if err != nil {
		t.Error(err)
		return
	}

	for _, c := range comments.Results {
		base := csmap.GetDataEnvelope(&c)
		csmap.AugmentBaseModel(ctx, base)
		cut := csmap.ActivityActivityToIdl(c.Data)
		fmt.Println(cut)

		drawActivity(cut)
	}
}

func drawActivity(act idl.Activity) {
	fmt.Println("-------------------------------------------")
	fmt.Printf("ActivityID: %s\n", act.ID)
	fmt.Printf("Summary: %s\n", act.Summary)
	fmt.Printf("Context: %s\n", act.Context)
}
