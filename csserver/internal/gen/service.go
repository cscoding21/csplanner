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
	CFTag             string
}

func GenService(props GenProps) error {
	caser := cases.Title(language.English)
	//---ensure struct name is title case
	props.ServiceName = caser.String(props.ServiceName)
	props.CFTag = utils.WrapInBackticks("csval:\"validate\"")

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
	builder.WriteString(csgen.ExecuteTemplate("service_find", findTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_create", createTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_update", updateTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("service_upsert", upsertTemplateString, props))
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
	PubSub nats.PubSubProvider
}

// New{{.ServiceName}}Service creates a new {{.ServiceName}} service.
func New{{.ServiceName}}Service(
	db surreal.DBClient,
	ch config.ContextHelper,
	ps nats.PubSubProvider) *{{.ServiceName}}Service {

	return &{{.ServiceName}}Service{
		DBClient: db,
		ContextHelper: &ch,
		PubSub: ps,
	}
}

`

var deleteTemplateString = `
// Delete{{.ServiceName}} deletes a {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Delete{{.ServiceName}}(ctx context.Context, id string) error {
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	list, err := s.Get{{.ServiceName}}ByID(ctx, id)
	if err != nil {
		return common.HandleReturn(err)
	}

	return common.HandleReturn(
		s.DBClient.SoftDeleteObject(userEmail, list),
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
func (s *{{.ServiceName}}Service) Create{{.ServiceName}}(ctx context.Context, input *{{.ServiceName}}) (common.UpdateResult[{{.ServiceName}}], error) {
{{if .IncludeValidation}}
	val := input.Validate()
	if !val.Pass {
		return common.NewUpdateResult[{{.ServiceName}}](&val, input), fmt.Errorf("validation failed")
	}
{{else}}
	val := validate.NewSuccessValidationResult()
{{end}}
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	outData, err := s.DBClient.CreateObject(userEmail, {{.ServiceName }}Identifier, input)
	if err != nil {
		return common.NewUpdateResult[{{.ServiceName}}](&val, input), err
	}

	outArray, err := marshal.SurrealUnmarshal[[]{{.ServiceName}}](outData)
	if err != nil {
		return common.NewUpdateResult[{{.ServiceName}}](&val, input), err
	}

	outObj := utils.RefToVal(outArray)[0]
	return common.NewUpdateResult[{{.ServiceName}}](&val, &outObj), nil
}
	
`

var updateTemplateString = `
// Update{{.ServiceName}} update an existing {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Update{{.ServiceName}}(ctx context.Context, input *{{.ServiceName}}) (common.UpdateResult[{{.ServiceName}}], error) {
{{if .IncludeValidation}}
	val := input.Validate()
	if !val.Pass {
		return common.NewUpdateResult[{{.ServiceName}}](&val, input), fmt.Errorf("validation failed")
	}
{{else}}
	val := validate.NewSuccessValidationResult()
{{end}}
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	lastObj, err := s.Get{{.ServiceName}}ByID(ctx, input.ID)
	if err != nil {
		return common.NewUpdateResult[{{.ServiceName}}](&val, input), err
	}

	//---ensure the integrity of the creation data
	input.CreatedAt = lastObj.CreatedAt
	input.CreatedBy = lastObj.CreatedBy

	outData, err := s.DBClient.UpdateObject(userEmail, input.ID, input)
	if err != nil {
		return common.NewUpdateResult[{{.ServiceName}}](&val, input), err
	}

	outObj, err := marshal.SurrealUnmarshal[{{.ServiceName}}](outData)
	if err != nil {
		return common.NewUpdateResult[{{.ServiceName}}](&val, input), err
	}

	return common.NewUpdateResult[{{.ServiceName}}](&val, outObj), nil
}

`

var upsertTemplateString = `
// Upsert{{.ServiceName}} create or update a {{.ServiceName}}
func (s *{{.ServiceName}}Service) Upsert{{.ServiceName}}(ctx context.Context, obj {{.ServiceName}}) (*common.UpdateResult[{{.ServiceName}}], error) {
	existingObj, _ := s.Get{{.ServiceName}}ByID(ctx, obj.ID)

	if existingObj == nil {
		resp, err := s.Create{{.ServiceName}}(ctx, &obj)
		return &resp, err
	}

	obj.ID = existingObj.ID
	resp, err := s.Update{{.ServiceName}}(ctx, &obj)
	return &resp, err
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

var findTemplateString = `
// Find{{.ServiceName}} return a paged list of {{.ServiceName}} based on filter criteria
func (s *{{.ServiceName}}Service) Find{{.ServiceName}}s(ctx context.Context, paging common.Pagination, filters common.Filters) (common.PagedResults[{{.ServiceName}}], error) {
	out := common.NewPagedResults[{{.ServiceName}}](paging, filters)

	whereSql, _ := s.DBClient.BuildWhereClauseFromFilters(&filters)

	sql := fmt.Sprintf("SELECT * FROM {{.ServiceLower}} WHERE true AND deleted_at is null %s ORDER BY name", whereSql)

	rawResults, count, err := s.DBClient.FindPagedObjects(sql, paging, filters)
	if err != nil {
		return out, err
	}

	out.Pagination.TotalResults = &count
	unpacked, err := marshal.SurrealSmartUnmarshal[[]{{.ServiceName}}](rawResults)
	if err != nil {
		return out, err
	}

	out.Results = common.RefToVal(unpacked)
	return out, nil
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
	common.ControlFields {{.CFTag}}

	//---TODO: add fields here
}
`
