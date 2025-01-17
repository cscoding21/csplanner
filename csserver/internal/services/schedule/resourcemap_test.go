package schedule_test

import (
	"csserver/internal/services/schedule"
	"csserver/internal/testobjects"
	"fmt"
	"testing"
)

func TestFlattenPortfolio(t *testing.T) {
	rm := testobjects.GetResourceMap()
	port := testobjects.GetTestPortfolio()

	ram, err := schedule.FlattenPortfolio(port.Schedule, rm)
	if err != nil {
		t.Error(err)
	}

	for _, m := range ram.WeekActivities {
		fmt.Println(m)
	}
}
