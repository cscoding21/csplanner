package augment

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/services/project"
)

func AugmentProject(model *project.Project, proj *idl.Project) {
	if len(proj.ProjectMilestones) == 0 {
		return
	}

	skillList := getSkills()
	resourceList := findResources()

	proj.ProjectBasics.Owner = getUserByEmail(model.ProjectBasics.OwnerID)

	for i, m := range proj.ProjectMilestones {
		if len(m.Tasks) == 0 {
			continue
		}

		for j, t := range m.Tasks {
			if len(t.RequiredSkillIDs) > 0 {
				for _, s := range t.RequiredSkillIDs {
					skill := getSkillById(*skillList, s)
					idlSkill := idl.Skill{ID: skill.Value, Name: skill.Name}
					proj.ProjectMilestones[i].Tasks[j].Skills = append(proj.ProjectMilestones[i].Tasks[j].Skills, &idlSkill)
				}
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
					proj.ProjectDaci.Driver = append(proj.ProjectDaci.Driver, resource)
				}
			}
		}

		if len(model.ProjectDaci.ApproverIDs) > 0 {
			for _, r := range model.ProjectDaci.ApproverIDs {
				if r != nil {
					resource := getResourceById(*resourceList, *r)
					proj.ProjectDaci.Approver = append(proj.ProjectDaci.Approver, resource)
				}
			}
		}

		if len(model.ProjectDaci.ContributorIDs) > 0 {
			for _, r := range model.ProjectDaci.ContributorIDs {
				if r != nil {
					resource := getResourceById(*resourceList, *r)
					proj.ProjectDaci.Contributor = append(proj.ProjectDaci.Contributor, resource)
				}
			}
		}

		if len(model.ProjectDaci.InformedIDs) > 0 {
			for _, r := range model.ProjectDaci.InformedIDs {
				if r != nil {
					resource := getResourceById(*resourceList, *r)
					proj.ProjectDaci.Informed = append(proj.ProjectDaci.Informed, resource)
				}
			}
		}
	}
}
