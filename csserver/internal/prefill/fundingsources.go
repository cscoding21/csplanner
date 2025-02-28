package prefill

import "csserver/internal/services/list"

var FundingSourcePrefillOptions = map[string]Prefill[list.ListItem]{
	"External": {
		Name: "External",
		Object: list.ListItem{
			Name:  "External",
			Value: "external",
		},
	},
	"Internal": {
		Name: "Internal",
		Object: list.ListItem{
			Name:  "Internal",
			Value: "internal",
		},
	},
}
