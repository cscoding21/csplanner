package tests

import (
	"csserver/internal/appserv/factory"
	"testing"
	"time"
)

func TestCalculateProjectSchedule(t *testing.T) {
	ctx := getTestContext()

	ss := factory.GetScheduleService(ctx)
	ps := factory.GetProjectService(ctx)

	startTime := time.Now()

	rs := factory.GetResourceService(ctx)

	resourceMap, err := rs.GetResourceMap(ctx, false)
	if err != nil {
		t.Error(err)
	}

	put, err := ps.GetProjectByID(ctx, "6BkusOtrRZ6MGnk9DmwNSA")
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
