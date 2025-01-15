package testobjects

import (
	"csserver/internal/common"
	"csserver/internal/services/project"
	"csserver/internal/services/project/ptypes/featurepriority"
	"csserver/internal/services/project/ptypes/projectstatus"
	"csserver/internal/utils"

	"github.com/google/uuid"
)

var testProject = project.Project{
	ControlFields: common.ControlFields{
		ID: "project:123",
	},
	ProjectBasics: &project.ProjectBasics{
		Name:        "Test Project",
		Description: "This is a test project",
		OwnerID:     "user:123",
		Status:      projectstatus.Draft,
	},
	ProjectCost: &project.ProjectCost{
		Ongoing:     utils.ValToRef(100.0),
		BlendedRate: utils.ValToRef(200.0),
		Calculated:  project.ProjectCostCalculatedData{},
	},
	ProjectFeatures: []*project.ProjectFeature{
		{ID: utils.ValToRef("projectfeature:1"), Priority: featurepriority.High, Name: "Topic Sentence", Description: "It lays out the path", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:2"), Priority: featurepriority.VeryHigh, Name: "Intro", Description: "It builds the brand", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:3"), Priority: featurepriority.VeryHigh, Name: "Content", Description: "The meat of the video", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:4"), Priority: featurepriority.Medium, Name: "Outro", Description: "Like a linked list to other videos", Status: "proposed"},
	},
	ProjectValue: &project.ProjectValue{
		DiscountRate:   7.0,
		YearOneValue:   1000.0,
		YearTwoValue:   1000.0,
		YearThreeValue: 2000.0,
		YearFourValue:  1000.0,
		YearFiveValue:  4000.0,
		Calculated:     project.ProjectValueCalculatedData{},
	},
	ProjectMilestones: []*project.ProjectMilestone{
		{
			ID:         utils.ValToRef("milestone:1"),
			Calculated: project.ProjectMilestoneCalculatedData{},
			Phase: &project.ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Initiation",
				Description: "The project initiation phase is the first stage of turning an abstract idea into a meaningful goal. In this stage, you need to develop a business case and define the project on a broad level. In order to do that, you have to determine the need for the project and create a project charter",
				Order:       1,
			},
			Tasks: []*project.ProjectMilestoneTask{
				{
					ID:              utils.ValToRef("task:1"),
					Name:            "Stakeholder Alignment",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "backend",
					ResourceIDs:     []string{"resource:3"},
					Calculated:      project.ProjectTaskCalculatedData{},
				},
				{
					ID:              utils.ValToRef("task:2"),
					Name:            "Signoff",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "security",
					ResourceIDs:     []string{"resource:1"},
					Calculated:      project.ProjectTaskCalculatedData{},
				},
				{
					ID:              utils.ValToRef("task:3"),
					Name:            "Gather Requirements",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "requirements-gathering",
					ResourceIDs:     []string{"resource:5"},
					Calculated:      project.ProjectTaskCalculatedData{},
				},
			},
		},

		{
			ID: utils.ValToRef("milestone:2"),
			Phase: &project.ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Planning",
				Description: "The project planning stage requires complete diligence as it lays out the project’s roadmap. Unless you are using a modern project management methodology like agile project management, the second phase of project management is expected to take almost half of the entire project’s timespan.",
				Order:       2,
			},
			Tasks: []*project.ProjectMilestoneTask{
				{
					ID:              utils.ValToRef("task:4"),
					Name:            "Technical Requirements Gathering",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "devops",
					ResourceIDs:     []string{"resource:1", "resource:3"},
					Calculated:      project.ProjectTaskCalculatedData{},
				},
				{
					ID:              utils.ValToRef("task:5"),
					Name:            "Resourcing Plan",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "project-management",
					ResourceIDs:     []string{"resource:1"},
					Calculated:      project.ProjectTaskCalculatedData{},
				},
				{
					ID:              utils.ValToRef("task:6"),
					Name:            "User Interface Design",
					HourEstimate:    160,
					Status:          "new",
					RequiredSkillID: "",
					ResourceIDs:     []string{},
					Calculated:      project.ProjectTaskCalculatedData{},
				},
			},
		},
	},
}

var updateProject = project.Project{
	ControlFields: common.ControlFields{
		ID: "project:1",
	},
	ProjectBasics: &project.ProjectBasics{
		Name:        "YouTube Sensation",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		OwnerID:     "user:123",
		Status:      projectstatus.Draft,
	},
	ProjectCost: &project.ProjectCost{
		Ongoing:     utils.ValToRef(100.0),
		BlendedRate: utils.ValToRef(200.0),
	},
	ProjectFeatures: []*project.ProjectFeature{
		{ID: utils.ValToRef(uuid.New().String()), Priority: featurepriority.High, Name: "Topic Sentence", Description: "It lays out the path", Status: "proposed"},
		{ID: utils.ValToRef(uuid.New().String()), Priority: featurepriority.VeryHigh, Name: "Intro", Description: "It builds the brand", Status: "proposed"},
		{ID: utils.ValToRef(uuid.New().String()), Priority: featurepriority.VeryHigh, Name: "Content", Description: "The meat of the video", Status: "proposed"},
		{ID: utils.ValToRef(uuid.New().String()), Priority: featurepriority.Medium, Name: "Outro", Description: "Like a linked list to other videos", Status: "proposed"},
	},
	ProjectValue: &project.ProjectValue{
		FundingSource:  "internal",
		DiscountRate:   7.0,
		YearOneValue:   10000.0,
		YearTwoValue:   10000.0,
		YearThreeValue: 20000.0,
		YearFourValue:  10000.0,
		YearFiveValue:  40000.0,
	},
	ProjectDaci: &project.ProjectDaci{
		DriverIDs:      []*string{common.ValToRef("jeph@jmk21.com")},
		ApproverIDs:    []*string{common.ValToRef("jeph@jmk21.com")},
		ContributorIDs: []*string{common.ValToRef("jeph@jmk21.com")},
		InformedIDs:    []*string{common.ValToRef("jeph@jmk21.com")},
	},
	ProjectMilestones: []*project.ProjectMilestone{
		{
			ID: utils.ValToRef(uuid.New().String()),
			Phase: &project.ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Initiation",
				Description: "The project initiation phase is the first stage of turning an abstract idea into a meaningful goal. In this stage, you need to develop a business case and define the project on a broad level. In order to do that, you have to determine the need for the project and create a project charter",
				Order:       1,
			},
			Tasks: []*project.ProjectMilestoneTask{
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Stakeholder Alignment",
					Description:     "Hob-nobbing the stakeholders",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "project-management",
					ResourceIDs:     []string{"resource:barret"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Signoff",
					Description:     "Get out your pen and sign off on the project charter",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "project-management",
					ResourceIDs:     []string{"resource:barret"},
				},
			},
		},
		{
			ID: utils.ValToRef(uuid.New().String()),
			Phase: &project.ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Planning",
				Description: "The project planning stage requires complete diligence as it lays out the project’s roadmap. Unless you are using a modern project management methodology like agile project management, the second phase of project management is expected to take almost half of the entire project’s timespan.",
				Order:       2,
			},
			Tasks: []*project.ProjectMilestoneTask{
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Technical Requirements Gathering",
					Description:     "Gather the requirements for the technical implementation of the project",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "requirements-gathering",
					ResourceIDs:     []string{"resource:yuffie"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Resourcing Plan",
					Description:     "Assign resources tentatively throughout the project lifecycle",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "project-management",
					ResourceIDs:     []string{"resource:biggs"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "User Interface Design",
					Description:     "Full mockups for the user interface",
					HourEstimate:    160,
					Status:          "new",
					RequiredSkillID: "ux",
					ResourceIDs:     []string{"resource:aerith"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Technical Architecture",
					Description:     "Design the system components and integrations",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "technical-architecture",
					ResourceIDs:     []string{"resource:cid", "resource:zack"},
				},
			},
		},

		{
			ID: utils.ValToRef(uuid.New().String()),
			Phase: &project.ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Execution",
				Description: "The project execution stage is where your team does the actual work. As a project manager, your job is to establish efficient workflows and carefully monitor the progress of your team.",
				Order:       3,
			},
			Tasks: []*project.ProjectMilestoneTask{
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Database Development",
					Description:     "Write the schema, tables, and database objects",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "database",
					ResourceIDs:     []string{"resource:zack"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Backend Development",
					Description:     "Write the backend coding artifacts",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "backend",
					ResourceIDs:     []string{"resource:cloud"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Frontend Development",
					Description:     "Write the frontend coding artifacts",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "frontend",
					ResourceIDs:     []string{"resource:wedge"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Testing",
					Description:     "Make sure all of the code works",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "ui",
					ResourceIDs:     []string{"resource:barret"},
				},
			},
		},

		{
			ID: utils.ValToRef(uuid.New().String()),
			Phase: &project.ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Monitoring & Control",
				Description: "In the project management process, the third and fourth phases are not sequential in nature. The project monitoring and controlling phase run simultaneously with project execution, thereby ensuring that objectives and project deliverables are met.",
				Order:       4,
			},
			Tasks: []*project.ProjectMilestoneTask{
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Ongoing Project Management",
					Description:     "Make sure the project is moving forward",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "project-management",
					ResourceIDs:     []string{"resource:barret"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Documentation",
					Description:     "Write the project diary",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "technical-writing",
					ResourceIDs:     []string{"resource:wedge"},
				},
			},
		},

		{
			ID: utils.ValToRef(uuid.New().String()),
			Phase: &project.ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Completion",
				Description: "This is the final phase of the project management process. The project closure stage indicates the end of the project after the final delivery. There are times when external talent is hired specifically for the project on contract. Terminating these contracts and completing the necessary paperwork is also the responsibility of the project manager.",
				Order:       5,
			},
			Tasks: []*project.ProjectMilestoneTask{
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Final Signoff",
					Description:     "Get the approvals",
					HourEstimate:    16,
					Status:          "new",
					RequiredSkillID: "project-management",
					ResourceIDs:     []string{"resource:barret"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Stakeholder Communication",
					Description:     "Write the project diary",
					HourEstimate:    8,
					Status:          "new",
					RequiredSkillID: "project-management",
					ResourceIDs:     []string{"resource:barret"},
				},
				{
					ID:              utils.ValToRef(uuid.New().String()),
					Name:            "Tell the World",
					Description:     "Let everybody know of your triumph",
					HourEstimate:    160,
					Status:          "new",
					RequiredSkillID: "marketing",
					ResourceIDs:     []string{"resource:tifa"},
				},
			},
		},
	},
}
