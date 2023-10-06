package euro

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
		{
			actual:      selectMatchDrawStr,
			expected:    "SELECT draw_date, day_of_week, ball1, ball2, ball3, ball4, ball5, ls1, ls2, uk_marker, draw_no FROM euro WHERE ball1=? OR ball2=? OR ball3=? OR ball4=? OR ball5=? OR ls1=? OR ls2=?",
			description: "selectMatchDrawStr",
		},
	}
	for i, tc := range testcases {
		if tc.actual != tc.expected {
			t.Errorf("Case: %d Description: %s Expected: %s Actual: %s", i, tc.description, tc.expected, tc.actual)
		}
	}
}

func Example_listAllDraw() {
	db, err := DBConn()
	if err != nil {
		fmt.Printf("Unable to connect to DB: %v", err)
		return
	}
	defer db.Close()

	err = CreateTable(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBCreateTbl) {
		fmt.Printf("Unable to create table: %v", err)
		return
	}

	stmt, err := prepareInsertDrawStmt(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBPrepareStmt) {
		fmt.Printf("Prepare insert statement: %v ", err)
		return
	}

	d := Draw{
		DrawDate:  time.Date(2023, time.January, 20, 12, 0, 0, 0, time.UTC),
		DayOfWeek: time.Date(2023, time.January, 20, 12, 0, 0, 0, time.UTC).Weekday(),
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
		fmt.Printf("Insert draw. %v", err)
	}

	draws, err := listAllDraw(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBQueryTbl) {
		fmt.Printf("Query table. %v", err)
	}
	if len(draws) != 1 {
		fmt.Printf("Expected: 1 Actual: %v", len(draws))
	}

	for _, d := range draws {
		fmt.Printf("draw: %v\n", d)
	}

	// Output:
	// draw: {2023-01-20 12:00:00 +0000 GMT Friday 1 2 3 4 5 1 2 uk marker 1234}
}

func Example_matchDraw() {
	db, err := DBConn()
	if err != nil {
		fmt.Printf("Unable to connect to DB: %v", err)
		return
	}
	defer db.Close()

	err = CreateTable(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBCreateTbl) {
		fmt.Printf("Unable to create table: %v", err)
		return
	}

	stmt, err := prepareInsertDrawStmt(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBPrepareStmt) {
		fmt.Printf("Prepare insert statement: %v ", err)
		return
	}

	draws := []Draw{
		{
			DrawDate:  time.Date(2003, time.January, 10, 12, 00, 00, 0, time.UTC),
			DayOfWeek: time.Date(2003, time.January, 10, 12, 00, 00, 0, time.UTC).Weekday(),
			Ball1:     1,
			Ball2:     2,
			Ball3:     3,
			Ball4:     4,
			Ball5:     5,
			LS1:       1,
			LS2:       2,
			UKMarker:  "uk marker",
			DrawNo:    1233,
		},
		{
			DrawDate:  time.Date(2003, time.February, 10, 12, 00, 00, 0, time.UTC),
			DayOfWeek: time.Date(2003, time.February, 10, 12, 00, 00, 0, time.UTC).Weekday(),
			Ball1:     2,
			Ball2:     10,
			Ball3:     20,
			Ball4:     27,
			Ball5:     41,
			LS1:       3,
			LS2:       6,
			UKMarker:  "uk marker",
			DrawNo:    1234,
		},
		{
			DrawDate:  time.Date(2003, time.March, 10, 12, 00, 00, 0, time.UTC),
			DayOfWeek: time.Date(2003, time.March, 10, 12, 00, 00, 0, time.UTC).Weekday(),
			Ball1:     1,
			Ball2:     10,
			Ball3:     20,
			Ball4:     27,
			Ball5:     40,
			LS1:       2,
			LS2:       5,
			UKMarker:  "uk marker",
			DrawNo:    1235,
		},
		{
			DrawDate:  time.Date(2003, time.April, 10, 12, 00, 00, 0, time.UTC),
			DayOfWeek: time.Date(2003, time.April, 10, 12, 00, 00, 0, time.UTC).Weekday(),
			Ball1:     3,
			Ball2:     27,
			Ball3:     30,
			Ball4:     35,
			Ball5:     40,
			LS1:       3,
			LS2:       6,
			UKMarker:  "uk marker",
			DrawNo:    1236,
		},
		{
			DrawDate:  time.Date(2003, time.May, 10, 12, 00, 00, 0, time.UTC),
			DayOfWeek: time.Date(2003, time.May, 10, 12, 00, 00, 0, time.UTC).Weekday(),
			Ball1:     3,
			Ball2:     27,
			Ball3:     30,
			Ball4:     35,
			Ball5:     41,
			LS1:       3,
			LS2:       6,
			UKMarker:  "uk marker",
			DrawNo:    1237,
		},
	}

	for _, d := range draws {
		_, err = insertDraw(context.TODO(), stmt, d)
		if errors.Is(err, dbutil.ErrDBInsertTbl) {
			fmt.Printf("Insert draw. %v", err)
		}
	}

	mdraws, err := matchDraw(context.TODO(), db, 1, 12, 15, 21, 40, 1, 1)
	if err != nil {
		fmt.Printf("Match not found. %v", err)
	}

	fmt.Printf("Number of input: %d Number of match: %d\n", len(draws), len(mdraws))

	for _, md := range mdraws {
		fmt.Println(md)
	}

	// Output:
	// Number of input: 5 Number of match: 3
	// {2003-01-10 12:00:00 +0000 GMT Friday 1 2 3 4 5 1 2 uk marker 1233}
	// {2003-03-10 12:00:00 +0000 GMT Monday 1 10 20 27 40 2 5 uk marker 1235}
	// {2003-04-10 13:00:00 +0100 BST Thursday 3 27 30 35 40 3 6 uk marker 1236}
}
