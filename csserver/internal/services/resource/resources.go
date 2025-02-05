package resource

import (
	"context"
	"csserver/internal/common"
	"slices"

	"github.com/cscoding21/csval/validate"

	log "github.com/sirupsen/logrus"
)

var (
	resourceMap map[string]Resource
	roleMap     map[string]Role
)

// PatchResource performs a surgical update of a resource, specially handing certain fields
func (s *ResourceService) PatchResource(ctx context.Context, resource Resource) (common.UpdateResult[Resource], error) {
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

	out, err := s.UpdateResource(ctx, res)
	if err != nil {
		log.Error(err)
		return common.NewUpdateResult(&val, res), err
	}

	return out, nil
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
