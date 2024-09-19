package project

import (
	"csserver/internal/common"
	"csserver/internal/services/project/ptypes"
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
		Status:      ptypes.Draft,
	},
	ProjectCost: &ProjectCost{
		Ongoing:     utils.ValToRef(100.0),
		BlendedRate: utils.ValToRef(200.0),
	},
	ProjectFeatures: []*ProjectFeature{
		{ID: utils.ValToRef("projectfeature:1"), Priority: ptypes.High, Name: "Topic Sentence", Description: "It lays out the path", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:2"), Priority: ptypes.VeryHigh, Name: "Intro", Description: "It builds the brand", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:3"), Priority: ptypes.VeryHigh, Name: "Content", Description: "The meat of the video", Status: "proposed"},
		{ID: utils.ValToRef("projectfeature:4"), Priority: ptypes.Medium, Name: "Outro", Description: "Like a linked list to other videos", Status: "proposed"},
	},
	ProjectValue: &ProjectValue{
		DiscountRate:   7.0,
		YearOneValue:   1000.0,
		YearTwoValue:   1000.0,
		YearThreeValue: 2000.0,
		YearFourValue:  1000.0,
		YearFiveValue:  4000.0,
	},
	ProjectMilestones: []*ProjectMilestone{
		{
			ID: utils.ValToRef("milestone:1"),
			Phase: &ProjectMilestonePhase{
				ID:          uuid.New().String(),
				Name:        "Initiation",
				Description: "The project initiation phase is the first stage of turning an abstract idea into a meaningful goal. In this stage, you need to develop a business case and define the project on a broad level. In order to do that, you have to determine the need for the project and create a project charter",
				Order:       1,
			},
			Tasks: []*ProjectMilestoneTask{
				{
					ID:               utils.ValToRef("task:1"),
					Name:             "Stakeholder Alignment",
					HourEstimate:     120,
					Status:           "new",
					RequiredSkillIDs: []string{},
					ResourceIDs:      []string{},
					StartDate:        nil,
					EndDate:          nil,
				},
				{
					ID:               utils.ValToRef("task:2"),
					Name:             "Signoff",
					HourEstimate:     40,
					Status:           "new",
					RequiredSkillIDs: []string{},
					ResourceIDs:      []string{},
					StartDate:        nil,
					EndDate:          nil,
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
					ID:               utils.ValToRef("task:3"),
					Name:             "Technical Requirements Gathering",
					HourEstimate:     120,
					Status:           "new",
					RequiredSkillIDs: []string{},
					ResourceIDs:      []string{},
					StartDate:        nil,
					EndDate:          nil,
				},
				{
					ID:               utils.ValToRef("task:4"),
					Name:             "Resourcing Plan",
					HourEstimate:     40,
					Status:           "new",
					RequiredSkillIDs: []string{},
					ResourceIDs:      []string{},
					StartDate:        nil,
					EndDate:          nil,
				},
				{
					ID:               utils.ValToRef("task:5"),
					Name:             "User Interface Design",
					HourEstimate:     160,
					Status:           "new",
					RequiredSkillIDs: []string{},
					ResourceIDs:      []string{},
					StartDate:        nil,
					EndDate:          nil,
				},
			},
		},
	},
}
