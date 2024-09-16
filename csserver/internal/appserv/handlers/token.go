package handlers

import (
	"csserver/internal/appserv/factory"
	"csserver/internal/services/iam/auth"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type LoginHandler struct{}
type SignoutHandler struct{}
type RefreshHandler struct{}

type RefreshArgs struct {
	RefreshToken string `json:"refreshToken"`
}

// ServeHTTP handles login functionality
func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("LoginHandler")
	ctx := r.Context()
	as := factory.GetAuthService(ctx)
	var creds auth.AuthCredentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Error(err)
	}

	resp, err := as.Authenticate(ctx, creds)
	if err != nil {
		log.Error(err)
	}

	status := fmt.Sprintf(
		`{"success": %v, "accessToken": "%s", "refreshToken": "%s"}`,
		resp.Success,
		resp.Token,
		resp.RefreshToken)

	w.Header().Set("Content-Type", "application/json")

	if resp.Success {
		http.SetCookie(w, GetAuthCookie("accessToken", resp.Token))
		http.SetCookie(w, GetAuthCookie("refreshToken", resp.RefreshToken))

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.Write([]byte(status))
}

func GetLoginHandler() LoginHandler {
	return LoginHandler{}
}

// ServeHTTP handles logout functionality
func (h SignoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("SignoutHandler")
	ctx := r.Context()
	as := factory.GetAuthService(ctx)

	var args RefreshArgs
	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		log.Error(err)
	}

	status := `{"success": true }`
	err = as.Signout(ctx, args.RefreshToken)
	if err == nil {
		http.SetCookie(w, GetAuthCookie("accessToken", ""))
		http.SetCookie(w, GetAuthCookie("refreshToken", ""))

		w.WriteHeader(http.StatusOK)
	} else {
		log.Error(err)
		status = `{"success": false }`
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(status))
}

func GetSignoutHandler() SignoutHandler {
	return SignoutHandler{}
}

// ServeHTTP handles refresh functionality
func (h RefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("RefreshHandler")
	ctx := r.Context()
	as := factory.GetAuthService(ctx)
	var args RefreshArgs

	err := json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		log.Error(err)
	}

	resp, err := as.RefreshToken(ctx, args.RefreshToken)
	if err != nil {
		log.Error(err)
	}

	w.Header().Set("Content-Type", "application/json")

	status := fmt.Sprintf(
		`{"success": %v, "accessToken": "%s", "refreshToken": "%s"}`,
		resp.Success,
		resp.Token,
		resp.RefreshToken)

	if resp.Success {
		http.SetCookie(w, GetAuthCookie("accessToken", resp.Token))
		http.SetCookie(w, GetAuthCookie("refreshToken", resp.RefreshToken))

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.Write([]byte(status))
}

func GetRefreshHandler() RefreshHandler {
	return RefreshHandler{}
}

func GetAuthCookie(key string, value string) *http.Cookie {
	cookie := http.Cookie{
		Name:     key,
		Value:    value,
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
	}

	return &cookie
}
