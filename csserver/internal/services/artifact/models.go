//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package artifact

import (
	"csserver/internal/common"
)

type Artifact struct {
	//---common for all DB objects
	common.ControlFields `csval:"validate"`

	//---TODO: add fields here
	Context        string `json:"context"`
	FileType       string `json:"file_type"`
	FileName       string `json:"file_name"`
	SizeInBytes    int64  `json:"size_in_bytes"`
	ClientFileName string `json:"client_file_name"`
	//ArtifactType   artifacttype.ArtifactType `json:"artifact_type"`
	Path     string      `json:"path"`
	FileInfo interface{} `json:"file_info"`
}
