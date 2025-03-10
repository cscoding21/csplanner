package surreal

/*
func TestBuildWhereClauseFromFilters(t *testing.T) {
	testCases := []struct {
		have    common.Filters
		wantSql string
		wantMap map[string]interface{}
	}{
		{have: common.Filters{Filters: []common.Filter{{Key: "name", Value: "jeph", Operation: common.FilterOperationEqual}}},
			wantSql: "AND name = $name \n",
			wantMap: map[string]interface{}{
				"name": "jeph",
			},
		},
	}

	for _, tc := range testCases {
		sql, m := buildWhereClauseFromFilters(&tc.have)

		if tc.wantSql != sql {
			t.Errorf("Sql is incorrect for %s - want %s, got %s", tc.have, tc.wantSql, sql)
		}

		if !reflect.DeepEqual(tc.wantMap, m) {
			t.Errorf("Map is incorrect for %v - %v", tc.wantMap, m)
		}
	}
}
*/
