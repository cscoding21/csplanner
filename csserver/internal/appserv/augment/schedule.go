package augment

import "csserver/internal/appserv/graph/idl"

func AugmentSchedule(schedule *idl.Schedule) {
	resourceList := findResources()
	projectList := findProjects()

	schedule.Project = getProjectById(*projectList, schedule.ProjectID)

	for i, week := range schedule.ProjectActivityWeeks {
		for j, act := range week.Activities {
			resource := getResourceById(*resourceList, act.ResourceID)
			schedule.ProjectActivityWeeks[i].Activities[j].Resource = resource
		}
	}
}
