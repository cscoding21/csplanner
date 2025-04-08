package prefill

import (
	"csserver/internal/services/list"
)

var valueCategoryPrefillOptions = map[string]Prefill[list.ListItem]{
	"Revenue": {
		Name: "Revenue",
		Object: list.ListItem{
			Name:  "Revenue",
			Value: "revenue",
		},
	},
	"Cost": {
		Name: "Cost Savings",
		Object: list.ListItem{
			Name:  "Cost Savings",
			Value: "cost-savings",
		},
	},
	"Risk": {
		Name: "Risk Mitigation",
		Object: list.ListItem{
			Name:  "Risk Mitigation",
			Value: "risk-mitigation",
		},
	},
	"Tax": {
		Name: "Tax Benefit",
		Object: list.ListItem{
			Name:  "Tax Benefit",
			Value: "tax-benefit",
		},
	},
}
