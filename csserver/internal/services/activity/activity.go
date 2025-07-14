package activity

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/events"
	"csserver/internal/services/iam/appuser"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"csserver/internal/utils"
	"encoding/json"
	"fmt"
)

type ActivityTemplate struct {
	Subject   string
	GetLink   func(wrapper events.MessageWrapper) string
	GetDetail func(ctx context.Context, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, wrapper events.MessageWrapper) string
}

// LogActivity logs a given activity
func (s *ActivityService) LogActivity(
	ctx context.Context,
	us appuser.AppuserService,
	ps project.ProjectService,
	rs resource.ResourceService,
	subject string,
	data []byte) error {

	sub := events.GetCSSubject(subject)
	fmt.Println(sub)

	var wrapper events.MessageWrapper
	err := json.Unmarshal(data, &wrapper)
	if err != nil {
		return err
	}

	detail, link, err := getActivityDetail(sub, us, ps, rs, wrapper)
	if err != nil {
		return err
	}

	act := Activity{
		ControlFields: common.ControlFields{
			ID: utils.GenerateBase64UUID(),
		},
		Detail:       detail,
		Link:         link,
		Context:      subject,
		ActivityDate: wrapper.Timestamp,
		UserEmail:    wrapper.UserEmail,
	}

	_, err = s.CreateActivity(ctx, act)
	return err
}

// getActivityDetail return the detail for a given activity by looking up the proper template based on the key
func getActivityDetail(sub events.CSSubject, us appuser.AppuserService, ps project.ProjectService, rs resource.ResourceService, act events.MessageWrapper) (string, string, error) {
	key := sub.LookupKey()
	ctx := getContextFromSubject(sub)

	temp, ok := templateMap[key]
	if !ok {
		return "", "", fmt.Errorf("no template found for subject key %s", sub.LookupKey())
	}

	delta := temp.GetDetail(ctx, us, ps, rs, act)
	link := temp.GetLink(act)

	return delta, link, nil
}

// getContextFromSubject return a context with values specific to the org that generated the activity
func getContextFromSubject(sub events.CSSubject) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, config.OrgUrlKey, sub.OrgKey)

	return ctx
}
