package utils

import (
	"time"

	log "github.com/sirupsen/logrus"
)

var TIME_FORMAT = "2006-01-02"

func GetTimeFromString(timeInput string) *time.Time {
	t, err := time.Parse(TIME_FORMAT, timeInput)
	if err != nil {
		//TODO: fix this logic
		log.Error(err)
		t = time.Now()
	}

	return &t
}

func GetStringFromTime(input *time.Time) *string {
	if input == nil {
		return nil
	}

	output := input.Format(TIME_FORMAT)
	return &output
}
