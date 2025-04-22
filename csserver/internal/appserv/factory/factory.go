package factory

import (
	"context"
	"csserver/internal/appserv/orgmap"
	"csserver/internal/config"
	"csserver/internal/events"
	"csserver/internal/providers/postgres"
	"csserver/internal/services/activity"
	"csserver/internal/services/comment"
	"csserver/internal/services/iam/appuser"
	"csserver/internal/services/iam/auth"
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
	"github.com/jackc/pgx/v5/pgxpool"

	log "github.com/sirupsen/logrus"
)

// ---define singletons
var (
	lock          = &sync.Mutex{}
	_saasDBClient *pgxpool.Pool
)

// GetDBClient return a configured DB client as a singleton
func GetDBClient(ctx context.Context) *pgxpool.Pool {
	org, err := orgmap.GetSaaSOrg(ctx)
	if err != nil {
		log.Error(err)
	}

	return org.DB
}

// GetDBClient return a configured DB client as a singleton
func GetSaasDBClient() *pgxpool.Pool {
	if _saasDBClient != nil {
		return _saasDBClient
	}

	lock.Lock()
	defer lock.Unlock()

	log.Info("Creating Master DBClient instance now.")

	// Connect to SurrealDB
	//"postgres://username:password@localhost:5432/database_name"
	host := fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
		config.Config.MasterDB.User,
		config.Config.MasterDB.Password,
		config.Config.MasterDB.Host,
		config.Config.MasterDB.Port,
		config.Config.MasterDB.Database)

	// host := fmt.Sprintf("ws://%s:%v/rpc", config.Config.Database.User, config.Config.Database.Password)
	db, err := postgres.GetDB(context.Background(), host)
	if err != nil {
		log.Error(err)
		return nil
	}

	_saasDBClient = db

	return _saasDBClient
}

// GetPubSubClient return a pubsub client
func GetPubSubClient(ctx context.Context) (events.PubSubProvider, error) {
	ps := events.NewPubSubProvider(
		config.Config.PubSub.Host,
		config.Config.PubSub.Name,
		config.Config.PubSub.SubjectFormat,
		config.Config.PubSub.StreamName,
	)

	return ps, nil
}

// GetKeycloakClient return a Keycloak client
func GetKeycloakClient(ctx context.Context) *gocloak.GoCloak {
	client := gocloak.NewClient(config.Config.Security.KeycloakURL)
	return client
}

// GetActivityService get activity service instance
func GetActivityService(ctx context.Context) *activity.ActivityService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}
	return activity.NewActivityService(dbClient, pubsub)
}

// GetAuthService get user service instance
func GetAuthService(ctx context.Context) *auth.AuthService {
	kc := GetKeycloakClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
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
func GetCommentService(ctx context.Context) *comment.CommentService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return comment.NewCommentService(dbClient, pubsub)
}

// GetListService get list service instance
func GetListService(ctx context.Context) *list.ListService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return list.NewListService(dbClient, pubsub)
}

// GetNotificationService get notification service instance
func GetNotificationService(ctx context.Context) *notification.NotificationService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return notification.NewNotificationService(dbClient, pubsub)
}

// GetOrganizationService get organization service instance
func GetOrganizationService(ctx context.Context) *organization.OrganizationService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return organization.NewOrganizationService(dbClient, pubsub)
}

// GetProjectService return a project service instance
func GetProjectService(ctx context.Context) *project.ProjectService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return project.NewProjectService(dbClient, pubsub)
}

// GetPortfolioService return a portfolio service instance
func GetPortfolioService(ctx context.Context) *portfolio.PortfolioService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	projectServce := GetProjectService(ctx)
	resourceService := GetResourceService(ctx)
	schedService := GetScheduleService(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return portfolio.NewPortfolioService(
		dbClient,
		pubsub,
		*projectServce,
		*resourceService,
		*schedService)
}

// GetProcessorService return a processor service instance
func GetProcessorService(ctx context.Context) *processor.ProcessorService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	projectServce := GetProjectService(ctx)
	resourceService := GetResourceService(ctx)
	schedService := GetScheduleService(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return processor.NewProcessorService(
		dbClient,
		pubsub,
		*projectServce,
		*resourceService,
		*schedService)
}

// GetProjectTemplateService return a project template service instance
func GetProjectTemplateService(ctx context.Context) *projecttemplate.ProjecttemplateService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return projecttemplate.NewProjecttemplateService(dbClient, pubsub)
}

// GetResourceService return a resource service instance
func GetResourceService(ctx context.Context) *resource.ResourceService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return resource.NewResourceService(dbClient, pubsub)
}

// GetScheduleService return a schedule service instance
func GetScheduleService(ctx context.Context) *schedule.ScheduleService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return schedule.NewScheduleService(dbClient, pubsub)
}

// GetIAMAdminService get user service instance
func GetIAMAdminService(ctx context.Context) *auth.IAMAdminService {
	client := GetKeycloakClient(ctx)
	userService := GetAppuserService(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	svc := auth.NewIAMAdminService(
		client,
		pubsub,
		config.Config.Security.KeycloakRealm,
		config.Config.Security.KeycloakAdminUser,
		config.Config.Security.KeycloakAdminPass,
		*userService)

	return &svc
}

// GetAppuserService return a resource service instance
func GetAppuserService(ctx context.Context) *appuser.AppuserService {
	dbClient := GetDBClient(ctx)
	pubsub, err := GetPubSubClient(ctx)
	if err != nil {
		log.Error(err)
		return nil
	}

	return appuser.NewAppuserService(dbClient, pubsub)
}

// GetDefaultOrganization return the default organization for the tenant
func GetDefaultOrganization(ctx context.Context) (*organization.Organization, error) {
	service := GetOrganizationService(ctx)
	org, err := service.GetDefaultOrganization(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return org, nil
}
