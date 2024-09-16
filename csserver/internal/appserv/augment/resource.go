package augment

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/services/list"
	"strings"

	log "github.com/sirupsen/logrus"
)

// AugmentResource fills in object data for a single reaource
func AugmentResource(res *idl.Resource) {
	if len(res.Skills) == 0 {
		return
	}

	ctx := context.Background()
	ss := factory.GetListService()

	skillList, err := ss.GetList(ctx, "Skills")
	if err != nil {
		log.Error(err)
		return
	}

	for i, s := range res.Skills {
		li := getSkillById(skillList.Values, s.ID)
		res.Skills[i].Name = li.Name
	}

}

func getSkillById(items []list.ListItem, id string) *list.ListItem {
	for _, item := range items {
		if strings.EqualFold(item.Value, id) {
			return &item
		}
	}

	return nil
}
