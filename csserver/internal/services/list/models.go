//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package list

import (
	"csserver/internal/common"
)

type List struct {
	//---common for all DB objects
	common.ControlFields

	//---TODO: add fields here
	Name   string     `json:"name" csval:"req"`
	Values []ListItem `json:"values"`
}

type ListItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
