package portfolio

import "time"

type CSWeek struct {
	Begin      time.Time
	End        time.Time
	Year       int
	WeekNumber int
}

func GetWeeks(start, end time.Time) []CSWeek {
	var weeks []CSWeek
	sow := GetStartOfWeek(start)

	for t := sow; t.Before(end); t = t.AddDate(0, 0, 7) {
		year, week := t.ISOWeek()
		eow := t.Add(1 * time.Hour * 24 * 7).Add(-1 * time.Second)
		w := CSWeek{
			Begin:      t,
			End:        eow,
			Year:       year,
			WeekNumber: week,
		}
		weeks = append(weeks, w)
	}
	return weeks
}

func GetStartOfWeek(date time.Time) time.Time {
	for int(date.Weekday()) != int(time.Monday) {
		date = date.Add(-1 * time.Hour * 24)
	}

	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	return date
}
