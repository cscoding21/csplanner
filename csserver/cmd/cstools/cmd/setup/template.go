package setup

import (
	"context"
	"fmt"

	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/projecttemplate"

	"github.com/google/uuid"
)

func CreateTestTemplates(ctx context.Context) error {
	ts := factory.GetProjectTemplateService(ctx)
	template := &projecttemplate.Projecttemplate{
		ControlFields: common.ControlFields{
			ID: "projecttemplate:primary",
		},
		Name:        "Primary",
		Description: "This template is used for complex projects that involve a large cross-functional effort.",
		Phases: []*projecttemplate.ProjecttemplatePhase{
			{
				ID:          uuid.New().String(),
				Name:        "Initiation",
				Description: "The project initiation phase is the first stage of turning an abstract idea into a meaningful goal. In this stage, you need to develop a business case and define the project on a broad level. In order to do that, you have to determine the need for the project and create a project charter",
				PhaseOrder:  1,
				Tasks: []*projecttemplate.ProjectTemplateTask{
					{Name: "Requirements gathering", Description: "Gather the requirements", RequiredSkillID: "td:requirements-gathering"},
					{Name: "Resource planning", Description: "Figure out who can do the thing", RequiredSkillID: "td:product-management"},
				},
			},
			{
				ID:          uuid.New().String(),
				Name:        "Planning",
				Description: "The project planning stage requires complete diligence as it lays out the project’s roadmap. Unless you are using a modern project management methodology like agile project management, the second phase of project management is expected to take almost half of the entire project’s timespan.",
				PhaseOrder:  2,
			},
			{
				ID:          uuid.New().String(),
				Name:        "Execution",
				Description: "The project execution stage is where your team does the actual work. As a project manager, your job is to establish efficient workflows and carefully monitor the progress of your team.",
				PhaseOrder:  3,
			},
			{
				ID:          uuid.New().String(),
				Name:        "Monitoring & Control",
				Description: "In the project management process, the third and fourth phases are not sequential in nature. The project monitoring and controlling phase run simultaneously with project execution, thereby ensuring that objectives and project deliverables are met.",
				PhaseOrder:  4,
			},
			{
				ID:          uuid.New().String(),
				Name:        "Completion",
				Description: "This is the final phase of the project management process. The project closure stage indicates the end of the project after the final delivery. There are times when external talent is hired specifically for the project on contract. Terminating these contracts and completing the necessary paperwork is also the responsibility of the project manager.",
				PhaseOrder:  5,
			},
		},
	}
	result, err := ts.SaveTemplate(ctx, *template)
	if err != nil {
		fmt.Printf("CreateTestTemplate Primary Error: %s\n", err)
	} else {
		fmt.Printf("CreateTestTemplate Primary: %v\n", result)
	}

	simpleTemplate := &projecttemplate.Projecttemplate{
		ControlFields: common.ControlFields{
			ID: "projecttemplate:simple",
		},
		Name:        "Simple",
		Description: "This template can be used for small projects that typically take less than 1 month to complete.",
		Phases: []*projecttemplate.ProjecttemplatePhase{
			{
				ID:          uuid.New().String(),
				Name:        "Execution",
				Description: "A single-phase project plan suitable for smaller initiatives.",
				PhaseOrder:  1,
				Tasks: []*projecttemplate.ProjectTemplateTask{
					{Name: "Backend development", Description: "Engineering work that takes place in the server-side components", RequiredSkillID: "td:backend"},
					{Name: "Testing", Description: "Complete testing of the new features as well as regression", RequiredSkillID: "td:product-management"},
				},
			},
		},
	}
	result, err = ts.SaveTemplate(ctx, *simpleTemplate)
	if err != nil {
		fmt.Printf("CreateTestTemplate Simple Error: %s\n", err)
	} else {
		fmt.Printf("CreateTestTemplate Simple: %v\n", result)
	}

	return nil
}
