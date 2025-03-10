package postgres

import (
	"csserver/internal/common"
	"fmt"
	"testing"
)

func TestConvertDotsToPG(t *testing.T) {
	testData := []struct {
		val            string
		expected       string
		lastItemIsJson bool
	}{
		{val: "data.basics.name", expected: "data->'basics'->>'name'"},
		{val: "data.id", expected: "data->>'id'"},
		{val: "data.basics.name", expected: "data->'basics'->>'name'"},
	}

	for _, d := range testData {
		result := convertDotsToPG(d.val, d.lastItemIsJson)

		if result != d.expected {
			t.Errorf("unexpected result.  Wanted %s - got %s", d.expected, result)
		}
	}
}

func TestConvertMultiParams(t *testing.T) {
	testData := []struct {
		val      string
		expected string
		count    int
		pointer  int
	}{
		{val: "scheduled,inflight", expected: "($1, $2)", count: 2, pointer: 1},
		{val: "scheduled", expected: "($1)", count: 1, pointer: 1},
	}

	for _, d := range testData {
		str, params := convertMultiParams(d.val, d.pointer)

		if len(params) != d.count {
			t.Errorf("unexpected number of params returned.  Wanted %v - got %v", d.count, len(params))
		}

		if str != d.expected {
			t.Errorf("unexpected string returned.  Wanted %s - got %s", d.expected, str)
		}
	}
}

func TestBuildWhereClauseFromFilters(t *testing.T) {
	filters := common.Filters{
		Filters: []common.Filter{
			{Key: "data.status.status", Value: "inflight, scheduled", Operation: common.FilterOperationIn},
			{Key: "data.basics.owner_id", Value: "jeph@jmk21.com", Operation: common.FilterOperationEqual},
			{Key: "data.skills.id", Value: "backend", Operation: common.FilterContains},
		}}

	sql, params := BuildWhereClauseFromFilters(&filters)

	fmt.Println(sql)
	fmt.Println(params)
}
