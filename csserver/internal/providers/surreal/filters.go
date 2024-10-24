package surreal

import (
	log "github.com/sirupsen/logrus"
)

func GetCustomFilter(name string) string {
	customFilters := loadFilters()
	out, ok := customFilters[name]
	if !ok {
		log.Warnf("cannot find custom filter key %s", name)
		return ""
	}

	log.Warn("out", out)

	return out
}

func loadFilters() map[string]string {
	f := make(map[string]string)

	f["resource_in_project"] = `
	AND (
        daci.driver_ids ANYINSIDE [$resourceID]
        OR daci.approver_ids ANYINSIDE [$resourceID]
        OR daci.contributor_ids ANYINSIDE [$resourceID]
        OR daci.informed_ids ANYINSIDE [$resourceID]
    )
	`

	return f
}
