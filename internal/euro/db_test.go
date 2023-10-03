package euro

import (
	"context"
	"database/sql"
	"errors"
	"paulwizviz/lotterystat/internal/dbutil"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func DBConn() (*sql.DB, error) {
	return sql.Open("sqlite3", ":memory:")
}

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

func TestInsertDraw(t *testing.T) {
	db, err := DBConn()
	if err != nil {
		t.Errorf("Unable to connect: %v", err)
	}
	defer db.Close()

	err = CreateTable(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBCreateTbl) {
		t.Fatalf("Unable to create table: %v", err)
	}

	stmt, err := prepareInsertDrawStmt(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBPrepareStmt) {
		t.Errorf("Prepare insert statement: %v ", err)
	}

	now := time.Now()

	d := Draw{
		DrawDate:  now,
		DayOfWeek: now.Weekday(),
		Ball1:     1,
		Ball2:     2,
		Ball3:     3,
		Ball4:     4,
		Ball5:     5,
		LS1:       1,
		LS2:       2,
		UKMarker:  "uk marker",
		DrawNo:    1234,
	}

	_, err = insertDraw(context.TODO(), stmt, d)
	if errors.Is(err, dbutil.ErrDBInsertTbl) {
		t.Errorf("Insert draw. %v", err)
	}

	draws, err := ListAllDraw(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBQueryTbl) {
		t.Errorf("Query table. %v", err)
	}
	if len(draws) != 1 {
		t.Errorf("Expected: 1 Actual: %v", len(draws))
	}
}
