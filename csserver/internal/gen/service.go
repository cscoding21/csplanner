package gen

import (
	"os"
	"path"
	"strings"

	"github.com/cscoding21/csgen"
)

type GenProps struct {
	ServiceName  string
	ServiceLower string
	OutputPath   string
}

func Generate(props GenProps) error {
	props.ServiceLower = strings.ToLower(props.ServiceName)

	builder := csgen.NewCSGenBuilderForFile("csplanner", strings.ToLower(props.ServiceName))

	builder.WriteString(csgen.ExecuteTemplate("list_identifier", identifierTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("list_service_struct", serviceStructTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("list_service_getbyid", getByIDTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("list_service_findall", findAllTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("list_service_create", createTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("list_service_update", updateTemplateString, props))
	builder.WriteString(csgen.ExecuteTemplate("list_service_delete", deleteTemplateString, props))

	path := path.Join(props.OutputPath, strings.ToLower(props.ServiceName))
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	fullPath := csgen.GetFileName(
		"csplanner",
		path,
		strings.ToLower(props.ServiceName))

	err = csgen.WriteGeneratedGoFile(fullPath, builder.String())
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
	Logger        interfaces.Logger
	ContextHelper interfaces.ContextHelpers
}

// New{{.ServiceName}}Service creates a new {{.ServiceName}} service.
func New{{.ServiceName}}Service(
	db surreal.DBClient,
	logger interfaces.Logger) *{{.ServiceName}}Service {

	return &{{.ServiceName}}Service{
		DBClient: db,
		Logger:   logger,
	}
}

`

var deleteTemplateString = `
// Delete{{.ServiceName}} deletes a {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Delete{{.ServiceName}}(ctx context.Context, id string) error {
	userID := s.ContextHelper.GetUserIDFromContext(ctx)
	s.Logger.Info(id)
	list, err := s.Get{{.ServiceName}}ByID(ctx, id)
	if err != nil {
		return common.HandleReturn(s.Logger, err)
	}

	return common.HandleReturn(
		s.Logger,
		s.DBClient.SoftDeleteObject(userID, list),
	)
}

`

var getByIDTemplateString = `
// Get{{.ServiceName}}ByID gets a {{.ServiceName}} by its ID.
func (s *{{.ServiceName}}Service) Get{{.ServiceName}}ByID(ctx context.Context, id string) (*{{.ServiceName}}, error) {
	outData, err := s.DBClient.GetObjectById(id)
	if err != nil {
		return common.HandleReturnWithValue[{{.ServiceName}}](s.Logger, nil, err)
	}

	output, err := common.SurrealUnmarshal[{{.ServiceName}}](s.Logger, outData)

	return common.HandleReturnWithValue(
		s.Logger,
		output,
		err,
	)
}
	
`

var createTemplateString = `
// {{.ServiceName}}List creates a new {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Create{{.ServiceName}}(ctx context.Context, input *{{.ServiceName}}) (*{{.ServiceName}}, error) {
	val := input.Validate()
	if !val.Pass {
		return common.HandleReturnWithValue[{{.ServiceName}}](s.Logger, nil, val.Error("{{.ServiceName}} validation failed"))
	}

	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	outData, err := s.DBClient.CreateObject(userID, {{.ServiceName }}Identifier, input)
	if err != nil {
		return common.HandleReturnWithValue[{{.ServiceName}}](s.Logger, nil, err)
	}

	list, err := common.SurrealSmartUnmarshal[{{.ServiceName}}](s.Logger, outData)

	return common.HandleReturnWithValue(s.Logger, list, err)
}
	
`

var updateTemplateString = `
// Update{{.ServiceName}} update an existing {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Update{{.ServiceName}}(ctx context.Context, input *{{.ServiceName}}) (*{{.ServiceName}}, error) {
	userID := s.ContextHelper.GetUserIDFromContext(ctx)

	outData, err := s.DBClient.UpdateObject(userID, input.ID, input)
	if err != nil {
		return common.HandleReturnWithValue[{{.ServiceName}}](s.Logger, nil, err)
	}

	output, err := common.SurrealUnmarshal[{{.ServiceName}}](s.Logger, outData)

	return common.HandleReturnWithValue(s.Logger, output, err)
}
	
`

var findAllTemplateString = `
// FindAll{{.ServiceName}} return all {{.ServiceName}} in the system
func (s *{{.ServiceName}}Service) FindAll{{.ServiceName}}(ctx context.Context) (common.PagedResults[{{.ServiceName}}], error) {
	pagingResults := common.NewPagedResultsForAllRecords[{{.ServiceName}}]()
	sql := "select * from {{.ServiceLower}} where deleted_at is null order by name"

	results, resultCount, err := s.DBClient.FindPagedObjects(sql, pagingResults.Pagination, pagingResults.Filters)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Pagination.TotalResults = &resultCount
	unpacked, err := common.SurrealSmartUnmarshal[[]{{.ServiceName}}](s.Logger, results)
	if err != nil {
		return pagingResults, err
	}

	pagingResults.Results = common.RefToVal(unpacked)
	return pagingResults, nil
}
	
`
