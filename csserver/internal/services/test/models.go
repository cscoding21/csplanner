//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package test

import (
	"csserver/internal/common"
)

type Test struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
}
