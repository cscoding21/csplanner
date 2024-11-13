package portfolio

import (
	"fmt"
	"testing"
	"time"
)

func TestWeeksBetween(t *testing.T) {
	start := time.Now()
	end := start.Add(24 * time.Hour * 7 * 5)
	expected := 6
	weeks := GetWeeks(start, end)
	for _, week := range weeks {
		fmt.Println(week.Year, week.WeekNumber, week.Begin.Format("2006-01-02T15:04:05Z07:00"), week.End.Format("2006-01-02T15:04:05Z07:00")) // adjust format as needed
	}

	if len(weeks) != expected {
		t.Errorf("unexpected number of weeks.  wanted %v - got %v\n", expected, len(weeks))
	}
}
