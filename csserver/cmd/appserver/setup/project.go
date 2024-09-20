package setup

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	"csserver/internal/appserv/factory"
	"csserver/internal/services/project"
	"csserver/internal/services/project/ptypes"
	"csserver/internal/utils"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

func CreateTestProjects(ctx context.Context) error {
	ps := factory.GetProjectService()
	rs := factory.GetResourceService()

	projectList, _ := ps.FindAllProjects(ctx)
	if len(projectList.Results) > 0 {
		return nil
	}

	allResources, _ := rs.FindAllResources(ctx)
	//allResources := ar.Results

	updateProject := project.Project{
		ProjectBasics: &project.ProjectBasics{
			Name:        "YouTube Sensation",
			Description: "This is my ticket to the big time",
			OwnerID:     allResources.Results[0].ID,
			Status:      ptypes.Draft,
		},
		ProjectCost: &project.ProjectCost{
			Ongoing:     utils.ValToRef(100.0),
			BlendedRate: utils.ValToRef(200.0),
		},
		ProjectFeatures: []*project.ProjectFeature{
			{ID: utils.ValToRef(uuid.New().String()), Priority: ptypes.High, Name: "Topic Sentence", Description: "It lays out the path", Status: "proposed"},
			{ID: utils.ValToRef(uuid.New().String()), Priority: ptypes.VeryHigh, Name: "Intro", Description: "It builds the brand", Status: "proposed"},
			{ID: utils.ValToRef(uuid.New().String()), Priority: ptypes.VeryHigh, Name: "Content", Description: "The meat of the video", Status: "proposed"},
			{ID: utils.ValToRef(uuid.New().String()), Priority: ptypes.Medium, Name: "Outro", Description: "Like a linked list to other videos", Status: "proposed"},
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
			Driver:      []*string{&allResources.Results[0].ID, &allResources.Results[1].ID},
			Approver:    []*string{&allResources.Results[3].ID, &allResources.Results[4].ID},
			Contributor: []*string{&allResources.Results[6].ID, &allResources.Results[7].ID, &allResources.Results[8].ID},
			Informed:    []*string{&allResources.Results[9].ID},
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
						ID:               utils.ValToRef(uuid.New().String()),
						Name:             "Stakeholder Alignment",
						Description:      "Hob-nobbing the stakeholders",
						HourEstimate:     120,
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{},
						StartDate:        nil,
						EndDate:          nil,
					},
					{
						ID:               utils.ValToRef(uuid.New().String()),
						Name:             "Signoff",
						Description:      "Get out your pen and sign off on the project charter",
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
				ID: utils.ValToRef(uuid.New().String()),
				Phase: &project.ProjectMilestonePhase{
					ID:          uuid.New().String(),
					Name:        "Planning",
					Description: "The project planning stage requires complete diligence as it lays out the project’s roadmap. Unless you are using a modern project management methodology like agile project management, the second phase of project management is expected to take almost half of the entire project’s timespan.",
					Order:       2,
				},
				Tasks: []*project.ProjectMilestoneTask{
					{
						ID:               utils.ValToRef(uuid.New().String()),
						Name:             "Technical Requirements Gathering",
						Description:      "Gather the requirements for the technical implementation of the project",
						HourEstimate:     120,
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{},
						StartDate:        nil,
						EndDate:          nil,
					},
					{
						ID:               utils.ValToRef(uuid.New().String()),
						Name:             "Resourcing Plan",
						Description:      "Assign resources tentatively throughout the project lifecycle",
						HourEstimate:     40,
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{},
						StartDate:        nil,
						EndDate:          nil,
					},
					{
						ID:               utils.ValToRef(uuid.New().String()),
						Name:             "User Interface Design",
						Description:      "Full mockups for the user interface",
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

	_, err := ps.CreateProject(ctx, &updateProject)
	if err != nil {
		log.Error(err)
	}

	otherProjects := findPortfolioProjects()

	for _, project := range otherProjects {
		_, err := ps.SaveProject(ctx, project)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

func findPortfolioProjects() []project.Project {
	outProjects := []project.Project{}

	type PT struct {
		name   string
		status ptypes.ProjectState
	}

	names := []PT{
		{name: "Video: Project Overview & Tech Stack", status: ptypes.Approved},
		{name: "Video: Kuberetes is Magic", status: ptypes.Approved},
		{name: "Video: Setting Local Dev With Tilt", status: ptypes.Approved},
		{name: "Video: Monorepo / Microservice", status: ptypes.Rejected},
		{name: "Video: Golang Project Setup", status: ptypes.Backlogged},
		{name: "Video: Mapping Objects in our Project", status: ptypes.Backlogged},
		{name: "Video: CRUD Operations Using Surreal", status: ptypes.Draft},
		{name: "Video: Scheduled 1", status: ptypes.Scheduled},
		{name: "Video: Scheduled 2", status: ptypes.Scheduled},
		{name: "Video: Inflight 1", status: ptypes.InFlight},
		{name: "Video: Inflight 2", status: ptypes.InFlight},
		{name: "Video: Inflight 3", status: ptypes.InFlight},
	}

	for _, p := range names {
		proj, _ := utils.DeepCopy[project.Project](GetVideoProjectTemplate(p.name, p.status))

		proj.ProjectValue.FundingSource = "external"
		proj.ProjectValue.YearOneValue = math.Round(1000.0 * rand.Float64())
		proj.ProjectValue.YearTwoValue = math.Round(1000.0 * rand.Float64())
		proj.ProjectValue.YearThreeValue = math.Round(2000.0 * rand.Float64())
		proj.ProjectValue.YearFourValue = math.Round(1000.0 * rand.Float64())
		proj.ProjectValue.YearFiveValue = math.Round(4000.0 * rand.Float64())

		outProjects = append(outProjects, *proj)
	}

	return outProjects
}

func GetVideoProjectTemplate(name string, status ptypes.ProjectState) project.Project {
	rs := factory.GetResourceService()
	ctx := context.Background()
	allResources, _ := rs.FindAllResources(ctx)

	startDate1 := time.Now().AddDate(0, 0, 1*rand.Intn(60))
	endDate1 := startDate1.AddDate(0, 0, 1*rand.Intn(4))

	startDate2 := endDate1.AddDate(0, 0, 1*rand.Intn(6))
	endDate2 := startDate2.AddDate(0, 0, 1*rand.Intn(4))

	startDate3 := endDate2.AddDate(0, 0, 1*rand.Intn(6))
	endDate3 := startDate3.AddDate(0, 0, 1*rand.Intn(4))

	startDate4 := endDate3.AddDate(0, 0, 1*rand.Intn(6))
	endDate4 := startDate4.AddDate(0, 0, 1*rand.Intn(4))

	startDate5 := endDate4.AddDate(0, 0, 1*rand.Intn(6))
	endDate5 := startDate5.AddDate(0, 0, 1*rand.Intn(4))

	startDate6 := endDate5.AddDate(0, 0, 1*rand.Intn(6))
	endDate6 := startDate6.AddDate(0, 0, 1*rand.Intn(4))

	proj := project.Project{
		ProjectBasics: &project.ProjectBasics{
			Name:        name,
			Description: "Here's a video to add to the portfolio",
			OwnerID:     allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].GetID(),
			Status:      status,
		},
		ProjectCost: &project.ProjectCost{
			Ongoing:     utils.ValToRef(0.0),
			BlendedRate: utils.ValToRef(121.0),
		},
		ProjectFeatures: []*project.ProjectFeature{
			{ID: utils.ValToRef("projectfeature:1"), Priority: ptypes.High, Name: "Topic Sentence", Description: "It lays out the path", Status: "proposed"},
			{ID: utils.ValToRef("projectfeature:2"), Priority: ptypes.VeryHigh, Name: "Intro", Description: "It builds the brand", Status: "proposed"},
			{ID: utils.ValToRef("projectfeature:3"), Priority: ptypes.VeryHigh, Name: "Content", Description: "The meat of the video", Status: "proposed"},
			{ID: utils.ValToRef("projectfeature:4"), Priority: ptypes.Medium, Name: "Outro", Description: "Like a linked list to other videos", Status: "proposed"},
		},
		ProjectValue: &project.ProjectValue{
			DiscountRate:   7.0,
			YearOneValue:   100.0,
			YearTwoValue:   100.0,
			YearThreeValue: 200.0,
			YearFourValue:  100.0,
			YearFiveValue:  400.0,
		},
		ProjectDaci: &project.ProjectDaci{
			Driver:      []*string{&allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
			Approver:    []*string{&allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID, &allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
			Contributor: []*string{&allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID, &allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
			Informed:    []*string{&allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
		},
		ProjectMilestones: []*project.ProjectMilestone{
			{
				ID:        utils.ValToRef(fmt.Sprintf("milestone:%s", uuid.New().String())),
				StartDate: &startDate1,
				EndDate:   &endDate6,
				Phase: &project.ProjectMilestonePhase{
					ID:          uuid.New().String(),
					Name:        "Execution",
					Description: "A single-phase project plan suitable for smaller initiatives.",
					Order:       1,
				},
				Tasks: []*project.ProjectMilestoneTask{
					{
						ID:               utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:             "Research",
						HourEstimate:     rand.Intn(10),
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
						StartDate:        &startDate1,
						EndDate:          &endDate1,
					},
					{
						ID:               utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:             "Video Outline",
						HourEstimate:     rand.Intn(11),
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
						StartDate:        &startDate2,
						EndDate:          &endDate2,
					},
					{
						ID:               utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:             "Video Shoot",
						HourEstimate:     rand.Intn(14),
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
						StartDate:        &startDate3,
						EndDate:          &endDate3,
					},
					{
						ID:               utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:             "Edit Video",
						HourEstimate:     rand.Intn(6),
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
						StartDate:        &startDate4,
						EndDate:          &endDate4,
					},
					{
						ID:               utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:             "Upload to All Platforms",
						HourEstimate:     rand.Intn(3),
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
						StartDate:        &startDate5,
						EndDate:          &endDate5,
					},
					{
						ID:               utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:             "Promote Upload",
						HourEstimate:     rand.Intn(3),
						Status:           "new",
						RequiredSkillIDs: []string{},
						ResourceIDs:      []string{allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
						StartDate:        &startDate6,
						EndDate:          &endDate6,
					},
				},
			},
		},
	}

	return proj
}
