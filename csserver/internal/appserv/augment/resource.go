package augment

import (
	"csserver/internal/appserv/graph/idl"

	log "github.com/sirupsen/logrus"
)

// AugmentResource fills in object data for a single reaource
func AugmentResource(res *idl.Resource) {
	if res.UserEmail != nil {
		res.User = getUserByEmail(*res.UserEmail)
	}

	if res.RoleID != nil {
		res.Role = getRoleById(*res.RoleID)
	}

	if len(res.Skills) == 0 {
		return
	}

	skills := getSkills()

	for i, s := range res.Skills {
		li := getSkillById(*skills, s.ID)

		if li == nil {
			log.Warnf("skill not found: %s", s.ID)
		} else {
			res.Skills[i].Name = li.Name
		}
	}
}
