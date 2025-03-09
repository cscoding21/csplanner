package organization

import (
	"context"
	"csserver/internal/common"
)

const (
	DefaultOrganizationID = "organization:default"
)

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
