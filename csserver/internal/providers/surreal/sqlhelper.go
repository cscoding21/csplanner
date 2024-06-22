package surreal

import (
	"strings"

	common "csserver/internal/common"
	commonInterfaces "csserver/internal/interfaces"

	"github.com/surrealdb/surrealdb.go"
)

// getCountSql modify a SQL query so that it returns a count
func getCountSql(sql string) string {
	var builder strings.Builder

	builder.WriteString("SELECT COUNT() FROM (")
	builder.WriteString(sql)
	builder.WriteString(") GROUP ALL;")

	return builder.String()
}

// getPageSql amend a SQL query so that it support pagination
func getPageSql(sql string) string {
	return sql + " LIMIT $limit START $start;"
}

// parseCountFromSurrealResult take the results of a scalar query that returns an int and parse the returned value
func parseCountFromSurrealResult(logger commonInterfaces.Logger, raw interface{}) (*int, error) {
	var countSlice []map[string]int

	_, err := surrealdb.UnmarshalRaw(raw, &countSlice)
	if err != nil {
		return common.HandleReturnWithValue[int](logger, nil, err)
	} else if len(countSlice) == 0 {
		return common.HandleReturnWithValue[int](logger, common.ValToRef(0), nil)
	}

	countMap := countSlice[0]
	count := countMap["count"]

	return &count, nil
}
