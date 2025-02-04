package factory

import (
	"context"
	"csserver/internal/config"
	"csserver/internal/providers/nats"
	"csserver/internal/providers/surreal"
	"csserver/internal/services/activity"
	"csserver/internal/services/comment"
	"csserver/internal/services/iam/auth"
	"csserver/internal/services/iam/user"
	"csserver/internal/services/list"
	"csserver/internal/services/notification"
	"csserver/internal/services/organization"
	"csserver/internal/services/portfolio"
	"csserver/internal/services/processor"
	"csserver/internal/services/project"
	"csserver/internal/services/projecttemplate"
	"csserver/internal/services/resource"
	"csserver/internal/services/schedule"
	"fmt"
	"sync"

	"github.com/Nerzal/gocloak/v13"
	"github.com/surrealdb/surrealdb.go"

	log "github.com/sirupsen/logrus"
)

// ---define singletons
var (
	lock      = &sync.Mutex{}
	_dbclient *surreal.DBClient
)

// GetContextHelper get context helper
func GetContextHelper() *config.ContextHelper {
	ch := config.ContextHelper{}
	return &ch
}

// GetDBClient return a configured DB client as a singleton
func GetDBClient() *surreal.DBClient {
	if _dbclient != nil {
		return _dbclient
	}

	lock.Lock()
	defer lock.Unlock()

	fmt.Println("Creating DBClient instance now.")

	// Connect to SurrealDB
	host := fmt.Sprintf("ws://%s:%v/rpc", config.Config.Database.Host, config.Config.Database.Port)
	_db, err := surrealdb.New(host)
	if err != nil {
		log.Error(err)
		return nil
	}

	// Sign in
	if _, err = _db.Signin(map[string]string{
		"user": config.Config.Database.User,
		"pass": config.Config.Database.Password,
	}); err != nil {
		log.Error(err)
		return nil
	}

	// Select namespace and database
	if _, err = _db.Use(
		config.Config.Database.Namespace,
		config.Config.Database.Database,
	); err != nil {
		log.Error(err)
		return nil
	}

	_dbclient = surreal.NewDBClient(*_db)

	return _dbclient
}

// GetPubSubClient return a pubsub client
func GetPubSubClient() (nats.PubSubProvider, error) {
	ps := nats.NewPubSubProvider(
		config.Config.PubSub.Host,
		config.Config.PubSub.Name,
		config.Config.PubSub.SubjectFormat,
		config.Config.PubSub.StreamName,
	)

	return ps, nil
}

// GetKeycloakClient return a Keycloak client
func GetKeycloakClient() *gocloak.GoCloak {
	client := gocloak.NewClient(config.Config.Security.KeycloakURL)
	return client
}

// GetActivityService get activity service instance
func GetActivityService() *activity.ActivityService {
	surrealClient := GetDBClient()
	contextHelper := config.ContextHelper{}
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}
	return activity.NewActivityService(*surrealClient, contextHelper, pubsub)
}

// GetAuthService get user service instance
func GetAuthService(ctx context.Context) *auth.AuthService {
	kc := GetKeycloakClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return auth.NewAuthService(
		kc,
		pubsub,
		config.Config.Security.KeycloakClientID,
		config.Config.Security.KeycloakClientSecret,
		config.Config.Security.KeycloakRealm)
}

// GetCommentService get comment service instance
func GetCommentService() *comment.CommentService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return comment.NewCommentService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetListService get list service instance
func GetListService() *list.ListService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return list.NewListService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetNotificationService get notification service instance
func GetNotificationService() *notification.NotificationService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return notification.NewNotificationService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetOrganizationService get organization service instance
func GetOrganizationService() *organization.OrganizationService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return organization.NewOrganizationService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetProjectService return a project service instance
func GetProjectService() *project.ProjectService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return project.NewProjectService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetPortfolioService return a portfolio service instance
func GetPortfolioService() *portfolio.PortfolioService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	projectServce := GetProjectService()
	resourceService := GetResourceService()
	schedService := GetScheduleService()
	if err != nil {
		log.Error(err)
		return nil
	}

	return portfolio.NewPortfolioService(
		*surrealClient,
		config.ContextHelper{},
		pubsub,
		*projectServce,
		*resourceService,
		*schedService)
}

// GetProcessorService return a processor service instance
func GetProcessorService() *processor.ProcessorService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	projectServce := GetProjectService()
	resourceService := GetResourceService()
	schedService := GetScheduleService()
	if err != nil {
		log.Error(err)
		return nil
	}

	return processor.NewProcessorService(
		*surrealClient,
		config.ContextHelper{},
		pubsub,
		*projectServce,
		*resourceService,
		*schedService)
}

// GetProjectTemplateService return a project template service instance
func GetProjectTemplateService() *projecttemplate.ProjecttemplateService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return projecttemplate.NewProjecttemplateService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetResourceService return a resource service instance
func GetResourceService() *resource.ResourceService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return resource.NewResourceService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetScheduleService return a schedule service instance
func GetScheduleService() *schedule.ScheduleService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return schedule.NewScheduleService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetIAMAdminService get user service instance
func GetIAMAdminService() *auth.IAMAdminService {
	contextHelper := config.ContextHelper{}
	client := GetKeycloakClient()
	userService := GetUserService()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	svc := auth.NewIAMAdminService(
		contextHelper,
		client,
		pubsub,
		config.Config.Security.KeycloakRealm,
		config.Config.Security.KeycloakAdminUser,
		config.Config.Security.KeycloakAdminPass,
		*userService)

	return &svc
}

// GetUserService return a resource service instance
func GetUserService() *user.UserService {
	surrealClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return user.NewUserService(*surrealClient, config.ContextHelper{}, pubsub)
}

// GetDefaultOrganization return the default organization for the tenant
func GetDefaultOrganization(ctx context.Context) (*organization.Organization, error) {
	service := GetOrganizationService()
	org, err := service.GetOrganizationByID(ctx, "organization:default")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return org, nil
}
