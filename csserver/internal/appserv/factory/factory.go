package factory

import (
	"csserver/internal/config"
	"csserver/internal/providers/surreal"
	"csserver/internal/services/iam/auth"
	"csserver/internal/services/iam/user"
	"csserver/internal/services/list"
	"fmt"

	"github.com/surrealdb/surrealdb.go"

	log "github.com/sirupsen/logrus"
)

// ---define singletons
var (
	_dbclient *surreal.DBClient
)

// GetContextHelper get context helper singleton
func GetContextHelper() *config.ContextHelper {
	ch := config.ContextHelper{}
	return &ch
}

// GetDBClient get dbclient singleton
func GetDBClient() *surreal.DBClient {
	if _dbclient != nil {
		return _dbclient
	}

	// Connect to SurrealDB
	host := fmt.Sprintf("ws://%s:%v/rpc", config.Config.Database.Host, config.Config.Database.Port)
	_dbclient, err := surrealdb.New(host)
	if err != nil {
		log.Error(err)
		return nil
	}

	// Sign in
	if _, err = _dbclient.Signin(map[string]string{
		"user": config.Config.Database.User,
		"pass": config.Config.Database.Password,
	}); err != nil {
		log.Error(err)
		return nil
	}

	// Select namespace and database
	if _, err = _dbclient.Use(
		config.Config.Database.Namespace,
		config.Config.Database.Database,
	); err != nil {
		log.Error(err)
		return nil
	}

	return surreal.NewDBClient(*_dbclient)
}

// GetUserService get user service singleton
func GetUserService() *user.UserService {
	surrealClient := GetDBClient()
	contextHelper := config.ContextHelper{}
	return user.NewUserService(*surrealClient, contextHelper)
}

// GetAuthService get user service singleton
func GetAuthService() *auth.AuthService {
	return auth.NewAuthService([]auth.AuthProvider{
		auth.NewCSAuthProvider(),
	})
}

// GetListService get list service singleton
func GetListService() *list.ListService {
	surrealClient := GetDBClient()
	return list.NewListService(*surrealClient, config.ContextHelper{})
}
