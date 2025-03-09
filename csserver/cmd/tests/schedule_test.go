package tests

import (
	"csserver/internal/appserv/factory"
	"testing"
	"time"
)

func TestCalculateProjectSchedule(t *testing.T) {
	ss := factory.GetScheduleService()
	ps := factory.GetProjectService()
	ctx := getTestContext()

	startTime := time.Now()

	rs := factory.GetResourceService()

	resourceMap, err := rs.GetResourceMap(ctx, false)
	if err != nil {
		t.Error(err)
	}

	put, err := ps.GetProjectByID(ctx, "project:1")
	if err != nil {
		t.Error(err)
	}

	ram, err := ss.GetInitialResourceAllocationMap(resourceMap)
	if err != nil {
		t.Error(err)
	}

	sch, err := ss.CalculateProjectSchedule(ctx, &put.Data, startTime, ram)
	if err != nil {
		t.Error(err)
	}

	drawScheduleResult(sch)
}
