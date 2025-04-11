package provision

import (
	"csserver/internal/common"
	"csserver/internal/services/list"
)

var skillsList = list.List{
	ControlFields: common.ControlFields{
		ID: "list:skills",
	},
	Name:        list.ListNameSkills,
	Description: "Skills are assigned to a resource.  This assignment allows for accurate project task assignment.",
	Values:      []*list.ListItem{},
}

var fundingSourceList = list.List{
	ControlFields: common.ControlFields{
		ID: "list:funding-source",
	},
	Name:        list.ListNameFundingSource,
	Description: "Funding sources are used to identify where a projects value comes from.  This list appears when assigning a projects value.",
	Values: []*list.ListItem{
		{Value: "internal", Name: "Internal"},
		{Value: "external", Name: "External"},
	},
}

var valueCategoryList = list.List{
	ControlFields: common.ControlFields{
		ID: "list:value-cats",
	},
	Name:        list.ListNameValueCategory,
	Description: "Value categories are used to characterize the nature of a projects value.  This information informs decisions about priority.",
	Values: []*list.ListItem{
		{Value: "revenue", Name: "Revenue increase"},
		{Value: "tax-benefit", Name: "Tax write-off"},
		{Value: "risk-mitigation", Name: "Risk mitigation"},
		{Value: "cost-reduction", Name: "Cost reduction"},
	},
}
