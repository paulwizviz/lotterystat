package sforl

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
			expected:    "CREATE TABLE IF NOT EXISTS set_for_life (draw_date INTEGER,day_of_week INTEGER,ball1 INTEGER,ball2 INTEGER,ball3 INTEGER,ball4 INTEGER,ball5 INTEGER,lb INTEGER, ball_set TEXT,machine TEXT,draw_no INTEGER PRIMARY KEY)",
			description: "createTableStmtStr",
		},
		{
			actual:      insertDrawStmtStr,
			expected:    "INSERT INTO set_for_life (draw_date,day_of_week,ball1,ball2,ball3,ball4,ball5,lb,ball_set,machine,draw_no) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )",
			description: "insertDrawStmtStr",
		},
		{
			actual:      selectAllDrawStmtStr,
			expected:    "SELECT * FROM set_for_life",
			description: "selectAllDrawStmtStr",
		},
	}
	for i, tc := range testcases {
		if tc.actual != tc.expected {
			t.Errorf("Case: %d Description: %s Expected: %s Actual: %s", i, tc.description, tc.expected, tc.actual)
		}
	}
}
