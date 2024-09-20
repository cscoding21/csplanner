package augment

import (
	"context"
	"csserver/internal/appserv/factory"
	"csserver/internal/services/list"
	"strings"

	log "github.com/sirupsen/logrus"
)

func getSkills() *[]list.ListItem {
	ctx := context.Background()
	ss := factory.GetListService()

	skillList, err := ss.GetList(ctx, "Skills")
	if err != nil {
		log.Error(err)
		return nil
	}

	return &skillList.Values
}

func getSkillById(items []list.ListItem, id string) *list.ListItem {
	for _, item := range items {
		if strings.EqualFold(item.Value, id) {
			return &item
		}
	}

	return nil
}
