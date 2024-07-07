package auth

import "csserver/internal/services/iam/user"

type AuthCredentials struct {
	Token string
}

type AuthResult struct {
	Success bool
	Token   string
	User    user.User
	Claims  interface{}
	Message string
}

type AuthProvider interface {
	Authenticate(token string) AuthResult
}

type AuthService struct {
	Providers []AuthProvider
}

func NewAuthService(providers []AuthProvider) *AuthService {
	return &AuthService{
		Providers: providers,
	}
}

// Authenticate iterate over the provided auth providers and return the first valid AuthResult if successful
func (s AuthService) Authenticate(creds AuthCredentials) (AuthResult, error) {
	for _, provider := range s.Providers {
		result := provider.Authenticate(creds.Token)
		if result.Success {
			return result, nil
		}
	}

	return NewFailingAuthResult(creds.Token), nil
}

// NewSuccessAuthResult return a new AuthResult with a success flag set to true and the given token
func NewSuccessAuthResult(token string) AuthResult {
	return AuthResult{
		Success: true,
		Token:   token,
	}
}

// NewFailingAuthResult return a new AuthResult with a success flag set to false and the given token
func NewFailingAuthResult(token string) AuthResult {
	return AuthResult{
		Success: false,
		Token:   token,
	}
}
