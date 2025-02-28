package organization

import (
	"context"
	"csserver/internal/common"

	"github.com/cscoding21/csval/validate"
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

	return org, nil
}

// SaveOrgDefaults update the default values in an organization
func (s *OrganizationService) SaveOrgDefaults(ctx context.Context, defaults OrganizationDefaults) (common.UpdateResult[Organization], error) {
	val := validate.NewSuccessValidationResult()

	org, err := s.GetDefaultOrganization(ctx)
	if err != nil {
		return common.NewUpdateResult(&val, org), err
	}

	org.Defaults = defaults
	return s.UpdateOrganization(ctx, org)
}
