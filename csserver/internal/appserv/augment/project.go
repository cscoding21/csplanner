package augment

import (
	"csserver/internal/appserv/graph/idl"
)

func AugmentProject(proj *idl.Project) {
	if len(proj.ProjectMilestones) == 0 {
		return
	}

	skillList := getSkills()

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
		}
	}
}
