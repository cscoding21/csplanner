package setup

import (
	"context"
	"csserver/internal/appserv/orgmap"
	"csserver/internal/config"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func DeleteAllData(ctx context.Context) error {
	ctx = context.WithValue(ctx, config.OrgUrlKey, TEST_KEY)

	ctx = GetBotContext(ctx)
	org, err := orgmap.GetSaaSOrg(ctx)
	if err != nil {
		log.Fatalf("ERROR setting up test org: %s\n", err)
	} else {
		log.Infof("Deleting data for test org %s", org.Info.Org.Name)
	}

	tables := []string{"list", "organization", "project", "projecttemplate", "resource", "role", "appuser", "comment", "reaction"}

	for _, t := range tables {
		sql := fmt.Sprintf("DELETE FROM %s;", t)

		_, err = org.DB.Exec(ctx, sql)
		if err != nil {
			log.Errorf("ERROR deleting from %s: %s", t, err)
			err = errors.Join(err, err)
		} else {
			log.Infof("Deleted all from %s", t)
		}

	}

	return err
}
