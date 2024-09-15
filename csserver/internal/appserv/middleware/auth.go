package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler/transport"

	"csserver/internal/appserv/factory"
	"csserver/internal/config"
	"csserver/internal/services/iam/auth"

	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type AuthResult struct {
	Valid bool
	Token interface{}
}

type GraphQLShape struct {
	Query         string
	OperationName string
	Variables     interface{}
}

// ValidationMiddleware
func ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//---temporary for testing
		if checkIsFile(r) {
			log.Info("Auth bypass: is file")
			next.ServeHTTP(w, r)
			return
		}

		if swallowHealthCheck(r) || checkIsPlayground(r) {
			log.Info("Auth bypass: playground or swallow")
			next.ServeHTTP(w, r)
			return
		}

		if checkWebsocketKey(r) {
			log.Info("Auth bypass: web socket key")
			next.ServeHTTP(w, r)
			return
		}

		us := factory.GetUserService()
		bctx := r.Context()

		if allowAnonomousOperation(r) || bypassSecurity(r) {
			anonEmail := config.Config.Default.BotUserEmail
			log.Warnf("AI Bot credentials set for %s", anonEmail)

			anonID, err := us.GetUser(bctx, anonEmail)
			if err != nil {
				log.Error(err)
			} else {
				log.Warnf("resolved user %s %s", anonID.FirstName, anonID.LastName)
			}

			bctx = context.WithValue(bctx, config.UserEmailKey, anonEmail)
			bctx = context.WithValue(bctx, config.UserIDKey, anonID)
			r = r.WithContext(bctx)

			next.ServeHTTP(w, r)
			return
		}

		token := getTokenFromHeader(r)
		log.Infof("Auth token: %s", token)

		authService := factory.GetAuthService(bctx)
		result, err := authService.ValidateToken(bctx, token)
		if err != nil {
			log.Errorf("error on validate: %s", err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		log.Debugf("AUTH Validation Result for token %s: %v", token, result)
		if !result.Success {
			err := fmt.Errorf("login failed for token %s", token)
			log.Error(err)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if result.Success {
			log.Debugf("Authentication success for user %s", result.User.Email)

			ctx := context.WithValue(r.Context(), config.UserEmailKey, result.User.Email)
			ctx = context.WithValue(ctx, config.UserIDKey, result.User.ID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
			return
		}

		log.Warn("Authentication failed")
		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}

// AuthMiddleware sets authorization cookies
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// Pass control back to the handler
		next.ServeHTTP(w, r)
	})
}

func WebSocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
	// Get the token from payload
	log.Infof("INIT PAYLOAD: %v", initPayload)
	any := initPayload["Authorization"]

	token, ok := any.(string)
	if !ok || token == "" {
		log.Errorf("error in init.  will attempt auth token - Authorization: %s", token)
		//---try authToken.  these settings are handled by the providers
		//TODO: search for a way to handle properly
		any = initPayload["authToken"]
		token, ok = any.(string)
		if !ok || token == "" {
			log.Errorf("authToken failed: %s", token)
			return nil, &initPayload, fmt.Errorf("authorization not found in transport payload: %v", initPayload)
		}
	}

	bctx := config.NewContext()
	authService := factory.GetAuthService(bctx)

	result, _ := authService.Authenticate(bctx, auth.AuthCredentials{Token: token})
	if result.Success {
		log.Debugf("WebSocket authentication success for user %s", result.User.Email)

		bctx = context.WithValue(bctx, config.UserEmailKey, result.User.Email)
		bctx = context.WithValue(bctx, config.UserIDKey, result.User.ID)

		return bctx, &initPayload, nil
	}

	return ctx, &initPayload, fmt.Errorf("login failed for token %s", token)
}

// swallowHealthCheck determine if the request is a kubernetes health check
func swallowHealthCheck(r *http.Request) bool {
	userAgent := r.Header["User-Agent"]

	if len(userAgent) > 0 && strings.Contains(userAgent[0], "kube-probe") {
		return true
	}

	return false
}

// allowAnonomousOperation audit the requested operation to see if if can be handled without authentication
func allowAnonomousOperation(r *http.Request) bool {
	var graphqlQuery GraphQLShape
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Error reading body from request:: %s", err)
		return false
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	json.Unmarshal([]byte(body), &graphqlQuery)
	log.Debug(graphqlQuery)

	allowed := []string{"IntrospectionQuery", "login", "createUser", "currentUser"}

	//---this will handle calls from the JS client
	if slices.Contains(allowed, graphqlQuery.OperationName) {
		log.Infof("Operation allowed via operation match %s", graphqlQuery.OperationName)
		return true
	}

	//---TODO: need a better way to confirm requests from playground.  This could be exploited
	for _, a := range allowed {
		if strings.Contains(graphqlQuery.Query, a) {
			log.Infof("Operation allowed via string search %s", a)
			return true
		}
	}

	return false
}

// bypassSecurity check to see if security should be bypassed.  For local development
func bypassSecurity(r *http.Request) bool {
	_, isPlayground := strings.CutPrefix(r.Referer(), "http://localhost:5000/playground")

	return config.Config.Security.BypassAuth || isPlayground
}

// getTokenFromHeader gets the jwt from the request header and strips the bearer prefix if necessary
func getTokenFromHeader(r *http.Request) string {
	token := r.Header.Get("Authorization")
	strippedToken := strings.Replace(token, "Bearer ", "", 1)

	return strippedToken
}

func getAccessTokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("accessToken")
	if err != nil {
		return ""
	}

	return cookie.Value
}

func getRefreshTokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		return ""
	}

	return cookie.Value
}

// checkIsPlayground return true if the request is to the playground
func checkIsPlayground(r *http.Request) bool {
	return r.URL.String() == "/playground"
}

// checkIsFile return true if the request is to the playground
func checkIsFile(r *http.Request) bool {
	return strings.Contains(r.URL.String(), "files")
}

// checkWebsocketKey handle web socket auth
func checkWebsocketKey(r *http.Request) bool {
	token := r.Header.Get("Sec-Websocket-Key")
	log.Debugf("WEBSOCKET KEY: %s", token)

	return len(token) > 0
}
