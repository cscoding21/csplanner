package organization

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/events"
	"csserver/internal/providers/postgres"
	"fmt"

	"github.com/cscoding21/csval/validate"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DefaultOrganizationID  = "organization:default"
	OrganizationIdentifier = postgres.TableName("organization")

	saaSOrgSQL = `
		SELECT 	id, 
				name, 
				db_host, 
				database, 
				url_key, 
				realm, 
				is_provisioned,
				created_at, 
				created_by, 
				updated_at, 
				updated_by 
			FROM organization 
			WHERE true
				AND deleted_at IS NULL
				AND url_key = $1
			;
		`
	licenseOrgSQL = `
		SELECT 	id,
				org_id, 
				app, 
				email, 
				expires_on, 
				cancelled_on, 
				auto_renew, 
				created_at, 
				created_by, 
				updated_at, 
				updated_by
			FROM license
			WHERE true
				AND deleted_at IS NULL
				AND org_id = $1
	`

	localOrgByKeySQL = `
		SELECT 	id, 
				parent_id, 
				created_at, 
				created_by, 
				updated_at, 
				updated_by, 
				deleted_at, 
				deleted_by, 
				data
		FROM %s
		WHERE true 
			AND deleted_at is null
			AND data->>'url' = $1;`
)

// OrganizationService is a service for interacting with lists.
type OrganizationService struct {
	db     *pgxpool.Pool
	pubsub events.PubSubProvider
}

// NewOrganizationService creates a new Organization service.
func NewOrganizationService(
	db *pgxpool.Pool,
	ps events.PubSubProvider) *OrganizationService {

	return &OrganizationService{
		db:     db,
		pubsub: ps,
	}
}

// GetSaaSInfo retrieves the saas information for an organization
func GetSaaSInfo(ctx context.Context, db *pgxpool.Pool) (*SaaSInfo, error) {
	urlKey := config.GetOrgUrlKeyFromContext(ctx)

	so, err := getSaaSOrg(ctx, db, urlKey)
	if err != nil {
		return nil, err
	}

	licenses, err := getOrgLicenses(ctx, db, so.ID)
	if err != nil {
		return nil, err
	}

	out := SaaSInfo{
		Org:      so,
		Licenses: licenses,
	}

	return &out, nil

	/*
		org, err := getLocalOrgByUrlKey(ctx, s.db, urlKey)
		if err != nil {
			return nil, err
		}

		licenses, err := getOrgLicenses(ctx, s.saasDB, saasOrg.ID)
		if err != nil {
			return nil, err
		}

		org.Data.Licenses = *licenses

		return org, nil
	*/
}

func getSaaSOrg(ctx context.Context, db *pgxpool.Pool, key string) (*SaaSOrg, error) {
	var org SaaSOrg
	rows, err := db.Query(ctx, saaSOrgSQL, key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	output, err := pgx.CollectRows(rows, pgx.RowToStructByName[SaaSOrg])
	if err != nil {
		return nil, err
	}

	if len(output) > 0 {
		org = output[0]
	} else {
		return &org, fmt.Errorf("organization %s not found", key)
	}

	return &org, nil
}

func getOrgLicenses(ctx context.Context, db *pgxpool.Pool, orgID string) (*[]License, error) {
	rows, err := db.Query(ctx, licenseOrgSQL, orgID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	licenses, err := pgx.CollectRows(rows, pgx.RowToStructByName[License])
	if err != nil {
		return nil, err
	}

	return &licenses, nil
}

func getLocalOrgByUrlKey(ctx context.Context, db *pgxpool.Pool, key string) (*common.BaseModel[Organization], error) {
	sql := fmt.Sprintf(localOrgByKeySQL, OrganizationIdentifier)

	return postgres.GetObject[Organization](ctx, db, OrganizationIdentifier, sql, key)
}

// GetOrganizationByID gets a Organization by its ID.
func (s *OrganizationService) GetOrganizationByID(ctx context.Context, id string) (*common.BaseModel[Organization], error) {
	return postgres.GetObjectByID[Organization](ctx, s.db, OrganizationIdentifier, id)
}

// GetDefaultOrganization return the default organization
func (s *OrganizationService) GetDefaultOrganization(
	ctx context.Context) (*Organization, error) {
	org, err := s.GetOrganizationByID(ctx, DefaultOrganizationID)
	if err != nil {
		return nil, err
	}

	return &org.Data, nil
}

// SaveOrgDefaults update the default values in an organization
func (s *OrganizationService) SaveOrgDefaults(ctx context.Context, defaults OrganizationDefaults) (common.UpdateResult[*common.BaseModel[Organization]], error) {
	//val := validate.NewSuccessValidationResult()

	org, err := s.GetDefaultOrganization(ctx)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Organization]](nil, err)
	}

	org.Defaults = defaults
	return s.UpdateOrganization(ctx, *org)
}

// UpdateOrganization update an existing Organization.
func (s *OrganizationService) UpdateOrganization(ctx context.Context, input Organization) (common.UpdateResult[*common.BaseModel[Organization]], error) {

	val := validate.NewSuccessValidationResult()

	obj, err := postgres.UpdateObject(ctx, s.db, input, OrganizationIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

// CreateOrganization creates a new Organization.
func (s *OrganizationService) CreateOrganization(ctx context.Context, input Organization) (common.UpdateResult[*common.BaseModel[Organization]], error) {

	val := validate.NewSuccessValidationResult()

	obj, err := postgres.UpdateObject(ctx, s.db, input, OrganizationIdentifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}
