package setup

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	"csserver/internal/appserv/factory"
	"csserver/internal/common"
	"csserver/internal/services/project"
	"csserver/internal/services/project/ptypes/featurepriority"
	"csserver/internal/services/project/ptypes/projectstatus"
	"csserver/internal/utils"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

func CreateTestProjects(ctx context.Context) error {
	ps := factory.GetProjectService()
	rs := factory.GetResourceService()
	us := factory.GetUserService()
	org, _ := factory.GetDefaultOrganization(ctx)

	projectList, _ := ps.FindAllProjects(ctx)
	if len(projectList.Results) > 0 {
		return nil
	}

	allResources, _ := rs.FindAllResources(ctx)
	allUsers, _ := us.FindAllUsers(ctx)

	updateProject := project.Project{
		ControlFields: common.ControlFields{
			ID: "project:1",
		},
		ProjectBasics: &project.ProjectBasics{
			Name:        "YouTube Sensation",
			Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			OwnerID:     allUsers.Results[0].Email,
			StartDate:   utils.ValToRef(time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC)),
		},
		ProjectStatusBlock: &project.ProjectStatusBlock{
			Status: projectstatus.Scheduled,
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
			DiscountRate: 7.0,
			ProjectValueLines: []*project.ProjectValueLine{
				{
					ID:             utils.ValToRef(uuid.New().String()),
					FundingSource:  "internal",
					ValueCategory:  "revenue",
					YearOneValue:   10000.0,
					YearTwoValue:   10000.0,
					YearThreeValue: 20000.0,
					YearFourValue:  10000.0,
					YearFiveValue:  40000.0,
				},
			},
		},
		ProjectDaci: &project.ProjectDaci{
			DriverIDs:      []*string{&allResources.Results[0].ID, &allResources.Results[1].ID},
			ApproverIDs:    []*string{&allResources.Results[3].ID, &allResources.Results[4].ID},
			ContributorIDs: []*string{&allResources.Results[6].ID, &allResources.Results[7].ID, &allResources.Results[8].ID},
			InformedIDs:    []*string{&allResources.Results[9].ID},
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

	pr, err := ps.SaveProject(ctx, updateProject, *org)
	if err != nil {
		log.Errorf("Save Project Error (YTS): %s", err)
	} else {
		log.Infof("PROJECT CREATED: %s", pr.Object.ProjectBasics.Name)
	}
	pr, err = ps.SetProjectStatus(ctx, pr.Object.ID, projectstatus.Scheduled, true)
	if err != nil {
		log.Errorf("Set Project Status Error (YTS): %s", err)
	} else {
		log.Infof("PROJECT STATUS UPDATED: %s", pr.Object.ProjectBasics.Name)
	}

	otherProjects := findPortfolioProjects()

	for _, otherProject := range otherProjects {
		st := otherProject.ProjectStatusBlock.Status
		pr, err = ps.SaveProject(ctx, otherProject, *org)
		if err != nil {
			log.Errorf("Save Project Error (OTHER): %s", err)
		} else {
			log.Infof("PROJECT CREATED: %s", pr.Object.ProjectBasics.Name)
		}
		pr, err = ps.SetProjectStatus(ctx, pr.Object.ID, st, true)
		if err != nil {
			log.Errorf("Set Project Status Error (YTS): %s", err)
		} else {
			log.Infof("PROJECT STATUS UPDATED: %s", pr.Object.ProjectBasics.Name)
		}
	}

	return nil
}

func findPortfolioProjects() []project.Project {
	outProjects := []project.Project{}

	type PT struct {
		name   string
		status projectstatus.ProjectState
	}

	names := []PT{
		{name: "Video: Project Overview & Tech Stack", status: projectstatus.Approved},
		{name: "Video: Kuberetes is Magic", status: projectstatus.Approved},
		{name: "Video: Setting Local Dev With Tilt", status: projectstatus.Approved},
		{name: "Video: Monorepo / Microservice", status: projectstatus.Rejected},
		{name: "Video: Golang Project Setup", status: projectstatus.Backlogged},
		{name: "Video: Mapping Objects in our Project", status: projectstatus.Backlogged},
		{name: "Video: CRUD Operations Using Surreal", status: projectstatus.Draft},
		{name: "Video: Scheduled 1", status: projectstatus.Scheduled},
		{name: "Video: Scheduled 2", status: projectstatus.Scheduled},
		{name: "Video: Inflight 1", status: projectstatus.InFlight},
		{name: "Video: Inflight 2", status: projectstatus.InFlight},
		{name: "Video: Inflight 3", status: projectstatus.InFlight},
	}

	for i, p := range names {
		proj, _ := utils.DeepCopy[project.Project](GetVideoProjectTemplate(p.name, p.status, (i + 2)))

		proj.ProjectValue.ProjectValueLines[0].YearOneValue = math.Round(1000.0 * rand.Float64())
		proj.ProjectValue.ProjectValueLines[0].YearTwoValue = math.Round(1000.0 * rand.Float64())
		proj.ProjectValue.ProjectValueLines[0].YearThreeValue = math.Round(2000.0 * rand.Float64())
		proj.ProjectValue.ProjectValueLines[0].YearFourValue = math.Round(1000.0 * rand.Float64())
		proj.ProjectValue.ProjectValueLines[0].YearFiveValue = math.Round(4000.0 * rand.Float64())

		outProjects = append(outProjects, *proj)
	}

	return outProjects
}

func GetVideoProjectTemplate(name string, status projectstatus.ProjectState, id int) project.Project {
	rs := factory.GetResourceService()
	us := factory.GetUserService()
	ctx := context.Background()
	allResources, _ := rs.FindAllResources(ctx)
	allUsers, _ := us.FindAllUsers(ctx)

	proj := project.Project{
		ControlFields: common.ControlFields{
			ID: fmt.Sprintf("project:%v", id),
		},
		ProjectBasics: &project.ProjectBasics{
			Name:        name,
			Description: "Here's a video to add to the portfolio",
			OwnerID:     allUsers.Results[rand.Intn(*allUsers.Pagination.TotalResults-1)].Email,
		},
		ProjectStatusBlock: &project.ProjectStatusBlock{
			Status: status,
		},
		ProjectCost: &project.ProjectCost{
			Ongoing:     utils.ValToRef(0.0),
			BlendedRate: utils.ValToRef(121.0),
		},
		ProjectFeatures: []*project.ProjectFeature{
			{ID: utils.ValToRef("projectfeature:1"), Priority: featurepriority.High, Name: "Topic Sentence", Description: "It lays out the path", Status: "proposed"},
			{ID: utils.ValToRef("projectfeature:2"), Priority: featurepriority.VeryHigh, Name: "Intro", Description: "It builds the brand", Status: "proposed"},
			{ID: utils.ValToRef("projectfeature:3"), Priority: featurepriority.VeryHigh, Name: "Content", Description: "The meat of the video", Status: "proposed"},
			{ID: utils.ValToRef("projectfeature:4"), Priority: featurepriority.Medium, Name: "Outro", Description: "Like a linked list to other videos", Status: "proposed"},
		},
		ProjectValue: &project.ProjectValue{
			DiscountRate: 7.0,
			ProjectValueLines: []*project.ProjectValueLine{
				{
					ID:             utils.ValToRef(uuid.New().String()),
					FundingSource:  "internal",
					ValueCategory:  "revenue",
					YearOneValue:   100.0,
					YearTwoValue:   100.0,
					YearThreeValue: 200.0,
					YearFourValue:  100.0,
					YearFiveValue:  400.0,
				},
			},
		},
		ProjectDaci: &project.ProjectDaci{
			DriverIDs:      []*string{&allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
			ApproverIDs:    []*string{&allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID, &allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
			ContributorIDs: []*string{&allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID, &allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
			InformedIDs:    []*string{&allResources.Results[rand.Intn(*allResources.Pagination.TotalResults-1)].ID},
		},
		ProjectMilestones: []*project.ProjectMilestone{
			{
				ID: utils.ValToRef(fmt.Sprintf("milestone:%s", uuid.New().String())),
				Phase: &project.ProjectMilestonePhase{
					ID:          uuid.New().String(),
					Name:        "Execution",
					Description: "A single-phase project plan suitable for smaller initiatives.",
					Order:       1,
				},
				Tasks: []*project.ProjectMilestoneTask{
					{
						ID:              utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:            "Research",
						HourEstimate:    rand.Intn(10),
						Status:          "new",
						RequiredSkillID: "business-analysis",
						ResourceIDs:     []string{"resource:yuffie"},
					},
					{
						ID:              utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:            "Video Outline",
						HourEstimate:    rand.Intn(11),
						Status:          "new",
						RequiredSkillID: "content-writing",
						ResourceIDs:     []string{"resource:tifa"},
					},
					{
						ID:              utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:            "Video Shoot",
						HourEstimate:    rand.Intn(14),
						Status:          "new",
						RequiredSkillID: "communications",
						ResourceIDs:     []string{"resource:barret"},
					},
					{
						ID:              utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:            "Edit Video",
						HourEstimate:    rand.Intn(6),
						Status:          "new",
						RequiredSkillID: "video-editing",
						ResourceIDs:     []string{"resource:jessie"},
					},
					{
						ID:              utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:            "Upload to All Platforms",
						HourEstimate:    rand.Intn(3),
						Status:          "new",
						RequiredSkillID: "frontend",
						ResourceIDs:     []string{"resource:wedge"},
					},
					{
						ID:              utils.ValToRef(fmt.Sprintf("task:%s", uuid.New().String())),
						Name:            "Promote Upload",
						HourEstimate:    rand.Intn(3),
						Status:          "new",
						RequiredSkillID: "marketing",
						ResourceIDs:     []string{"resource:wedge"},
					},
				},
			},
		},
	}

	return proj
}
