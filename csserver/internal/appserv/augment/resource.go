package augment

import (
	"csserver/internal/appserv/graph/idl"
)

// AugmentResource fills in object data for a single reaource
func AugmentResource(res *idl.Resource) {
	if len(res.Skills) == 0 {
		return
	}

	skills := getSkills()

	for i, s := range res.Skills {
		li := getSkillById(*skills, s.ID)
		res.Skills[i].Name = li.Name
	}
}
