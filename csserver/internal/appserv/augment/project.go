package augment

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/services/project"
	"csserver/internal/services/project/ptypes/projectstatus"
)

func AugmentProject(model *project.Project, proj *idl.Project) {
	if len(proj.ProjectMilestones) == 0 {
		return
	}

	skillList := getSkills()
	resourceList := findResources()

	proj.ProjectBasics.Owner = getUserByEmail(model.ProjectBasics.OwnerID)

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
				idlSkill := idl.Skill{ID: skill.Value, Name: skill.Name}
				proj.ProjectMilestones[i].Tasks[j].Skills = append(proj.ProjectMilestones[i].Tasks[j].Skills, &idlSkill)
			}

			if len(t.ResourceIDs) > 0 {
				for _, r := range t.ResourceIDs {
					resource := getResourceById(*resourceList, r)
					proj.ProjectMilestones[i].Tasks[j].Resources = append(proj.ProjectMilestones[i].Tasks[j].Resources, resource)
				}
			}
		}
	}

	if model.ProjectDaci != nil {
		if len(model.ProjectDaci.DriverIDs) > 0 {
			for _, r := range model.ProjectDaci.DriverIDs {
				if r != nil {
					resource := getResourceById(*resourceList, *r)
					AugmentResource(resource)
					proj.ProjectDaci.Driver = append(proj.ProjectDaci.Driver, resource)
				}
			}
		}

		if len(model.ProjectDaci.ApproverIDs) > 0 {
			for _, r := range model.ProjectDaci.ApproverIDs {
				if r != nil {
					resource := getResourceById(*resourceList, *r)
					AugmentResource(resource)
					proj.ProjectDaci.Approver = append(proj.ProjectDaci.Approver, resource)
				}
			}
		}

		if len(model.ProjectDaci.ContributorIDs) > 0 {
			for _, r := range model.ProjectDaci.ContributorIDs {
				if r != nil {
					resource := getResourceById(*resourceList, *r)
					AugmentResource(resource)
					proj.ProjectDaci.Contributor = append(proj.ProjectDaci.Contributor, resource)
				}
			}
		}

		if len(model.ProjectDaci.InformedIDs) > 0 {
			for _, r := range model.ProjectDaci.InformedIDs {
				if r != nil {
					resource := getResourceById(*resourceList, *r)
					AugmentResource(resource)
					proj.ProjectDaci.Informed = append(proj.ProjectDaci.Informed, resource)
				}
			}
		}
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
