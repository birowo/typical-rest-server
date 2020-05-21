package dbkit_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/typical-go/typical-rest-server/pkg/dbkit"
)

func TestSelectOption_Where(t *testing.T) {
	testcases := []SelectTestCase{
		{
			SelectOption: dbkit.Equal("", ""),
			Builder:      sq.Select("name", "version").From("sometables"),
			ExpectedErr:  "Filter column can't be empty",
		},
		{
			SelectOption: dbkit.Equal("name", "dummy-name"),
			Builder:      sq.Select("name", "version").From("sometables"),
			Expected:     "SELECT name, version FROM sometables WHERE name = ?",
			ExpectedArgs: []interface{}{"dummy-name"},
		},
		{
			SelectOption: dbkit.Like("name", "%dum%"),
			Builder:      sq.Select("name", "version").From("sometables"),
			Expected:     "SELECT name, version FROM sometables WHERE name LIKE ?",
			ExpectedArgs: []interface{}{"%dum%"},
		},
	}

	for _, tt := range testcases {
		tt.Execute(t)
	}
}