package csmap

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/common"
	"csserver/internal/services/resource"
	"csserver/internal/utils"
)

// ---RESOURCE
func ConvertResourceResultToEnvelope(model *common.BaseModel[resource.Resource]) *idl.ResourceEnvelope {
	out := idl.ResourceEnvelope{}

	out.Meta = GetDataEnvelope(model)
	out.Data = utils.ValToRef(ResourceResourceToIdl(model.Data))

	AugmentResource(out.Data)

	return &out
}

func ConvertResourceResultToEnvelopeSlice(results []*common.BaseModel[resource.Resource]) []*idl.ResourceEnvelope {
	out := []*idl.ResourceEnvelope{}

	for _, r := range results {
		out = append(out, ConvertResourceResultToEnvelope(r))
	}

	return out
}

// ---ROLE
func ConvertRoleResultToEnvelope(model *common.BaseModel[resource.Role]) *idl.RoleEnvelope {
	out := idl.RoleEnvelope{}

	out.Meta = GetDataEnvelope(model)
	out.Data = utils.ValToRef(RoleResourceToIdl(model.Data))

	AugmentRole(out.Data)

	return &out
}

func ConvertRoleResultToEnvelopeSlice(results []*common.BaseModel[resource.Role]) []*idl.RoleEnvelope {
	out := []*idl.RoleEnvelope{}

	for _, r := range results {
		out = append(out, ConvertRoleResultToEnvelope(r))
	}

	return out
}
