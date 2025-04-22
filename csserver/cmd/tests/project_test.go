package tests

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/config"
	"fmt"
	"testing"
)

func init() {
	config.InitConfig()
}

func TestSaveProject(t *testing.T) {
	ctx := getTestContext()

	service := factory.GetProjectService(ctx)
	org, err := factory.GetDefaultOrganization(ctx)
	if err != nil {
		t.Error(err)
	}

	rs := factory.GetResourceService(ctx)

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

	put.Data.ProjectBasics.Name = put.Data.ProjectBasics.Name + " T1"

	result, err := service.SaveProject(ctx, put.Data, resourceMap, roleMap, *org)
	if err != nil {
		t.Error(err)
	}

	p := common.UpwrapFromUpdateResult(result)
	fmt.Println(p.ID)
}

func TestGetProject(t *testing.T) {
	projectUnderTest := "project:1"

	ctx := getTestContext()

	service := factory.GetProjectService(ctx)

	pro, err := service.GetProjectByID(ctx, projectUnderTest)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(pro.Data.ProjectBasics.Name)
}
