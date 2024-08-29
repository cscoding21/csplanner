package setup

import (
	"context"

	"csserver/internal/appserv/factory"
	listPackage "csserver/internal/services/list"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

func CreateTestLists(ctx context.Context) error {
	service := factory.GetListService()
	skillsList := listPackage.List{
		Name: "Skills",
		Values: []listPackage.ListItem{
			{Value: uuid.New().String(), Name: "Backend Development"},
			{Value: uuid.New().String(), Name: "Frontend Development"},
			{Value: uuid.New().String(), Name: "Database Development"},
			{Value: uuid.New().String(), Name: "DevOps"},
			{Value: uuid.New().String(), Name: "UX"},
			{Value: uuid.New().String(), Name: "UI"},
			{Value: uuid.New().String(), Name: "Project Management"},
			{Value: uuid.New().String(), Name: "Requirements Gathering"},
			{Value: uuid.New().String(), Name: "Security"},
			{Value: uuid.New().String(), Name: "Marketing"},
			{Value: uuid.New().String(), Name: "Content Writing"},
			{Value: uuid.New().String(), Name: "Video Editing"},
			{Value: uuid.New().String(), Name: "Technical Writing"},
			{Value: uuid.New().String(), Name: "Data Analysis"},
		},
	}

	fundingSourceList := listPackage.List{
		Name: "Funding Source",
		Values: []listPackage.ListItem{
			{Value: "internal", Name: "Internal"},
			{Value: "external", Name: "External"},
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

	return nil
}
