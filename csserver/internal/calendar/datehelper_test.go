package calendar

import (
	"fmt"
	"strings"
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

func TestGetNextWeek(t *testing.T) {
	start := time.Date(2024, time.November, 21, 0, 0, 0, 0, time.UTC)
	want := "2024-11-25"

	week := GetWeek(start)
	w := GetNextWeek(week)

	if !strings.EqualFold(want, w.Begin.Format("2006-01-02")) {
		t.Errorf("week is wrong.  wanted 2024-11-25, got %v", w.Begin.Format("2006-01-02"))
	}
}
