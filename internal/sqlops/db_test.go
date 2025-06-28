package sqlops_test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/stretchr/testify/assert"
)

func TestDriverType(t *testing.T) {
	testcases := []struct {
		name      string
		sqlDriver func() (*sql.DB, error)
		expected  sqlops.SQLType
	}{
		{
			name: "github.com/mattn/go-sqlite3",
			sqlDriver: func() (*sql.DB, error) {
				return sqlops.NewSQLiteMem()
			},
			expected: sqlops.SQLiteType,
		},
		{
			name: "github.com/lib/pq",
			sqlDriver: func() (*sql.DB, error) {
				return sqlops.NewPSQL("a", "b", "c", 123, "efg")
			},
			expected: sqlops.PSQLType,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d-%s", i, tc.name), func(t *testing.T) {
			db, _ := tc.sqlDriver()
			actual := sqlops.DriverType(db)
			assert.Equal(t, tc.expected, actual)
		})

	}
}
