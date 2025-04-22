package middleware

import (
	"context"
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
