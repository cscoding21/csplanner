package factory

import (
	"context"
	"csserver/internal/config"
	"csserver/internal/providers/contentful"
	"csserver/internal/providers/nats"
	"csserver/internal/providers/postgres"
	"csserver/internal/services/activity"
	"csserver/internal/services/comment"
	"csserver/internal/services/content"
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
	"github.com/jackc/pgx/v5"

	log "github.com/sirupsen/logrus"
)

// ---define singletons
var (
	lock      = &sync.Mutex{}
	_dbclient *pgx.Conn
)

// GetDBClient return a configured DB client as a singleton
func GetDBClient() *pgx.Conn {
	if _dbclient != nil {
		return _dbclient
	}

	lock.Lock()
	defer lock.Unlock()

	fmt.Println("Creating DBClient instance now.")

	// Connect to SurrealDB
	//"postgres://username:password@localhost:5432/database_name"
	host := fmt.Sprintf("postgres://%s:%s@%s:%v/%s",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Database)

	db, err := postgres.GetDB(context.Background(), host)
	if err != nil {
		log.Error(err)
		return nil
	}

	_dbclient = db

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

func GetContentfulProvider() (*contentful.ContentfulProvider, error) {
	cp := &contentful.ContentfulProvider{
		OrgID:   config.Config.CMS.OrgID,
		SpaceID: config.Config.CMS.SpaceID,
		PAT:     config.Config.CMS.PAT,
	}

	return cp, nil
}

// GetKeycloakClient return a Keycloak client
func GetKeycloakClient() *gocloak.GoCloak {
	client := gocloak.NewClient(config.Config.Security.KeycloakURL)
	return client
}

// GetActivityService get activity service instance
func GetActivityService() *activity.ActivityService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}
	return activity.NewActivityService(dbClient, pubsub)
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
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return comment.NewCommentService(dbClient, pubsub)
}

// GetListService get list service instance
func GetListService() *list.ListService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return list.NewListService(dbClient, pubsub)
}

// GetListService get list service instance
func GetContentService() *content.ContentService {
	cms, err := GetContentfulProvider()
	if err != nil {
		log.Error(err)
		return nil
	}

	return content.NewContentService(*cms)
}

// GetNotificationService get notification service instance
func GetNotificationService() *notification.NotificationService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return notification.NewNotificationService(dbClient, pubsub)
}

// GetOrganizationService get organization service instance
func GetOrganizationService() *organization.OrganizationService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return organization.NewOrganizationService(dbClient, pubsub)
}

// GetProjectService return a project service instance
func GetProjectService() *project.ProjectService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return project.NewProjectService(dbClient, pubsub)
}

// GetPortfolioService return a portfolio service instance
func GetPortfolioService() *portfolio.PortfolioService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	projectServce := GetProjectService()
	resourceService := GetResourceService()
	schedService := GetScheduleService()
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
func GetProcessorService() *processor.ProcessorService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	projectServce := GetProjectService()
	resourceService := GetResourceService()
	schedService := GetScheduleService()
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
func GetProjectTemplateService() *projecttemplate.ProjecttemplateService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return projecttemplate.NewProjecttemplateService(dbClient, pubsub)
}

// GetResourceService return a resource service instance
func GetResourceService() *resource.ResourceService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return resource.NewResourceService(dbClient, pubsub)
}

// GetScheduleService return a schedule service instance
func GetScheduleService() *schedule.ScheduleService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return schedule.NewScheduleService(dbClient, pubsub)
}

// GetIAMAdminService get user service instance
func GetIAMAdminService() *auth.IAMAdminService {
	client := GetKeycloakClient()
	userService := GetAppuserService()
	pubsub, err := GetPubSubClient()
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
func GetAppuserService() *appuser.AppuserService {
	dbClient := GetDBClient()
	pubsub, err := GetPubSubClient()
	if err != nil {
		log.Error(err)
		return nil
	}

	return appuser.NewAppuserService(dbClient, pubsub)
}

// GetDefaultOrganization return the default organization for the tenant
func GetDefaultOrganization(ctx context.Context) (*organization.Organization, error) {
	service := GetOrganizationService()
	org, err := service.GetOrganizationByID(ctx, "organization:default")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &org.Data, nil
}
