package setup

import (
	"context"

	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/list"

	log "github.com/sirupsen/logrus"
)

func CreateTestLists(ctx context.Context) error {
	service := factory.GetListService(ctx)
	skillsList := list.List{
		ControlFields: common.ControlFields{
			ID: "list:skills",
		},
		Name:        list.ListNameSkills,
		Description: "Skills are assigned to a resource.  This assignment allows for accurate project task assignment.",
		Values: []*list.ListItem{
			{Value: "td:backend", Name: "Backend Development"},
			{Value: "td:frontend", Name: "Frontend Development"},
			{Value: "td:database", Name: "Database Development"},
			{Value: "td:technical-architecture", Name: "Technical Architecture"},
			{Value: "td:devops", Name: "DevOps"},
			{Value: "td:ux", Name: "UX"},
			{Value: "td:ui", Name: "UI"},
			{Value: "td:product-design", Name: "Product Design"},
			{Value: "td:business-analysis", Name: "Business Analysis"},
			{Value: "td:project-management", Name: "Project Management"},
			{Value: "td:product-management", Name: "Product Management"},
			{Value: "td:requirements-gathering", Name: "Requirements Gathering"},
			{Value: "td:security", Name: "Security"},
			{Value: "td:marketing", Name: "Marketing"},
			{Value: "td:content-writing", Name: "Content Writing"},
			{Value: "td:video-editing", Name: "Video Editing"},
			{Value: "td:technical-writing", Name: "Technical Writing"},
			{Value: "td:data-analysis", Name: "Data Analysis"},
			{Value: "td:data-science", Name: "Data Science"},
			{Value: "td:communications", Name: "Communications"},
			{Value: "td:testing", Name: "Testing Automation"},
		},
	}

	fundingSourceList := list.List{
		ControlFields: common.ControlFields{
			ID: "list:funding-source",
		},
		Name:        list.ListNameFundingSource,
		Description: "Funding sources are used to identify where a projects value comes from.  This list appears when assigning a projects value.",
		Values: []*list.ListItem{
			{Value: "internal", Name: "Internal"},
			{Value: "external", Name: "External"},
		},
	}

	valueCategoryList := list.List{
		ControlFields: common.ControlFields{
			ID: "list:value-cats",
		},
		Name:        list.ListNameValueCategory,
		Description: "Value categories are used to characterize the nature of a projects value.  This information informs decisions about priority.",
		Values: []*list.ListItem{
			{Value: "revenue", Name: "Revenue increase"},
			{Value: "tax-benefit", Name: "Tax write-off"},
			{Value: "risk-mitigation", Name: "Risk mitigation"},
			{Value: "cost-reduction", Name: "Cost reduction"},
		},
	}

	existingList, err := service.GetList(ctx, list.ListNameSkills)
	if existingList == nil {
		log.Errorf("Error getting skills list: %v", err)

		ur, err := service.CreateList(ctx, skillsList)
		if err != nil {
			log.Errorf("Error creating skills list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	} else {
		_, err := service.UpdateList(ctx, skillsList)
		if err != nil {
			log.Errorf("Error updating skills list: %v", err)
		}
	}

	existingList, err = service.GetList(ctx, list.ListNameFundingSource)
	if existingList == nil {
		log.Errorf("Error getting funding source list: %v", err)

		ur, err := service.CreateList(ctx, fundingSourceList)
		if err != nil {
			log.Errorf("Error creating funding source list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	} else {
		_, err := service.UpdateList(ctx, fundingSourceList)
		if err != nil {
			log.Errorf("Error updating funding source list: %v", err)
		}
	}

	existingList, err = service.GetList(ctx, list.ListNameValueCategory)
	if existingList == nil {
		log.Errorf("Error getting value category list: %v", err)

		ur, err := service.CreateList(ctx, valueCategoryList)
		if err != nil {
			log.Errorf("Error creating value category list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	} else {
		_, err := service.UpdateList(ctx, valueCategoryList)
		if err != nil {
			log.Errorf("Error updating value category list: %v", err)
		}
	}

	return nil
}
