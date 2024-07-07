package auth

import (
	"csserver/internal/config"
	"fmt"

	"github.com/golang-jwt/jwt/v5"

	log "github.com/sirupsen/logrus"
)

type CSAuthProvider struct{}

func NewCSAuthProvider() CSAuthProvider {
	return CSAuthProvider{}
}

func (p CSAuthProvider) Authenticate(tokenString string) AuthResult {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//---TODO: properly validate audience
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// }

		return []byte(config.Config.Security.JwtSigningKey), nil
	})
	if err != nil {
		log.Error(err)
		rv := NewFailingAuthResult(tokenString)

		rv.Message = err.Error()

		return rv
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && token.Valid {
		rv := NewFailingAuthResult(tokenString)

		rv.Message = "error processing claims from token"

		return rv
	}

	rv := NewSuccessAuthResult(tokenString)
	rv.Claims = claims
	rv.Message = "success"

	return rv
}
