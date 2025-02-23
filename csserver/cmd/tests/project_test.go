package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"fmt"
	"testing"
)

func init() {
	config.InitConfig()
}

func TestSaveProject(t *testing.T) {
	ctx := getTestContext()

	service := factory.GetProjectService()
	org, err := factory.GetDefaultOrganization(ctx)
	if err != nil {
		t.Error(err)
	}

	rs := factory.GetResourceService()

	resourceMap, err := rs.GetResourceMap(ctx, false)
	if err != nil {
		t.Error(err)
	}

	roleMap, err := rs.GetRoleMap(ctx, false)
	if err != nil {
		t.Error(err)
	}

	put, err := service.GetProjectByID(ctx, "project:1")
	if err != nil {
		t.Error(err)
	}

	put.ProjectBasics.Name = put.ProjectBasics.Name + " T1"

	result, err := service.SaveProject(ctx, *put, resourceMap, roleMap, *org)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(result.Object.ID)
}
