package testobjects

import (
	"csserver/internal/services/organization"
	"csserver/internal/services/portfolio"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"csserver/internal/services/schedule"
	"csserver/internal/utils"
	"time"
)

// GetTestProject return a deep copy of a project graph for testing
func GetTestProject(name string) project.Project {
	proj, err := utils.DeepCopy[project.Project](testProject)
	proj.ProjectBasics.Name = name
	proj.ID = name
	if err != nil {
		panic(err)
	}

	return *proj
}

// GetTestProjectWithCalcsDone return a test project with calculations done
func GetTestProjectWithCalcsDone(name string) project.Project {
	proj := GetTestProject(name)
	rm := GetResourceMap()
	org := GetTestOrganization()

	proj.CalculateProjectTasksStats(org, rm)

	return proj
}

// GetTestProject return a deep copy of a project graph for testing
func GetTestUpdateProject(name string) project.Project {
	proj, err := utils.DeepCopy[project.Project](updateProject)
	proj.ProjectBasics.Name = name
	proj.ID = name
	if err != nil {
		panic(err)
	}

	return *proj
}

// GetTestUpdateProjectWithCalcsDone
func GetTestUpdateProjectWithCalcsDone(name string) project.Project {
	proj := GetTestUpdateProject(name)
	rm := GetResourceMap()
	org := GetTestOrganization()

	proj.CalculateProjectTasksStats(org, rm)

	return proj
}

// GetTestResources return a deep copy of a resource array for testing
func GetTestResources() []resource.Resource {
	res, err := utils.DeepCopy[[]resource.Resource](testResources)
	if err != nil {
		panic(err)
	}

	return *res
}

// GetTestOrganization return a deep copy of an organization graph for testing
func GetTestOrganization() organization.Organization {
	org, err := utils.DeepCopy[organization.Organization](testOrganization)
	if err != nil {
		panic(err)
	}

	return *org
}

// ConvertResourceSliceToMap convert a slice of resource into a map with the id as the key
func convertResourceSliceToMap(resources []resource.Resource) map[string]resource.Resource {
	m := make(map[string]resource.Resource)
	for _, r := range resources {
		m[r.ID] = r
	}

	return m
}

// GetResourceMap return a map of resourceIDs to resource objects
func GetResourceMap() map[string]resource.Resource {
	resources := GetTestResources()
	return convertResourceSliceToMap(resources)
}

// GetTestPortfolio return a portfolio of project schedule based on test projects
func GetTestPortfolio() portfolio.Portfolio {
	rm := GetResourceMap()
	ram := schedule.NewResourceAllocationMapFromResourceMap(rm)
	portfolio := portfolio.Portfolio{Schedule: []schedule.Schedule{}}

	testProject := GetTestProjectWithCalcsDone("project:1")
	startDate := time.Date(2025, time.January, 21, 0, 0, 0, 0, time.UTC)
	sch, err := schedule.ScheduleProjectAlgo(&testProject, startDate, ram)
	if err != nil {
		panic(err)
	}

	portfolio.Schedule = append(portfolio.Schedule, sch)

	testProject2 := GetTestProjectWithCalcsDone("project:2")
	startDate = time.Date(2025, time.March, 21, 0, 0, 0, 0, time.UTC)
	sch, err = schedule.ScheduleProjectAlgo(&testProject2, startDate, ram)
	if err != nil {
		panic(err)
	}

	portfolio.Schedule = append(portfolio.Schedule, sch)

	return portfolio
}
