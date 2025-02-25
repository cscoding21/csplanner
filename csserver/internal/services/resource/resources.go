package resource

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/services/organization"
	"csserver/internal/services/resource/rtypes/resourcetype"
	"slices"

	"github.com/cscoding21/csval/validate"

	log "github.com/sirupsen/logrus"
)

var (
	resourceMap map[string]Resource
	roleMap     map[string]Role
)

// SaveResource upserts a resource and updates calculated fields
func (s *ResourceService) SaveResource(ctx context.Context, res Resource, org organization.Organization) (common.UpdateResult[Resource], error) {
	s.CalculateResourceInfo(ctx, &res, org)

	if len(res.ID) == 0 {
		return s.addNewResource(ctx, res)
	}

	_, err := s.GetResourceByID(ctx, res.ID)
	if err != nil {
		//---there's an ID, but no existing record.  this is a new resource
		return s.addNewResource(ctx, res)
	}

	return s.PatchResource(ctx, res, org)
}

func (s *ResourceService) addNewResource(ctx context.Context, res Resource) (common.UpdateResult[Resource], error) {
	if res.Type == resourcetype.Human {
		if len(res.Skills) == 0 {
			roleMap, err := s.GetRoleMap(ctx, true)
			if err != nil {
				return common.NewFailingUpdateResult(&res, err)
			}

			role, ok := roleMap[*res.RoleID]
			if ok {
				res.Skills = append(res.Skills, role.DefaultSkills...)
			}
		}
	}

	return s.CreateResource(ctx, &res)
}

// PatchResource performs a surgical update of a resource, specially handing certain fields
func (s *ResourceService) PatchResource(ctx context.Context, resource Resource, org organization.Organization) (common.UpdateResult[Resource], error) {
	val := validate.NewSuccessValidationResult()

	res, err := s.GetResourceByID(ctx, resource.ID)
	if err != nil {
		log.Error(err)
		return common.NewUpdateResult(&val, res), err
	}

	res.Name = resource.Name
	res.ProfileImage = resource.ProfileImage
	res.RoleID = resource.RoleID
	res.AnnualizedCost = resource.AnnualizedCost
	res.InitialCost = resource.InitialCost
	res.Type = resource.Type
	res.Status = resource.Status

	s.CalculateResourceInfo(ctx, res, org)

	out, err := s.UpdateResource(ctx, res)
	if err != nil {
		log.Error(err)
		return common.NewUpdateResult(&val, res), err
	}

	return out, nil
}

func (s *ResourceService) CalculateResourceInfo(ctx context.Context, res *Resource, org organization.Organization) {
	rm, err := s.GetRoleMap(ctx, true)
	if err != nil {
		return
	}

	res.Calculated.HourlyCost, res.Calculated.HourlyCostMethod = res.GetHourlyRate(rm, float64(org.Defaults.GenericBlendedHourlyRate), org.Defaults.WorkingHoursPerYear)
}

// UpdateSkillForResource add or update a skill to s resource
func (s *ResourceService) UpdateSkillForResource(ctx context.Context, id string, skill Skill) (common.UpdateResult[Resource], error) {
	val := validate.NewSuccessValidationResult()

	res, err := s.GetResourceByID(ctx, id)
	if err != nil {
		log.Error(err)
		return common.NewUpdateResult(&val, res), err
	}

	if skillsContain(res.Skills, skill.ID) {
		//---update the existing skill
		for i, s := range res.Skills {
			if s.ID == skill.ID {
				res.Skills[i].Proficiency = skill.Proficiency
			}
		}
	} else {
		res.Skills = append(res.Skills, &skill)
	}

	out, err := s.UpdateResource(ctx, res)
	if err != nil {
		log.Error(err)
		return common.NewUpdateResult(&val, res), err
	}

	return out, nil
}

// RemoveSkillFromResource remove a skill from a resource
func (s *ResourceService) RemoveSkillFromResource(ctx context.Context, resourceID string, skillID string) (common.UpdateResult[Resource], error) {
	val := validate.NewSuccessValidationResult()

	res, err := s.GetResourceByID(ctx, resourceID)
	if err != nil {
		log.Error(err)
		return common.NewUpdateResult(&val, res), err
	}

	if skillsContain(res.Skills, skillID) {
		//---update the existing skill
		for i, s := range res.Skills {
			if s.ID == skillID {
				res.Skills = slices.Delete(res.Skills, i, i+1)
			}
		}
	}

	out, err := s.UpdateResource(ctx, res)
	if err != nil {
		log.Error(err)
		return common.NewUpdateResult(&val, res), err
	}

	return out, nil
}

// GetResourceMap return a map of all resources keyed by ID
func (s *ResourceService) GetResourceMap(ctx context.Context, force bool) (map[string]Resource, error) {
	if resourceMap != nil && force {
		return resourceMap, nil
	}

	res, err := s.FindAllResources(ctx)
	if err != nil {
		return nil, err
	}

	m := make(map[string]Resource)
	for _, r := range res.Results {
		m[r.ID] = r
	}

	resourceMap = m

	return resourceMap, nil
}

// GetRoleMap return a map of all roles keyed by ID
func (s *ResourceService) GetRoleMap(ctx context.Context, force bool) (map[string]Role, error) {
	if resourceMap != nil && force {
		return roleMap, nil
	}

	res, err := s.FindAllRoles(ctx)
	if err != nil {
		return nil, err
	}

	m := make(map[string]Role)
	for _, r := range res.Results {
		m[r.ID] = r
	}

	roleMap = m

	return roleMap, nil
}

// skillsContain return true is a skill with the corresponging ID already exists
func skillsContain(skills []*Skill, id string) bool {
	if len(skills) == 0 {
		return false
	}

	for _, s := range skills {
		if s.ID == id {
			return true
		}
	}

	return false
}
