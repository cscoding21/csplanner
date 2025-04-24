package auth

import (
	"context"
	"csserver/internal/events"
	"csserver/internal/services/iam/appuser"
	"fmt"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/golang-jwt/jwt/v5"

	log "github.com/sirupsen/logrus"
)

type AuthCredentials struct {
	Token    string
	Username string
	Password string
	Realm    string
}

type AuthResult struct {
	Success      bool
	Token        string
	Realm        string
	RefreshToken string
	User         appuser.Appuser
	Claims       interface{}
	Message      string
}

type AuthService struct {
	KCClient       *gocloak.GoCloak
	KCClientID     string
	KCClientSecret string
	PubSub         events.PubSubProvider
}

// NewAuthService return an AuthService instance with proper configuration
func NewAuthService(client *gocloak.GoCloak,
	PubSubProvider events.PubSubProvider,
	clientID string,
	clientSecret string) *AuthService {
	return &AuthService{
		KCClient:       client,
		KCClientID:     clientID,
		KCClientSecret: clientSecret,
		PubSub:         PubSubProvider,
	}
}

// ValidateToken returns a valid result if a token is active
func (s *AuthService) ValidateToken(ctx context.Context, realm string, token string) (AuthResult, error) {
	rptResult, err := s.KCClient.RetrospectToken(ctx, token, s.KCClientID, s.KCClientSecret, realm)
	if err != nil {
		return NewFailingAuthResult(token, err), err
	}

	if rptResult.Active != nil && *rptResult.Active {
		ret := s.NewSuccessAuthResult(ctx, realm, token, "")

		return ret, nil
	}

	err = fmt.Errorf("token not active")
	return NewFailingAuthResult(token, err), err
}

// RefreshToken issue an updated authToken using the passed in refresh token
func (s *AuthService) RefreshToken(ctx context.Context, realm string, refreshToken string) (AuthResult, error) {
	newAuthToken, err := s.KCClient.RefreshToken(ctx, refreshToken, s.KCClientID, s.KCClientSecret, realm)
	if err != nil {
		return NewFailingAuthResult(refreshToken, err), err
	}

	return s.NewSuccessAuthResult(ctx, realm, newAuthToken.AccessToken, newAuthToken.RefreshToken), nil
}

// Signout issue a signout command to the auth server which will invalidate the refresh token
func (s *AuthService) Signout(ctx context.Context, realm string, refreshToken string) error {
	return s.KCClient.Logout(ctx, s.KCClientID, s.KCClientSecret, realm, refreshToken)
}

// Authenticate iterate over the provided auth providers and return the first valid AuthResult if successful
func (s *AuthService) Authenticate(ctx context.Context, creds AuthCredentials) (AuthResult, error) {
	log.Debugf("CREDS: %v", creds)
	log.Debugf("ClientID: %v", s.KCClientID)
	log.Debugf("Client Secret: %v", s.KCClientSecret)
	log.Debugf("Realm: %v", creds.Realm)

	token, err := s.KCClient.Login(ctx,
		s.KCClientID,
		s.KCClientSecret,
		creds.Realm,
		creds.Username,
		creds.Password)

	if err != nil {
		return NewFailingAuthResult("", err), err
	}

	return s.NewSuccessAuthResult(ctx, creds.Realm, token.AccessToken, token.RefreshToken), nil
}

// NewSuccessAuthResult return a new AuthResult with a success flag set to true and the given token
func (s *AuthService) NewSuccessAuthResult(ctx context.Context, realm string, token string, refreshToken string) AuthResult {
	_, claims, err := s.KCClient.DecodeAccessToken(ctx, token, realm)
	if err != nil {

	}

	return AuthResult{
		Success:      true,
		Token:        token,
		Realm:        realm,
		RefreshToken: refreshToken,
		User: appuser.Appuser{
			Email:        GetKeyFromClaims(*claims, "email").(string),
			FirstName:    GetKeyFromClaims(*claims, "given_name").(string),
			LastName:     GetKeyFromClaims(*claims, "family_name").(string),
			DBID:         GetKeyFromClaims(*claims, "dbid").(string),
			ProfileImage: GetKeyFromClaims(*claims, "profileImage").(string),
		},
	}
}

// NewFailingAuthResult return a new AuthResult with a success flag set to false and the given token
func NewFailingAuthResult(token string, err error) AuthResult {
	return AuthResult{
		Success: false,
		Token:   token,
		Message: err.Error(),
	}
}

// GetKeyFromClaims iterate over the claims and return the value of the key if found
func GetKeyFromClaims(claims jwt.MapClaims, key string) interface{} {
	for k, v := range claims {
		if strings.EqualFold(k, key) {
			return v
		}
	}

	return ""
}
