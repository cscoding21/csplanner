package augment

import (
	"csserver/internal/appserv/graph/idl"
	"csserver/internal/services/list"
)

func AugmentOrganization(org *idl.Organization) {
	if org == nil {
		return
	}

	templates := findTemplates()
	roles := findRoles()
	resources := findResources()
	lists := findLists()

	sl := getListById(*lists, list.ListNameSkills)
	fsl := getListById(*lists, list.ListNameFundingSource)
	vcl := getListById(*lists, list.ListNameValueCategory)

	org.Setup = &idl.OrganizationSetup{
		HasRoles:           len(*roles) > 0,
		HasTemplates:       len(*templates) > 0,
		HasResources:       len(*resources) > 0,
		HasSkills:          len(sl.Values) > 0,
		HasFundingSources:  len(fsl.Values) > 0,
		HasValueCategories: len(vcl.Values) > 0,
	}

	org.Setup.IsReadyForProjects = (org.Setup.HasRoles &&
		org.Setup.HasTemplates &&
		org.Setup.HasResources &&
		org.Setup.HasSkills &&
		org.Setup.HasFundingSources &&
		org.Setup.HasValueCategories)

}
