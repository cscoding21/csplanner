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
import (
	"context"
	"csserver/internal/common"
	"csserver/internal/providers/nats"
	"csserver/internal/providers/postgres"
	"fmt"

	"github.com/cscoding21/csval/validate"
	"github.com/jackc/pgx/v5"
)


//---This is the name of the object in the database
const {{.ServiceName}}Identifier = postgres.TableName("{{.ServiceLower}}")

`

var serviceStructTemplateString = `
// {{.ServiceName}}Service is a service for interacting with lists.
type {{.ServiceName}}Service struct {
	db      *pgxpool.Pool
	pubsub nats.PubSubProvider
}

// New{{.ServiceName}}Service creates a new {{.ServiceName}} service.
func New{{.ServiceName}}Service(
	db *pgxpool.Pool,
	ps nats.PubSubProvider) *{{.ServiceName}}Service {

	return &{{.ServiceName}}Service{
		db: db,
		pubsub: ps,
	}
}

`

var deleteTemplateString = `
// Delete{{.ServiceName}} deletes a {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Delete{{.ServiceName}}(ctx context.Context, id string) error {
	return postgres.SoftDelete(ctx, s.db, {{.ServiceName}}Identifier, id)
}

`

var getByIDTemplateString = `
// Get{{.ServiceName}}ByID gets a {{.ServiceName}} by its ID.
func (s *{{.ServiceName}}Service) Get{{.ServiceName}}ByID(ctx context.Context, id string) (*common.BaseModel[{{.ServiceName}}], error) {
	return postgres.GetObjectByID[{{.ServiceName}}](ctx, s.db, {{.ServiceName}}Identifier, id)
}
	
`

var createTemplateString = `
// Create{{.ServiceName}} creates a new {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Create{{.ServiceName}}(ctx context.Context, input {{.ServiceName}}) (common.UpdateResult[*common.BaseModel[{{.ServiceName}}]], error) {
{{if .IncludeValidation}}
	val := input.Validate()
	if !val.Pass {
		return common.NewUpdateResult[*common.BaseModel[{{.ServiceName}}]](val, nil), fmt.Errorf("validation failed")
	}
{{else}}
	val := validate.NewSuccessValidationResult()
{{end}}
	obj, err := postgres.UpdateObject(ctx, s.db, input, {{.ServiceName}}Identifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}
	
`

var updateTemplateString = `
// Update{{.ServiceName}} update an existing {{.ServiceName}}.
func (s *{{.ServiceName}}Service) Update{{.ServiceName}}(ctx context.Context, input {{.ServiceName}}) (common.UpdateResult[*common.BaseModel[{{.ServiceName}}]], error) {
{{if .IncludeValidation}}
	val := input.Validate()
	if !val.Pass {
		return common.NewUpdateResult[*common.BaseModel[{{.ServiceName}}]](val, nil), fmt.Errorf("validation failed")
	}
{{else}}
	val := validate.NewSuccessValidationResult()
{{end}}
	obj, err := postgres.UpdateObject(ctx, s.db, input, {{.ServiceName}}Identifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

`

var upsertTemplateString = `
// Upsert{{.ServiceName}} create or update a {{.ServiceName}}
func (s *{{.ServiceName}}Service) Upsert{{.ServiceName}}(ctx context.Context, input {{.ServiceName}}) (common.UpdateResult[*common.BaseModel[{{.ServiceName}}]], error) {
{{if .IncludeValidation}}
	val := input.Validate()
	if !val.Pass {
		return common.NewUpdateResult[*common.BaseModel[{{.ServiceName}}]](val, nil), fmt.Errorf("validation failed")
	}
{{else}}
	val := validate.NewSuccessValidationResult()
{{end}}
	obj, err := postgres.UpdateObject(ctx, s.db, input, {{.ServiceName}}Identifier, input.ID)
	if err != nil {
		return common.NewUpdateResult(val, &obj), err
	}

	return common.NewUpdateResult(val, &obj), nil
}

`

var findAllTemplateString = `
// FindAll{{.ServiceName}}s return all {{.ServiceName}} in the system
func (s *{{.ServiceName}}Service) FindAll{{.ServiceName}}s(ctx context.Context) (common.PagedResults[common.BaseModel[{{.ServiceName}}]], error) {
	return postgres.FindAllObjects[{{.ServiceName}}](ctx, s.db, {{.ServiceName}}Identifier)
}
	
`

var findTemplateString = `
// Find{{.ServiceName}} return a paged list of {{.ServiceName}} based on filter criteria
func (s *{{.ServiceName}}Service) Find{{.ServiceName}}s(ctx context.Context, paging common.Pagination, filters common.Filters) (common.PagedResults[common.BaseModel[{{.ServiceName}}]], error) {
	out := common.NewPagedResults[{{.ServiceName}}](paging, filters)

	whereSql, params := postgres.BuildWhereClauseFromFilters(&filters)

	sql := fmt.Sprintf("SELECT * FROM %s WHERE true AND deleted_at is null %s ORDER BY created_at DESC", {{.ServiceName}}Identifier, whereSql)

	return postgres.FindPagedObjects[{{.ServiceName}}](ctx, s.db, sql, out.Pagination, out.Filters, params)
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
