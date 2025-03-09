package postgres

import (
	"csserver/internal/common"
	"fmt"
	"strings"
)

func GetPageSql(sql string, paramNum int) string {
	return sql + fmt.Sprintf(" LIMIT $%v OFFSET $%v;", paramNum, (paramNum+1))
}

func GetCountSql(sql string) string {
	var builder strings.Builder

	builder.WriteString("SELECT COUNT(*) FROM (")
	builder.WriteString(sql)
	builder.WriteString(");")

	return builder.String()
}

func BuildWhereClauseFromFilters(filters *common.Filters) (string, map[string]interface{}) {
	if filters == nil || len(filters.Filters) == 0 {
		return "", filters.GetFiltersAsMap()
	}

	builder := strings.Builder{}

	for i, f := range filters.Filters {
		stringFormat := getPhraseFromFilterOp(f)
		ky := strings.Replace(f.Key, ".", "_", 1)
		sqlString := ""

		if f.Operation == common.FilterCustom {
			sqlString = stringFormat
		} else {
			sqlString = fmt.Sprintf(stringFormat, ky, (i + 1))
		}

		builder.WriteString(sqlString)
	}

	return builder.String(), filters.GetFiltersAsMap()
}

func getPhraseFromFilterOp(fo common.Filter) string {
	out := ""

	switch fo.Operation {
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
	case (common.FilterContains):
		return "AND %s CONTAINS $%v \n"
	case (common.FilterCustom):
		return GetCustomFilter(*fo.CustomName)
	}

	return out
}

func GetCustomFilter(name string) string {
	customFilters := loadFilters()
	out, ok := customFilters[name]
	if !ok {
		fmt.Printf("cannot find custom filter key %s", name)
		return ""
	}

	return out
}

func loadFilters() map[string]string {
	f := make(map[string]string)

	f["resource_in_project"] = `
	AND (
        calculated.team ANYINSIDE [$resourceID]
    )
	`

	return f
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
FilterCustom          FilterOperation = "custom"
*/
