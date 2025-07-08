package resource

import (
	"context"
	"csserver/internal/common"
	"csserver/internal/config"
	"csserver/internal/services/organization"
	"csserver/internal/services/resource/rtypes/resourcetype"
	"csserver/internal/utils"
	"slices"
)

var (
	resourceMap map[string]Resource
	roleMap     map[string]Role
)

// SaveResource upserts a resource and updates calculated fields
func (s *ResourceService) SaveResource(ctx context.Context, res Resource, org organization.Organization) (common.UpdateResult[*common.BaseModel[Resource]], error) {
	s.CalculateResourceInfo(ctx, &res, org)

	if len(res.ID) == 0 {
		return s.addNewResource(ctx, res)
	}

	wrapperR, err := s.GetResourceByID(ctx, res.ID)
	if err != nil {
		//---there's an ID, but no existing record.  this is a new resource
		return s.addNewResource(ctx, res)
	}

	if wrapperR == nil {
		return s.addNewResource(ctx, res)
	}

	return s.PatchResource(ctx, res, org)
}

func (s *ResourceService) addNewResource(ctx context.Context, res Resource) (common.UpdateResult[*common.BaseModel[Resource]], error) {
	if len(res.ID) == 0 {
		res.ID = utils.GenerateBase64UUID()
	}

	if res.Type == resourcetype.Human {
		if len(res.Skills) == 0 {
			roleMap, err := s.GetRoleMap(ctx, true)
			if err != nil {
				return common.NewFailingUpdateResult[*common.BaseModel[Resource]](nil, err)
			}

			role, ok := roleMap[*res.RoleID]
			if ok {
				res.Skills = append(res.Skills, role.DefaultSkills...)
			}
		}
	}

	r, err := s.CreateResource(ctx, res)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Resource]](nil, err)
	}

	s.pubsub.StreamPublish(ctx, config.GetOrgUrlKeyFromContext(ctx), string(ResourceIdentifier), "resource", "created", r)

	return r, nil
}

// PatchResource performs a surgical update of a resource, specially handing certain fields
func (s *ResourceService) PatchResource(ctx context.Context, resource Resource, org organization.Organization) (common.UpdateResult[*common.BaseModel[Resource]], error) {
	//val := validate.NewSuccessValidationResult()

	res, err := s.GetResourceByID(ctx, resource.ID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Resource]](nil, err)
	}

	res.Data.Name = resource.Name
	res.Data.ProfileImage = resource.ProfileImage
	res.Data.RoleID = resource.RoleID
	res.Data.AnnualizedCost = resource.AnnualizedCost
	res.Data.InitialCost = resource.InitialCost
	res.Data.Type = resource.Type
	res.Data.Status = resource.Status

	s.CalculateResourceInfo(ctx, &res.Data, org)

	out, err := s.UpdateResource(ctx, res.Data)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Resource]](nil, err)
	}

	s.pubsub.StreamPublish(ctx, config.GetOrgUrlKeyFromContext(ctx), string(ResourceIdentifier), "resource", "updated", out)

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
func (s *ResourceService) UpdateSkillForResource(ctx context.Context, id string, skill Skill) (common.UpdateResult[*common.BaseModel[Resource]], error) {
	//val := validate.NewSuccessValidationResult()

	res, err := s.GetResourceByID(ctx, id)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Resource]](nil, err)
	}

	if skillsContain(res.Data.Skills, skill.SkillID) {
		//---update the existing skill
		for i, s := range res.Data.Skills {
			if s.SkillID == skill.SkillID {
				res.Data.Skills[i].Proficiency = skill.Proficiency
			}
		}
	} else {
		res.Data.Skills = append(res.Data.Skills, &skill)
	}

	out, err := s.UpdateResource(ctx, res.Data)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Resource]](nil, err)
	}

	return out, nil
}

// RemoveSkillFromResource remove a skill from a resource
func (s *ResourceService) RemoveSkillFromResource(ctx context.Context, resourceID string, skillID string) (common.UpdateResult[*common.BaseModel[Resource]], error) {
	//val := validate.NewSuccessValidationResult()

	res, err := s.GetResourceByID(ctx, resourceID)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Resource]](nil, err)
	}

	if skillsContain(res.Data.Skills, skillID) {
		//---update the existing skill
		for i, s := range res.Data.Skills {
			if s.SkillID == skillID {
				res.Data.Skills = slices.Delete(res.Data.Skills, i, i+1)
				break
			}
		}
	}

	out, err := s.UpdateResource(ctx, res.Data)
	if err != nil {
		return common.NewFailingUpdateResult[*common.BaseModel[Resource]](nil, err)
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
		m[r.ID] = r.Data
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
		m[r.ID] = r.Data
	}

	roleMap = m

	return roleMap, nil
}

// skillsContain return true is a skill with the corresponging ID already exists
func skillsContain(skills []*Skill, skillID string) bool {
	if len(skills) == 0 {
		return false
	}

	for _, s := range skills {
		if s.SkillID == skillID {
			return true
		}
	}

	return false
}
