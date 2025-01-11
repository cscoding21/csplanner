package augment

import "csserver/internal/appserv/graph/idl"

func AugmentSchedule(schedule *idl.ProjectSchedule) {
	resourceList := findResources()

	for i, week := range schedule.ProjectActivityWeeks {
		for j, act := range week.Activities {
			resource := getResourceById(*resourceList, *act.ResourceID)
			schedule.ProjectActivityWeeks[i].Activities[j].Resource = resource
		}
	}
}
