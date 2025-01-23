package augment

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/services/resource/rtypes/resourcestatus"
	"csserver/internal/services/resource/rtypes/resourcetype"
)

func AugmentSchedule(schedule *idl.Schedule) {
	resourceList := findResources()
	projectList := findProjects()

	orgCapacity := 0

	for _, r := range *resourceList {
		if r.Type == resourcetype.Human && r.Status == resourcestatus.Inhouse {
			orgCapacity += r.AvailableHoursPerWeek
		}
	}

	schedule.Project = getProjectById(*projectList, schedule.ProjectID)

	for i, week := range schedule.ProjectActivityWeeks {
		for j, act := range week.Activities {
			resource := getResourceById(*resourceList, act.ResourceID)
			schedule.ProjectActivityWeeks[i].Activities[j].Resource = resource
			schedule.ProjectActivityWeeks[i].OrgCapacity = orgCapacity
		}
	}
}
