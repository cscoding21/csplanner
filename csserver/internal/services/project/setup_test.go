package project

import (
	"csserver/internal/common"
	"csserver/internal/services/organization"
	"csserver/internal/services/project/ptypes/featurepriority"
	"csserver/internal/services/project/ptypes/projectstatus"
	"csserver/internal/services/resource"
	"csserver/internal/services/resource/rtypes/resourcestatus"
	"csserver/internal/services/resource/rtypes/resourcetype"
	"csserver/internal/utils"

	"github.com/google/uuid"
)

var testProject = Project{
	ControlFields: common.ControlFields{
		ID: "project:123",
	},
	ProjectBasics: &ProjectBasics{
		Name:        "Test Project",
		Description: "This is a test project",
		OwnerID:     "user:123",
		Status:      projectstatus.Draft,
	},
	ProjectCost: &ProjectCost{
		Ongoing:     utils.ValToRef(100.0),
		BlendedRate: utils.ValToRef(200.0),
		Calculated:  ProjectCostCalculatedData{},
	},
	ProjectFeatures: []*ProjectFeature{
		{ID: utils.ValToRef("projectfeature:1"), Priority: featurepriority.High, Name: "Topic Sentence", Description: "It lays out the path", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:2"), Priority: featurepriority.VeryHigh, Name: "Intro", Description: "It builds the brand", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:3"), Priority: featurepriority.VeryHigh, Name: "Content", Description: "The meat of the video", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:4"), Priority: featurepriority.Medium, Name: "Outro", Description: "Like a linked list to other videos", Status: "proposed"},
	},
	ProjectValue: &ProjectValue{
		DiscountRate:   7.0,
		YearOneValue:   1000.0,
		YearTwoValue:   1000.0,
		YearThreeValue: 2000.0,
		YearFourValue:  1000.0,
		YearFiveValue:  4000.0,
		Calculated:     ProjectValueCalculatedData{},
	},
	ProjectMilestones: []*ProjectMilestone{
		{
			ID:         utils.ValToRef("milestone:1"),
			Calculated: ProjectMilestoneCalculatedData{},
			Phase: &ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Initiation",
				Description: "The project initiation phase is the first stage of turning an abstract idea into a meaningful goal. In this stage, you need to develop a business case and define the project on a broad level. In order to do that, you have to determine the need for the project and create a project charter",
				Order:       1,
			},
			Tasks: []*ProjectMilestoneTask{
				{
					ID:              utils.ValToRef("task:1"),
					Name:            "Stakeholder Alignment",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "backend",
					ResourceIDs:     []string{"resource:3"},
					Calculated:      ProjectTaskCalculatedData{},
				},
				{
					ID:              utils.ValToRef("task:2"),
					Name:            "Signoff",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "security",
					ResourceIDs:     []string{"resource:1"},
					Calculated:      ProjectTaskCalculatedData{},
				},
				{
					ID:              utils.ValToRef("task:3"),
					Name:            "gather-requirements",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "requirements-gathering",
					ResourceIDs:     []string{"resource:5"},
					Calculated:      ProjectTaskCalculatedData{},
				},
			},
		},

		{
			ID: utils.ValToRef("milestone:2"),
			Phase: &ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Planning",
				Description: "The project planning stage requires complete diligence as it lays out the project’s roadmap. Unless you are using a modern project management methodology like agile project management, the second phase of project management is expected to take almost half of the entire project’s timespan.",
				Order:       2,
			},
			Tasks: []*ProjectMilestoneTask{
				{
					ID:              utils.ValToRef("task:4"),
					Name:            "Technical Requirements Gathering",
					HourEstimate:    120,
					Status:          "new",
					RequiredSkillID: "devops",
					ResourceIDs:     []string{"resource:1", "resource:3"},
					Calculated:      ProjectTaskCalculatedData{},
				},
				{
					ID:              utils.ValToRef("task:5"),
					Name:            "Resourcing Plan",
					HourEstimate:    40,
					Status:          "new",
					RequiredSkillID: "project-management",
					ResourceIDs:     []string{"resource:1"},
					Calculated:      ProjectTaskCalculatedData{},
				},
				{
					ID:              utils.ValToRef("task:6"),
					Name:            "User Interface Design",
					HourEstimate:    160,
					Status:          "new",
					RequiredSkillID: "",
					ResourceIDs:     []string{},
					Calculated:      ProjectTaskCalculatedData{},
				},
			},
		},
	},
}

var testResources = []resource.Resource{
	{
		ControlFields: common.ControlFields{
			ID: "resource:1",
		},
		Role:           "CEO",
		Type:           resourcetype.Human,
		Name:           "Jeph",
		UserEmail:      utils.ValToRef("jeph@jmk21.com"),
		InitialCost:    0.0,
		AnnualizedCost: 200000,
		Status:         resourcestatus.Inhouse,
		Skills: []*resource.Skill{
			{ID: "devops", Proficiency: utils.ValToRef(1.0)},
			{ID: "security", Proficiency: utils.ValToRef(2.0)},
		},
	},
	{
		ControlFields: common.ControlFields{
			ID: "resource:2",
		},
		Role:           "Bartender",
		Type:           resourcetype.Human,
		Name:           "Tifa Lockhart",
		UserEmail:      utils.ValToRef("tifa@jmk21.com"),
		ProfileImage:   utils.ValToRef("/tifa.png"),
		InitialCost:    0.0,
		AnnualizedCost: 200000,
		Status:         resourcestatus.Inhouse,
		Skills: []*resource.Skill{
			{ID: "marketing", Proficiency: utils.ValToRef(1.0)},
			{ID: "content-writing", Proficiency: utils.ValToRef(2.0)},
			{ID: "video-editing", Proficiency: utils.ValToRef(3.0)},
		},
	},
	{
		ControlFields: common.ControlFields{
			ID: "resource:3",
		},
		Role:           "Ex-SOLDIER",
		Type:           resourcetype.Human,
		Name:           "Cloud Strife",
		UserEmail:      utils.ValToRef("cloud@jmk21.com"),
		ProfileImage:   utils.ValToRef("/cloud.png"),
		InitialCost:    0.0,
		AnnualizedCost: 150000,
		Status:         resourcestatus.Inhouse,
		Skills: []*resource.Skill{
			{ID: "devops", Proficiency: utils.ValToRef(2.0)},
			{ID: "backend", Proficiency: utils.ValToRef(1.0)},
			{ID: "database", Proficiency: utils.ValToRef(3.0)},
		},
	},
	{
		ControlFields: common.ControlFields{
			ID: "resource:4",
		},
		Role:           "Florist",
		Type:           resourcetype.Human,
		Name:           "Aerith Gainsborough",
		UserEmail:      utils.ValToRef("aerith@jmk21.com"),
		ProfileImage:   utils.ValToRef("/aerith.png"),
		InitialCost:    0.0,
		AnnualizedCost: 150000,
		Status:         resourcestatus.Inhouse,
		Skills: []*resource.Skill{
			{ID: "ui", Proficiency: utils.ValToRef(3.0)},
			{ID: "ux", Proficiency: utils.ValToRef(3.0)},
			{ID: "frontend", Proficiency: utils.ValToRef(3.0)},
		},
	},
	{
		ControlFields: common.ControlFields{
			ID: "resource:5",
		},
		Role:           "Avalache",
		Type:           resourcetype.Human,
		Name:           "Barret Wallace",
		ProfileImage:   utils.ValToRef("/barrett.png"),
		InitialCost:    0.0,
		AnnualizedCost: 150000,
		Status:         resourcestatus.Inhouse,
		Skills: []*resource.Skill{
			{ID: "project-management", Proficiency: utils.ValToRef(1.0)},
			{ID: "product-management", Proficiency: utils.ValToRef(2.0)},
			{ID: "requirements-gathering", Proficiency: utils.ValToRef(3.0)},
		},
	},
}

var testOrganization = organization.Organization{
	Name: "Test",
	Defaults: organization.OrganizationDefaults{
		FocusFactor:      5.0,
		CommsCoefficient: 5.0,
		HoursPerWeek:     40,
		DiscountRate:     7.0,
	},
}

// GetTestProject return a deep copy of a project graph for testing
func GetTestProject() Project {
	proj, err := utils.DeepCopy[Project](testProject)
	if err != nil {
		panic(err)
	}

	return *proj
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
