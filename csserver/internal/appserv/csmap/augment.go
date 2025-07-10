package csmap

import (
	"context"
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/calendar"
	"csserver/internal/common"
	"csserver/internal/services/list"
	"csserver/internal/services/project"
	"csserver/internal/services/project/ptypes/projectstatus"
	"csserver/internal/services/resource/rtypes/resourcestatus"
	"csserver/internal/services/resource/rtypes/resourcetype"

	log "github.com/sirupsen/logrus"
)

func AugmentComment(ctx context.Context, comment *idl.CommentEnvelope) {
	if comment.Data.Replies != nil {
		AugmentCommentSlice(ctx, &comment.Data.Replies)
	}
}

func AugmentBaseModel(ctx context.Context, base *idl.BaseModel) {
	if base == nil {
		return
	}

	base.CreateByUser = getUserByEmail(ctx, base.CreatedBy)
	base.UpdateByUser = getUserByEmail(ctx, base.UpdatedBy)

	if base.DeletedBy != nil {
		base.DeleteByUser = getUserByEmail(ctx, *base.DeletedBy)
	}
}

func AugmentCommentSlice(ctx context.Context, comments *[]*idl.CommentEnvelope) {
	for _, comment := range *comments {
		AugmentComment(ctx, comment)
	}
}

func AugmentOrganization(ctx context.Context, org *idl.Organization) {
	if org == nil {
		return
	}

	templates := findTemplates(ctx)
	roles := findRoles(ctx)
	resources := findResources(ctx)
	lists := findLists(ctx)

	sl := getListById(*lists, list.ListNameSkills)
	fsl := getListById(*lists, list.ListNameFundingSource)
	vcl := getListById(*lists, list.ListNameValueCategory)

	org.Setup = &idl.OrganizationSetup{
		HasRoles:           len(*roles) > 0,
		HasTemplates:       len(*templates) > 0,
		HasResources:       len(*resources) > 0,
		HasSkills:          len(sl.Values) > 0,
		HasFundingSources:  len(fsl.Values) > 0,
		HasValueCategories: len(vcl.Values) > 0,
	}

	org.Setup.IsReadyForProjects = (org.Setup.HasRoles &&
		org.Setup.HasTemplates &&
		org.Setup.HasResources &&
		org.Setup.HasSkills &&
		org.Setup.HasFundingSources &&
		org.Setup.HasValueCategories)

}

// AugmentActivity take an idl.Activity object and add the referenece user details
func AugmentActivity(ctx context.Context, act *idl.Activity) {
	u := getUserByEmail(ctx, act.UserEmail)
	act.User = u
}

func AugmentActivitySlice(ctx context.Context, activities []*idl.Activity) {
	for i := range activities {
		AugmentActivity(ctx, activities[i])
	}
}

// AugmentPortfolio enhance calculate and caches properties of a portfolio
func AugmentPortfolio(ctx context.Context, port *idl.Portfolio, resourceID *string) {
	//---exit conditions
	if port == nil || port.Begin == nil || port.End == nil {
		return
	}

	resourceList := findResources(ctx)

	orgCapacity := 0

	for _, r := range *resourceList {
		if resourceID != nil && r.ID != *resourceID {
			continue
		}

		if resourceID != nil && r.Type == resourcetype.Human {
			orgCapacity += r.AvailableHoursPerWeek
		} else if r.Type == resourcetype.Human && r.Status == resourcestatus.Inhouse {
			orgCapacity += r.AvailableHoursPerWeek
		}
	}

	weeks := calendar.GetWeeks(*port.Begin, *port.End)

	portWeeks := []*idl.PortfolioWeekSummary{}

	for _, w := range weeks {
		thisWeek := idl.PortfolioWeekSummary{
			WeekNumber:     w.WeekNumber,
			Year:           w.Year,
			Begin:          w.Begin,
			End:            w.End,
			OrgCapacity:    orgCapacity,
			AllocatedHours: getWeekHourSummary(common.RefToValSlice(port.Schedule), w),
		}

		portWeeks = append(portWeeks, &thisWeek)
	}

	port.WeekSummary = portWeeks
	calculated := getPortfolioCalculatedData(port)

	port.Calculated = &calculated
}

func getWeekHourSummary(sch []idl.Schedule, week calendar.CSWeek) int {
	out := 0
	for _, s := range sch {
		for _, w := range s.ProjectActivityWeeks {
			if w.End.Equal(week.End) {
				for _, a := range w.Activities {
					out += a.HoursSpent
				}
			}
		}
	}

	return out
}

func getPortfolioCalculatedData(port *idl.Portfolio) idl.PortfolioCalculatedData {
	out := idl.PortfolioCalculatedData{
		ValueInFlight:  0.0,
		ValueScheduled: 0.0,
		TotalValue:     0.0,
		CountInFlight:  0,
		CountScheduled: 0,
		TotalCount:     0,
	}

	if port == nil {
		return out
	}

	for _, s := range port.Schedule {
		out.TotalCount++
		out.TotalValue += *s.Project.ProjectValue.Calculated.NetPresentValue
		if s.Project.ProjectStatusBlock.Status == "inflight" {
			out.CountInFlight++
			out.ValueInFlight += *s.Project.ProjectValue.Calculated.NetPresentValue
		} else if s.Project.ProjectStatusBlock.Status == "scheduled" {
			out.CountScheduled++
			out.ValueScheduled += *s.Project.ProjectValue.Calculated.NetPresentValue
		}
	}

	return out
}

// AugmentProject add extra details to the properties of a project
func AugmentProject(ctx context.Context, model *project.Project, proj *idl.Project) {
	if proj == nil || model == nil {
		return
	}

	if len(proj.ProjectMilestones) == 0 {
		return
	}

	skillList := getSkills(ctx)
	resourceList := findResources(ctx)

	proj.ProjectBasics.Owner = getUserByEmail(ctx, model.ProjectBasics.OwnerID)

	team := make(map[string]idl.Resource)

	for i, s := range proj.ProjectStatusBlock.AllowedNextStates {
		tra := getStateTransition(*model, projectstatus.ProjectState(s.NextState))
		messages := []*idl.ValidationMessage{}

		for _, m := range tra.CheckResults.Messages {
			messages = append(messages, &idl.ValidationMessage{
				Message: m.Message,
				Field:   m.Field,
			})
		}

		proj.ProjectStatusBlock.AllowedNextStates[i].CheckResult = &idl.ValidationResult{
			Pass:     tra.CheckResults.Pass,
			Messages: messages,
		}
	}

	for i, m := range proj.ProjectMilestones {
		if len(m.Tasks) == 0 {
			continue
		}

		for j, t := range m.Tasks {
			if len(t.RequiredSkillID) > 0 {
				skill := getSkillById(*skillList, t.RequiredSkillID)
				if skill != nil {
					idlSkill := idl.Skill{ID: skill.Value, Name: skill.Name}
					proj.ProjectMilestones[i].Tasks[j].Skills = append(proj.ProjectMilestones[i].Tasks[j].Skills, &idlSkill)
				} else {
					log.Warnf("Skill %s not found when augmenting task %s", t.RequiredSkillID, t.Name)
				}
			}

			if len(t.ResourceIDs) > 0 {
				for _, r := range t.ResourceIDs {
					resource := getResourceById(ctx, *resourceList, r)

					if resource != nil {
						team[r] = *resource
						proj.ProjectMilestones[i].Tasks[j].Resources = append(proj.ProjectMilestones[i].Tasks[j].Resources, resource)
					}
				}
			}
		}
	}

	if model.ProjectDaci != nil {
		if len(model.ProjectDaci.DriverIDs) > 0 {
			for _, r := range model.ProjectDaci.DriverIDs {
				if r != nil {
					resource := getResourceById(ctx, *resourceList, *r)
					// AugmentResource(ctx, resource)
					team[*r] = *resource
					proj.ProjectDaci.Driver = append(proj.ProjectDaci.Driver, resource)
				}
			}
		}

		if len(model.ProjectDaci.ApproverIDs) > 0 {
			for _, r := range model.ProjectDaci.ApproverIDs {
				if r != nil {
					resource := getResourceById(ctx, *resourceList, *r)
					// AugmentResource(ctx, resource)
					team[*r] = *resource
					proj.ProjectDaci.Approver = append(proj.ProjectDaci.Approver, resource)
				}
			}
		}

		if len(model.ProjectDaci.ContributorIDs) > 0 {
			for _, r := range model.ProjectDaci.ContributorIDs {
				if r != nil {
					resource := getResourceById(ctx, *resourceList, *r)
					// AugmentResource(ctx, resource)
					team[*r] = *resource
					proj.ProjectDaci.Contributor = append(proj.ProjectDaci.Contributor, resource)
				}
			}
		}

		if len(model.ProjectDaci.InformedIDs) > 0 {
			for _, r := range model.ProjectDaci.InformedIDs {
				if r != nil {
					resource := getResourceById(ctx, *resourceList, *r)
					// AugmentResource(ctx, resource)
					team[*r] = *resource
					proj.ProjectDaci.Informed = append(proj.ProjectDaci.Informed, resource)
				}
			}
		}
	}

	for _, v := range team {
		proj.Calculated.TeamMembers = append(proj.Calculated.TeamMembers, &v)
	}
}

func getStateTransition(model project.Project, status projectstatus.ProjectState) *project.ProjectStatusTransition {
	for _, pst := range model.ProjectStatusBlock.AllowedNextStates {
		if pst.NextState == status {
			return pst
		}
	}

	return nil
}

// AugmentResource fills in object data for a single reaource
func AugmentResource(ctx context.Context, res *idl.Resource) {
	if res == nil {
		return
	}

	if res.UserEmail != nil {
		res.User = getUserByEmail(ctx, *res.UserEmail)
	}

	if res.RoleID != nil {
		res.Role = getRoleById(ctx, *res.RoleID)
	}

	if len(res.Skills) == 0 {
		return
	}

	skills := getSkills(ctx)

	for i, s := range res.Skills {
		li := getSkillById(*skills, *s.SkillID)

		if li == nil {
			log.Warnf("skill not found: %s for resource %s", *s.SkillID, res.Name)
		} else {
			res.Skills[i].Name = li.Name
		}
	}
}

// AugmentRole ensure the the skill name is inluded in the role's default skills
func AugmentRole(ctx context.Context, r *idl.Role) {
	if r == nil {
		return
	}

	skills := getSkills(ctx)

	if len(r.DefaultSkills) == 0 {
		return
	}

	for j, sk := range r.DefaultSkills {
		skill := getSkillById(*skills, *sk.SkillID)
		if skill != nil {
			r.DefaultSkills[j].Name = skill.Name
		} else {
			log.Warnf("Skill %s not found when augmenting role %s", *sk.SkillID, r.Name)
		}

	}
}

func AugmentSchedule(ctx context.Context, schedule *idl.Schedule) {
	if schedule == nil {
		return
	}

	resourceList := findResources(ctx)
	projectList := findProjects(ctx)

	orgCapacity := 0

	for _, r := range *resourceList {
		if r.Type == resourcetype.Human && r.Status == resourcestatus.Inhouse {
			orgCapacity += r.AvailableHoursPerWeek
		}
	}

	schedule.Project = getProjectById(ctx, *projectList, schedule.ProjectID)

	for i, week := range schedule.ProjectActivityWeeks {
		for j, act := range week.Activities {
			resource := getResourceById(ctx, *resourceList, act.ResourceID)
			schedule.ProjectActivityWeeks[i].Activities[j].Resource = resource
			schedule.ProjectActivityWeeks[i].OrgCapacity = orgCapacity
		}
	}
}
