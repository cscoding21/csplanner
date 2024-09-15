package resource

import (
	"context"
	"csserver/internal/common"
	"slices"

	"github.com/cscoding21/csval/validate"
	"github.com/opentracing/opentracing-go/log"
)

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
