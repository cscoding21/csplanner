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

		log.Warnf("OrgMiddleware: %v", orgKey)
		ctx = context.WithValue(ctx, config.OrgUrlKey, orgKey)

		m, err := orgmap.GetSaaSOrg(ctx)
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		} else {
			log.Infof("SaaS Org: %v", *m.Info.Org)
			log.Infof("SaaS Org: %v", *m.Info.Licenses)
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
