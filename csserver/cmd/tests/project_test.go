package tests

import (
	"csserver/internal/appserv/csmap"
	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/utils"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/wI2L/jsondiff"
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

	result, err := service.SaveProject(ctx, put.Data, resourceMap, roleMap, *org, utils.ValToRef("created"))
	if err != nil {
		t.Error(err)
	}

	p := common.UpwrapFromUpdateResult(result)
	fmt.Println(p.ID)
}

func TestGetProject(t *testing.T) {
	projectUnderTest := "zYfB9fG-SICFXH-ZpCV9vA"

	ctx := getTestContext()

	service := factory.GetProjectService(ctx)

	pro, err := service.GetProjectByID(ctx, projectUnderTest)
	if err != nil {
		t.Error(err)
	}

	put := csmap.ProjectProjectToIdl(pro.Data)

	csmap.AugmentProject(ctx, &pro.Data, &put)

	fmt.Println(pro.Data.ProjectBasics.Name)
}

func TestProjectDiff(t *testing.T) {
	ctx := getTestContext()

	service := factory.GetProjectService(ctx)
	put, err := service.GetProjectByID(ctx, "project:1")
	if err != nil {
		t.Error(err)
	}

	source := put.Data
	target, err := utils.DeepCopy(source)
	if err != nil {
		t.Error(err)
	}
	target.ID = fmt.Sprintf("%s-UPDATED", target.ID)
	target.ProjectBasics.Name = fmt.Sprintf("%s-UPDATED", target.ID)

	sourceJSON, err := json.Marshal(source)
	if err != nil {
		t.Error(err)
	}

	targetJSON, err := json.Marshal(target)
	if err != nil {
		t.Error(err)
	}

	diffs, err := jsondiff.CompareJSON(sourceJSON, targetJSON)
	if err != nil {
		t.Error(err)
	}

	for _, d := range diffs {
		fmt.Println(d)
	}
}

func TestSetProjectMilestoneFromTemplate(t *testing.T) {
	projectUnderTest := "y4pnyWsKQBag0QWyGwIdvg"
	templateID := "projecttemplate:simple"
	ctx := getTestContext()

	service := factory.GetProjectService(ctx)
	templateService := factory.GetProjectTemplateService(ctx)

	org, err := factory.GetDefaultOrganization(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	rs := factory.GetResourceService(ctx)

	resourceMap, err := rs.GetResourceMap(ctx, false)
	if err != nil {
		t.Error(err)
		return
	}

	roleMap, err := rs.GetRoleMap(ctx, false)
	if err != nil {
		t.Error(err)
		return
	}

	//---get the selected template
	template, err := templateService.GetProjecttemplateByID(ctx, templateID)
	if err != nil {
		t.Errorf("Error getting template: %s", err)
		return
	}

	pr, err := service.SetProjectMilestonesFromTemplate(ctx, projectUnderTest, template.Data, resourceMap, roleMap, *org)
	if err != nil {
		t.Errorf("Error setting template to project: %s", err)
		return
	}

	fmt.Println(pr.Object)
}
