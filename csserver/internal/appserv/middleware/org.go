package middleware

import (
	"context"
	"csserver/internal/appserv/orgmap"
	"csserver/internal/config"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func OrgMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orgKey := getOrgUrlKey(r.Host)

		ctx := r.Context()

		log.Debugf("OrgMiddleware: %v", orgKey)
		ctx = context.WithValue(ctx, config.OrgUrlKey, orgKey)

		m, err := orgmap.GetSaaSOrg(ctx)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		} else {
			log.Debugf("SaaS Org: %v", *m.Info.Org)
			log.Debugf("SaaS Licenses: %v", *m.Info.Licenses)
		}

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func getOrgUrlKey(host string) string {
	if strings.Contains(host, ":") {
		parts := strings.Split(host, ":")

		host = parts[0]
	}
	hostParts := strings.Split(host, ".")

	return hostParts[0]
}
