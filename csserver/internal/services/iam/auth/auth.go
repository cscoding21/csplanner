package auth

import (
	"context"
	"csserver/internal/providers/nats"
	"csserver/internal/services/iam/user"
	"fmt"

	"github.com/Nerzal/gocloak/v13"

	log "github.com/sirupsen/logrus"
)

type AuthCredentials struct {
	Token    string
	Username string
	Password string
}

type AuthResult struct {
	Success      bool
	Token        string
	RefreshToken string
	User         user.User
	Claims       interface{}
	Message      string
}

type AuthService struct {
	KCClient       *gocloak.GoCloak
	KCClientID     string
	KCClientSecret string
	KCRealm        string
	PubSub         nats.PubSubProvider
}

// NewAuthService return an AuthService instance with proper configuration
func NewAuthService(client *gocloak.GoCloak,
	PubSubProvider nats.PubSubProvider,
	clientID string,
	clientSecret string,
	realm string) *AuthService {
	return &AuthService{
		KCClient:       client,
		KCClientID:     clientID,
		KCClientSecret: clientSecret,
		KCRealm:        realm,
		PubSub:         PubSubProvider,
	}
}

// ValidateToken returns a valid result if a token is active
func (s *AuthService) ValidateToken(ctx context.Context, token string) (AuthResult, error) {
	rptResult, err := s.KCClient.RetrospectToken(ctx, token, s.KCClientID, s.KCClientSecret, s.KCRealm)
	if err != nil {
		return NewFailingAuthResult(token, err), err
	}

	decoded, claims, err := s.KCClient.DecodeAccessToken(ctx, token, s.KCRealm)
	if err != nil {
		return NewFailingAuthResult(token, err), err
	}

	log.Infof("%s - status: %v\n", token, rptResult)
	log.Infof("decoded: %v\n", decoded)
	log.Infof("claimss: %v\n", claims)

	if rptResult.Active != nil && *rptResult.Active {
		return NewSuccessAuthResult(token, ""), nil
	}

	err = fmt.Errorf("token not active")
	return NewFailingAuthResult(token, err), err
}

// RefreshToken issue an updated authToken using the passed in refresh token
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (AuthResult, error) {
	newAuthToken, err := s.KCClient.RefreshToken(ctx, refreshToken, s.KCClientID, s.KCClientSecret, s.KCRealm)
	if err != nil {
		return NewFailingAuthResult(refreshToken, err), err
	}

	return NewSuccessAuthResult(newAuthToken.AccessToken, newAuthToken.RefreshToken), nil
}

// Signout issue a signout command to the auth server which will invalidate the refresh token
func (s *AuthService) Signout(ctx context.Context, refreshToken string) error {
	return s.KCClient.Logout(ctx, s.KCClientID, s.KCClientSecret, s.KCRealm, refreshToken)
}

// Authenticate iterate over the provided auth providers and return the first valid AuthResult if successful
func (s *AuthService) Authenticate(ctx context.Context, creds AuthCredentials) (AuthResult, error) {
	log.Info(creds)
	log.Info(s.KCClientID)
	log.Info(s.KCClientSecret)
	log.Info(s.KCRealm)

	token, err := s.KCClient.Login(ctx,
		s.KCClientID,
		s.KCClientSecret,
		s.KCRealm,
		creds.Username,
		creds.Password)

	if err != nil {
		return NewFailingAuthResult("", err), err
	}

	return NewSuccessAuthResult(token.AccessToken, token.RefreshToken), nil
}

// NewSuccessAuthResult return a new AuthResult with a success flag set to true and the given token
func NewSuccessAuthResult(token string, refreshToken string) AuthResult {
	return AuthResult{
		Success:      true,
		Token:        token,
		RefreshToken: refreshToken,
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
