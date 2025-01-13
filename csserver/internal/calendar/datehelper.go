package calendar

import "time"

type CSWeek struct {
	Begin      time.Time
	End        time.Time
	Year       int
	WeekNumber int
}

func (c *CSWeek) Before(week CSWeek) bool {
	return week.Begin.Before(c.Begin)
}

func (c *CSWeek) After(week CSWeek) bool {
	return week.Begin.After(c.Begin)
}

func (c *CSWeek) Equal(week CSWeek) bool {
	return week.Begin.Equal(c.Begin)
}

// GetWeeks return a slice of weeks based on the passed in date range
func GetWeeks(start, end time.Time) []CSWeek {
	var weeks []CSWeek
	sow := getStartOfWeek(start)

	for t := sow; t.Before(end); t = t.AddDate(0, 0, 7) {
		w := GetWeek(t)

		weeks = append(weeks, w)
	}
	return weeks
}

// GetWeek get a quantized CSWeek based on the passed in date
func GetWeek(date time.Time) CSWeek {
	sow := getStartOfWeek(date)
	eow := sow.Add(1 * time.Hour * 24 * 7).Add(-1 * time.Second)
	year, week := sow.ISOWeek()

	w := CSWeek{
		Begin:      sow,
		End:        eow,
		Year:       year,
		WeekNumber: week,
	}

	return w
}

// GetWeekFromDateString parse a string into a date and return the result of GetWeek
func GetWeekFromDateString(date string) CSWeek {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		//---figure out what to do here
	}

	return GetWeek(d)
}

// GetNextWeek return the sequentially next week based on the passed in week
func GetNextWeek(w CSWeek) CSWeek {
	weeks := GetWeeks(w.End, w.End.Add(1*time.Hour*168))

	return weeks[1]
}

func getStartOfWeek(date time.Time) time.Time {
	for int(date.Weekday()) != int(time.Monday) {
		date = date.Add(-1 * time.Hour * 24)
	}

	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	return date
}
