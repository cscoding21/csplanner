package surreal

import (
	"fmt"
	"strings"

	"csserver/internal/common"

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
func parseCountFromSurrealResult(raw interface{}) (*int, error) {
	var countSlice []map[string]int

	_, err := surrealdb.UnmarshalRaw(raw, &countSlice)
	if err != nil {
		return common.HandleReturnWithValue[int](nil, err)
	} else if len(countSlice) == 0 {
		return common.HandleReturnWithValue[int](common.ValToRef(0), nil)
	}

	countMap := countSlice[0]
	count := countMap["count"]

	return &count, nil
}

func buildWhereClauseFromFilters(filters *common.Filters) (string, map[string]interface{}) {
	if filters == nil || len(filters.Filters) == 0 {
		return "", filters.GetFiltersAsMap()
	}

	builder := strings.Builder{}

	for _, f := range filters.Filters {
		stringFormat := getPhraseFromFilterOp(f.Operation)
		ky := strings.Replace(f.Key, ".", "_", 1)
		builder.WriteString(fmt.Sprintf(stringFormat, f.Key, ky))
	}

	return builder.String(), filters.GetFiltersAsMap()
}

func getPhraseFromFilterOp(fo common.FilterOperation) string {
	out := ""

	switch fo {
	case (common.FilterOperationEqual):
		return "AND %s = $%v \n"
	case (common.FilterOperationNotEqual):
		return "AND %s != $%v \n"
	case (common.FilterOperationGreater):
		return "AND %s > $%v \n"
	case (common.FilterOperationLess):
		return "AND %s < $%v \n"
	case (common.FilterOperationGreaterOrEqual):
		return "AND %s >= $%v \n"
	case (common.FilterOperationLessOrEqual):
		return "AND %s <= $%v \n"
	case (common.FilterOperationIn):
		return "AND %s IN ($%v) \n"
	case (common.FilterOperationNotIn):
		return "AND %s NOT IN ($%v) \n"
	case (common.FilterFuzzyLike):
		return "AND %s ~ $%v \n"
	}

	return out
}

/*
FilterOperationEqual          FilterOperation = "eq"
	FilterOperationNotEqual       FilterOperation = "ne"
	FilterOperationGreater        FilterOperation = "gt"
	FilterOperationLess           FilterOperation = "lt"
	FilterOperationGreaterOrEqual FilterOperation = "ge"
	FilterOperationLessOrEqual    FilterOperation = "le"
	FilterOperationIn             FilterOperation = "in"
	FilterOperationNotIn          FilterOperation = "nin"
*/
