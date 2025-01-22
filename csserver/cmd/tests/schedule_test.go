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

	put, err := ps.GetProjectByID(ctx, "project:4k1b1vu9f1t64it9sk46")
	if err != nil {
		t.Error(err)
	}

	rm, err := ps.GetResourceMap(true)
	if err != nil {
		t.Error(err)
	}

	ram, err := ss.GetInitialResourceAllocationMap(rm)
	if err != nil {
		t.Error(err)
	}

	sch, err := ss.CalculateProjectSchedule(ctx, put, startTime, ram)
	if err != nil {
		t.Error(err)
	}

	drawScheduleResult(sch)
}
