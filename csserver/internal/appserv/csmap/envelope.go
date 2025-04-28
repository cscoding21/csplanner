package csmap

import (
	"context"
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/common"
	"csserver/internal/services/comment"
	"csserver/internal/services/project"
	"csserver/internal/services/resource"
	"csserver/internal/utils"
)

// ---RESOURCE
func ConvertResourceResultToEnvelope(ctx context.Context, model *common.BaseModel[resource.Resource]) *idl.ResourceEnvelope {
	out := idl.ResourceEnvelope{}

	out.Meta = GetDataEnvelope(model)
	out.Data = utils.ValToRef(ResourceResourceToIdl(model.Data))

	AugmentBaseModel(ctx, out.Meta)
	AugmentResource(ctx, out.Data)

	return &out
}

func ConvertResourceResultToEnvelopeSlice(ctx context.Context, results []*common.BaseModel[resource.Resource]) []*idl.ResourceEnvelope {
	out := []*idl.ResourceEnvelope{}

	for _, r := range results {
		out = append(out, ConvertResourceResultToEnvelope(ctx, r))
	}

	return out
}

// ---ROLE
func ConvertRoleResultToEnvelope(ctx context.Context, model *common.BaseModel[resource.Role]) *idl.RoleEnvelope {
	out := idl.RoleEnvelope{}

	out.Meta = GetDataEnvelope(model)
	out.Data = utils.ValToRef(RoleResourceToIdl(model.Data))

	AugmentBaseModel(ctx, out.Meta)
	AugmentRole(ctx, out.Data)

	return &out
}

func ConvertRoleResultToEnvelopeSlice(ctx context.Context, results []*common.BaseModel[resource.Role]) []*idl.RoleEnvelope {
	out := []*idl.RoleEnvelope{}

	for _, r := range results {
		out = append(out, ConvertRoleResultToEnvelope(ctx, r))
	}

	return out
}

// ---PROJECT
func ConvertProjectResultToEnvelope(ctx context.Context, model *common.BaseModel[project.Project]) *idl.ProjectEnvelope {
	out := idl.ProjectEnvelope{}

	out.Meta = GetDataEnvelope(model)
	out.Data = utils.ValToRef(ProjectProjectToIdl(model.Data))

	AugmentBaseModel(ctx, out.Meta)
	AugmentProject(ctx, &model.Data, out.Data)

	return &out
}

func ConvertProjectResultToEnvelopeSlice(ctx context.Context, results []*common.BaseModel[project.Project]) []*idl.ProjectEnvelope {
	out := []*idl.ProjectEnvelope{}

	for _, r := range results {
		out = append(out, ConvertProjectResultToEnvelope(ctx, r))
	}

	return out
}

// ---COMMENT
func ConvertCommentResultToEnvelope(ctx context.Context, model *common.BaseModel[comment.Comment]) *idl.CommentEnvelope {
	out := idl.CommentEnvelope{}

	out.Meta = GetDataEnvelope(model)
	out.Data = utils.ValToRef(CommentCommentToIdl(model.Data))
	out.Data.Replies = ConvertCommentResultToEnvelopeSlice(ctx, utils.ValToRefSlice(model.Data.CommentReplies))

	AugmentBaseModel(ctx, out.Meta)

	return &out
}

func ConvertCommentResultToEnvelopeSlice(ctx context.Context, results []*common.BaseModel[comment.Comment]) []*idl.CommentEnvelope {
	out := []*idl.CommentEnvelope{}

	for _, r := range results {
		out = append(out, ConvertCommentResultToEnvelope(ctx, r))
	}

	return out
}
