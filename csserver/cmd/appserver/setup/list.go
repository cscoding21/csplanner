package setup

import (
	"context"

	"csserver/internal/appserv/factory"
	"csserver/internal/services/list"

	log "github.com/sirupsen/logrus"
)

func CreateTestLists(ctx context.Context) error {
	service := factory.GetListService()
	//uuid.New().String()
	skillsList := list.List{
		Name:        "Skills",
		Description: "Skills are assigned to a resource.  This assignment allows for accurate project task assignment.",
		Values: []*list.ListItem{
			{Value: "backend", Name: "Backend Development"},
			{Value: "frontend", Name: "Frontend Development"},
			{Value: "database", Name: "Database Development"},
			{Value: "technical-architecture", Name: "Technical Architecture"},
			{Value: "devops", Name: "DevOps"},
			{Value: "ux", Name: "UX"},
			{Value: "ui", Name: "UI"},
			{Value: "product-design", Name: "Product Design"},
			{Value: "business-analysis", Name: "Business Analysis"},
			{Value: "project-management", Name: "Project Management"},
			{Value: "product-management", Name: "Product Management"},
			{Value: "requirements-gathering", Name: "Requirements Gathering"},
			{Value: "security", Name: "Security"},
			{Value: "marketing", Name: "Marketing"},
			{Value: "content-writing", Name: "Content Writing"},
			{Value: "video-editing", Name: "Video Editing"},
			{Value: "technical-writing", Name: "Technical Writing"},
			{Value: "data-analysis", Name: "Data Analysis"},
			{Value: "data-science", Name: "Data Science"},
			{Value: "communications", Name: "Communications"},
			{Value: "testing", Name: "Testing Automation"},
		},
	}

	fundingSourceList := list.List{
		Name:        "Funding Source",
		Description: "Funding sources are used to identify where a projects value comes from.  This list appears when assigning a projects value.",
		Values: []*list.ListItem{
			{Value: "internal", Name: "Internal"},
			{Value: "external", Name: "External"},
		},
	}

	valueCategoryList := list.List{
		Name:        "Value Catetgory",
		Description: "Value categories are used to characterize the nature of a projects value.  This information informs decisions about priority.",
		Values: []*list.ListItem{
			{Value: "revenue", Name: "Revenue increase"},
			{Value: "tax-benefit", Name: "Tax write-off"},
			{Value: "risk-mitigation", Name: "Risk mitigation"},
			{Value: "cost-reduction", Name: "Cost reduction"},
		},
	}

	_, err := service.GetList(ctx, "Skills")
	if err != nil {
		log.Errorf("Error getting skills list: %v", err)

		ur, err := service.CreateList(ctx, &skillsList)
		if err != nil {
			log.Errorf("Error creating skills list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	}

	_, err = service.GetList(ctx, "Funding Source")
	if err != nil {
		log.Errorf("Error getting funding source list: %v", err)

		ur, err := service.CreateList(ctx, &fundingSourceList)
		if err != nil {
			log.Errorf("Error creating funding source list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	}

	_, err = service.GetList(ctx, "Value Catetgory")
	if err != nil {
		log.Errorf("Error getting value category list: %v", err)

		ur, err := service.CreateList(ctx, &valueCategoryList)
		if err != nil {
			log.Errorf("Error creating value category list: %v", err)
		}

		if !ur.ValidationResult.Pass {
			log.Warn(ur.ValidationResult.Pass, ur.ValidationResult.Messages)
		}
	}

	return nil
}
