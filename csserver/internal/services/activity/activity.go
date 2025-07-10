package activity

import (
	"context"
	"csserver/internal/events"
	"encoding/json"
	"fmt"
)

// LogActivity logs a given activity
func (s *ActivityService) LogActivity(
	ctx context.Context,
	subject string,
	data []byte) error {

	sub := events.GetCSSubject(subject)
	fmt.Println(sub)

	var wrapper events.MessageWrapper
	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return err
	}

	detail, err := getActivityDetail(sub, wrapper)
	if err != nil {
		return err
	}

	act := Activity{
		Detail:    detail,
		Link:      subject,
		Context:   subject,
		UserEmail: wrapper.UserEmail,
	}

	_, err = s.CreateActivity(ctx, act)
	return err
}

func getActivityDetail(sub events.CSSubject, act events.MessageWrapper) (string, error) {

	return "", nil
}
