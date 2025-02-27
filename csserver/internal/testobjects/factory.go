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
	resourceMap := GetResourceMap()
	roleMap := GetRoleMap()
	org := GetTestOrganization()

	proj.CalculateProjectTaskStats(org, resourceMap, roleMap)

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
	resourceMap := GetResourceMap()
	roleMap := GetRoleMap()
	org := GetTestOrganization()

	proj.CalculateProjectTaskStats(org, resourceMap, roleMap)

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

// GetTestRoles return a deep copy of a role array for testing
func GetTestRoles() []resource.Role {
	roles, err := utils.DeepCopy[[]resource.Role](testRoles)
	if err != nil {
		panic(err)
	}

	return *roles
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

// convertRoleSliceToMap convert a slice of role into a map with the id as the key
func convertRoleSliceToMap(items []resource.Role) map[string]resource.Role {
	m := make(map[string]resource.Role)
	for _, r := range items {
		m[r.ID] = r
	}

	return m
}

// GetResourceMap return a map of resourceIDs to resource objects
func GetResourceMap() map[string]resource.Resource {
	resources := GetTestResources()
	return convertResourceSliceToMap(resources)
}

// GetRoleMap return a map of roles to resource objects
func GetRoleMap() map[string]resource.Role {
	resources := GetTestRoles()
	return convertRoleSliceToMap(resources)
}

// GetTestPortfolio return a portfolio of project schedule based on test projects
func GetTestPortfolio() portfolio.Portfolio {
	rm := GetResourceMap()
	ram := schedule.NewResourceAllocationMapFromResourceMap(rm)
	scheduleList := []schedule.Schedule{}
	projectList := []project.Project{}

	testProject := GetTestProjectWithCalcsDone("project:1")
	projectList = append(projectList, testProject)
	startDate := time.Date(2025, time.January, 21, 0, 0, 0, 0, time.UTC)
	sch, err := schedule.ScheduleProjectAlgo(&testProject, startDate, ram)
	if err != nil {
		panic(err)
	}

	scheduleList = append(scheduleList, sch)

	testProject2 := GetTestProjectWithCalcsDone("project:2")
	projectList = append(projectList, testProject2)
	startDate = time.Date(2025, time.March, 21, 0, 0, 0, 0, time.UTC)
	sch, err = schedule.ScheduleProjectAlgo(&testProject2, startDate, ram)
	if err != nil {
		panic(err)
	}

	scheduleList = append(scheduleList, sch)

	port := portfolio.NewPortfolio(projectList, rm)
	port.Schedule = scheduleList

	return port
}
