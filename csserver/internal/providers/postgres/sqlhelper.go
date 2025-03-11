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

func BuildWhereClauseFromFilters(filters *common.Filters) (string, []any) {
	if filters == nil || len(filters.Filters) == 0 {
		return "", []any{}
	}

	sqlBuilder := strings.Builder{}
	outParams := []any{}

	for _, f := range filters.Filters {
		pointer := len(outParams) + 1
		stringFormat, params := getPhraseFromFilterOp(f, pointer)
		outParams = append(outParams, params...)
		sqlString := ""

		if f.Operation == common.FilterCustom {
			sqlString = stringFormat
		} else {
			sqlString = stringFormat
		}

		sqlBuilder.WriteString(sqlString)
	}

	return sqlBuilder.String(), outParams
}

func getPhraseFromFilterOp(fo common.Filter, pointer int) (string, []any) {
	outPhrase := ""
	outParams := []any{}

	ky := convertDotsToPG(fo.Key, fo.Operation == common.FilterContains)

	switch fo.Operation {
	case (common.FilterOperationEqual):
		outParams = append(outParams, fo.Value)
		outPhrase = fmt.Sprintf("AND %s = $%v \n", ky, pointer)
	case (common.FilterOperationNotEqual):
		outParams = append(outParams, fo.Value)
		outPhrase = fmt.Sprintf("AND %s != $%v \n", ky, pointer)
	case (common.FilterOperationGreater):
		outParams = append(outParams, fo.Value)
		outPhrase = fmt.Sprintf("AND %s > $%v \n", ky, pointer)
	case (common.FilterOperationLess):
		outParams = append(outParams, fo.Value)
		outPhrase = fmt.Sprintf("AND %s < $%v \n", ky, pointer)
	case (common.FilterOperationGreaterOrEqual):
		outParams = append(outParams, fo.Value)
		outPhrase = fmt.Sprintf("AND %s >= $%v \n", ky, pointer)
	case (common.FilterOperationLessOrEqual):
		outParams = append(outParams, fo.Value)
		outPhrase = fmt.Sprintf("AND %s <= $%v \n", ky, pointer)
	case (common.FilterOperationIn):
		parens, params := convertMultiParams(fo.Value.(string), pointer)
		outParams = append(outParams, params...)
		outPhrase = fmt.Sprintf("AND %s IN "+parens+"\n", ky)
	case (common.FilterOperationNotIn):
		parens, params := convertMultiParams(fo.Value.(string), pointer)
		outParams = append(outParams, params...)
		outPhrase = fmt.Sprintf("AND %s NOT IN "+parens+"\n", ky)
	case (common.FilterFuzzyLike):
		outParams = append(outParams, fo.Value.(string))
		outPhrase = fmt.Sprintf("AND %s ILIKE '%%' || $%v || '%%'\n", ky, pointer)
	case (common.FilterContains):
		//---skip the parameter here as we're embedding into the filter line :(
		//outParams = append(outParams, fo.Value)
		outPhrase = getContainsFilter(fo, pointer)
	case (common.FilterCustom):
		op, params := GetCustomFilter(fo, pointer)
		outParams = append(outParams, params...)
		outPhrase = op
	}

	return outPhrase, outParams
}

func getContainsFilter(fo common.Filter, pointer int) string {
	key := ""
	field := ""

	if strings.Contains(fo.Key, ".") {
		sp := strings.Split(fo.Key, ".")
		field = sp[len(sp)-1]

		firstParts := sp[:len(sp)-1]

		key = getKeyFromDotsArray(firstParts, true)
	}

	//---TODO: try to get the value option parameterized
	str := fmt.Sprintf(`AND %s @> '[{"%s": "%s"}]'`, key, field, fo.Value)

	return str
}

func convertMultiParams(params string, pointer int) (string, []any) {
	outString := strings.Builder{}
	outParams := []any{}
	ps := params

	if strings.Contains(ps, ",") {
		psa := strings.Split(ps, ",")

		outString.WriteString("(")
		for i, param := range psa {
			if i == 0 {
				outString.WriteString(fmt.Sprintf("$%v", pointer))
			} else {
				outString.WriteString(fmt.Sprintf(", $%v", pointer))
			}

			outParams = append(outParams, param)
			pointer++
		}
		outString.WriteString(")")
	} else {
		outParams = append(outParams, params)
		outString.WriteString(fmt.Sprintf("($%v)", pointer))
	}

	return outString.String(), outParams
}

func convertDotsToPG(input string, targetIsJson bool) string {
	if strings.Contains(input, ".") {
		parts := strings.Split(input, ".")

		return getKeyFromDotsArray(parts, targetIsJson)
	}

	return input
}

func getKeyFromDotsArray(parts []string, targetIsJson bool) string {
	out := strings.Builder{}
	lastPart := len(parts) - 1

	for i, p := range parts {
		if i == 0 {
			out.WriteString(p)
		} else if i != lastPart {
			out.WriteString(fmt.Sprintf("->'%s'", p))
		} else {
			if targetIsJson {
				out.WriteString(fmt.Sprintf("->'%s'", p))
			} else {
				out.WriteString(fmt.Sprintf("->>'%s'", p))
			}

		}
	}

	return out.String()
}

func GetCustomFilter(fo common.Filter, pointer int) (string, []any) {
	outString := strings.Builder{}
	outParams := []any{}
	outParams = append(outParams, fo.Value)

	customFilters := loadFilters()
	cf, ok := customFilters[*fo.CustomName]
	if !ok {
		fmt.Printf("cannot find custom filter key %s", *fo.CustomName)
		return "", outParams
	}

	outString.WriteString(fmt.Sprintf(cf, pointer))

	return outString.String(), outParams
}

func loadFilters() map[string]string {
	f := make(map[string]string)

	f["resource_in_project"] = `
	AND (
        data->'calculated'->'team' ?| array[$%v]
    )
	`

	return f
}

// func sanitizeSqlInput(input string) string {

// }

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
