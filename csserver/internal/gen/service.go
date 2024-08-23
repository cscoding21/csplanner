package gen

import (
	"csserver/internal/utils"
	"errors"
	"os"
	"path"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/cscoding21/csgen"
	csvalGen "github.com/cscoding21/csval/gen"
)

type GenProps struct {
	ServiceName       string
	ServiceLower      string
	OutputPath        string
	IncludeValidation bool
	IDTag             string
	CFTag             string
}

func GenService(props GenProps) error {
	caser := cases.Title(language.English)
	//---ensure struct name is title case
	props.ServiceName = caser.String(props.ServiceName)
	props.IDTag = utils.WrapInBackticks("json:\"id,omitempty\" csval:\"required\"")
	props.CFTag = utils.WrapInBackticks("json:\"control_fields\" csval:\"validate\"")

	//---create an all-lowercase version for the database & file names
	props.ServiceLower = strings.ToLower(props.ServiceName)

	servicePath := path.Join(props.OutputPath, props.ServiceLower)
	fullModelPath := path.Join(servicePath, "models.go")

	err := os.MkdirAll(servicePath, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = os.Open(fullModelPath)
	if errors.Is(err, os.ErrNotExist) {
		// handle the case where the file doesn't exist
		modelBuilder := strings.Builder{}

		modelBuilder.WriteString(csgen.ExecuteTemplate("models", modelsTemplateString, props))

		err = csgen.WriteGeneratedGoFile(fullModelPath, modelBuilder.String())
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	needsVal, err := csvalGen.CheckFileHasValidatorTags(fullModelPath)
	if err != nil {
		return err
	}

	props.IncludeValidation = needsVal

	//---create a builder for the service file
	builder := csgen.NewCSGenBuilderForFile("csplanner", props.ServiceLower)

	builder.WriteString(csgen.ExecuteTemplate("identifier", identifierTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_struct", serviceStructTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_getbyid", getByIDTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_findall", findAllTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_create", createTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_update", updateTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_delete", deleteTemplateString, props))

	fullPath := csgen.GetFileName(
		"csplanner",
		servicePath,
		props.ServiceLower)

	err = csgen.WriteGeneratedGoFile(fullPath, builder.String())
	if err != nil {
		return err
	}

	err = csvalGen.Generate(fullModelPath)
	if err != nil {
		return err
	}

	return nil
}

var identifierTemplateString = `
//---This is the name of the object in the database
const {{.ServiceName}}Identifier = "{{.ServiceLower}}"

`

var serviceStructTemplateString = `
// {{.ServiceName}}Service is a service for interacting with lists.
type {{.ServiceName}}Service struct {
	DBClient      surreal.DBClient
	ContextHelper interfaces.ContextHelpers
}

// New{{.ServiceName}}Service creates a new {{.ServiceName}} service.
func New{{.ServiceName}}Service(
	db surreal.DBClient,
	ch config.ContextHelper) *{{.ServiceName}}Service {

	return &{{.ServiceName}}Service{
		DBClient: db,
		ContextHelper: &ch,
	}
}

`

var deleteTemplateString = `
// Delete{{.ServiceName}} deletes a {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Delete{{.ServiceName}}(ctx context.Context, id string) error {
	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	list, err := s.Get{{.ServiceName}}ByID(ctx, id)
	if err != nil {
		return common.HandleReturn(err)
	}

	return common.HandleReturn(
		s.DBClient.SoftDeleteObject(userID, list),
	)
}

`

var getByIDTemplateString = `
// Get{{.ServiceName}}ByID gets a {{.ServiceName}} by its ID.
func (s *{{.ServiceName}}Service) Get{{.ServiceName}}ByID(ctx context.Context, id string) (*{{.ServiceName}}, error) {
	outData, err := s.DBClient.GetObjectById(id)
	if err != nil {
		return common.HandleReturnWithValue[{{.ServiceName}}](nil, err)
	}

	output, err := marshal.SurrealUnmarshal[{{.ServiceName}}](outData)

	return common.HandleReturnWithValue(
		output,
		err,
	)
}
	
`

var createTemplateString = `
// Create{{.ServiceName}} creates a new {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Create{{.ServiceName}}(ctx context.Context, input *{{.ServiceName}}) (*{{.ServiceName}}, error) {
{{if .IncludeValidation}}
	val := input.Validate()
	if !val.Pass {
		return common.HandleReturnWithValue[{{.ServiceName}}](nil, val.Error("{{.ServiceName}} validation failed"))
	}
{{end}}
	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	outData, err := s.DBClient.CreateObject(userID, {{.ServiceName }}Identifier, input)
	if err != nil {
		return common.HandleReturnWithValue[{{.ServiceName}}](nil, err)
	}

	outArray, err := marshal.SurrealUnmarshal[[]{{.ServiceName}}](outData)

	outObj := utils.RefToVal(outArray)[0]
	return common.HandleReturnWithValue(&outObj, err)
}
	
`

var updateTemplateString = `
// Update{{.ServiceName}} update an existing {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Update{{.ServiceName}}(ctx context.Context, input *{{.ServiceName}}) (*{{.ServiceName}}, error) {
{{if .IncludeValidation}}
	val := input.Validate()
	if !val.Pass {
		return common.HandleReturnWithValue[{{.ServiceName}}](nil, val.Error("{{.ServiceName}} validation failed"))
	}
{{end}}
	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	outData, err := s.DBClient.UpdateObject(userID, input.ID, input)
	if err != nil {
		return common.HandleReturnWithValue[{{.ServiceName}}](nil, err)
	}

	outArray, err := marshal.SurrealUnmarshal[[]{{.ServiceName}}](outData)

	outObj := utils.RefToVal(outArray)[0]
	return common.HandleReturnWithValue(&outObj, err)
}
	
`

var findAllTemplateString = `
// FindAll{{.ServiceName}}s return all {{.ServiceName}} in the system
func (s *{{.ServiceName}}Service) FindAll{{.ServiceName}}s(ctx context.Context) (common.PagedResults[{{.ServiceName}}], error) {
	pagingResults := common.NewPagedResultsForAllRecords[{{.ServiceName}}]()
	sql := "select * from {{.ServiceLower}} where deleted_at is null order by name"

	results, resultCount, err := s.DBClient.FindPagedObjects(sql, pagingResults.Pagination, pagingResults.Filters)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Pagination.TotalResults = &resultCount
	unpacked, err := marshal.SurrealSmartUnmarshal[[]{{.ServiceName}}](results)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Results = common.RefToVal(unpacked)
	return pagingResults, nil
}
	
`

var modelsTemplateString = `
//go:generate csval

//---NOTE TO DEVELOPERRS
//---This file will only be generated once.  Modify file as needed.

package {{.ServiceLower}}

import (
	"csserver/internal/common"
)

type {{.ServiceName}} struct {
	//---common for all DB objects
	ID string {{.IDTag}}
	ControlFields common.ControlFields {{.CFTag}}

	//---TODO: add fields here
}
`
