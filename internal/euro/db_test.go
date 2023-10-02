package euro

import (
	"testing"
)

func TestStmtStr(t *testing.T) {
	testcases := []struct {
		actual      string
		expected    string
		description string
	}{
		{
			actual:      createTableStmtStr,
			expected:    "CREATE TABLE IF NOT EXISTS euro (draw_date INTEGER,day_of_week INTEGER,ball1 INTEGER,ball2 INTEGER,ball3 INTEGER,ball4 INTEGER,ball5 INTEGER,ls1 INTEGER,ls2 INTEGER,uk_marker TEXT,draw_no INTEGER PRIMARY KEY)",
			description: "createTableStmtStr",
		},
		{
			actual:      insertDrawStmtStr,
			expected:    "INSERT INTO euro (draw_date,day_of_week,ball1,ball2,ball3, ball4,ball5,ls1,ls2,uk_marker,draw_no) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )",
			description: "insertDrawStmtStr",
		},
		{
			actual:      selectAllStmtStr,
			expected:    "SELECT * FROM euro",
			description: "selectAllStmtStr",
		},
	}
	for i, tc := range testcases {
		if tc.actual != tc.expected {
			t.Errorf("Case: %d Description: %s Expected: %s Actual: %s", i, tc.description, tc.expected, tc.actual)
		}
	}
}
