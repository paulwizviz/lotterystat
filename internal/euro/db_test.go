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

func dbConn() (*sql.DB, error) {
	return sql.Open("sqlite3", ":memory:")
}

func listTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM sqlite_schema WHERE type='table' ORDER BY name")
	if err != nil {
		return nil, err
	}
	var result []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Printf("Table name not found: %v", err)
		}
		result = append(result, name)
	}
	return result, nil
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

func Example_listTable() {
	db, err := dbConn()
	if err != nil {
		fmt.Printf("DB not found: %v", err)
		return
	}
	defer db.Close()
	err = createTable(context.TODO(), db)
	if err != nil {
		fmt.Printf("Create table error: %v", err)
		return
	}
	tbls, err := listTables(db)
	if err != nil {
		fmt.Printf("Unable to list tables: %v", err)
		return
	}

	for _, tbl := range tbls {
		fmt.Println(tbl)
	}

	// Output:
	// euro
}

func Example_listAllDraw() {
	db, err := dbConn()
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
	defer stmt.Close()

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
	db, err := dbConn()
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
	defer stmt.Close()

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
		{
			DrawDate:  time.Date(2003, time.May, 10, 12, 00, 00, 0, time.UTC),
			DayOfWeek: time.Date(2003, time.May, 10, 12, 00, 00, 0, time.UTC).Weekday(),
			Ball1:     1,
			Ball2:     12,
			Ball3:     15,
			Ball4:     21,
			Ball5:     40,
			LS1:       1,
			LS2:       1,
			UKMarker:  "uk marker",
			DrawNo:    1238,
		},
	}

	for _, d := range draws {
		_, err = insertDraw(context.TODO(), stmt, d)
		if errors.Is(err, dbutil.ErrDBInsertTbl) {
			fmt.Printf("Insert draw. %v", err)
		}
	}

	mStmt, err := prepareMatchDrawStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Prepare match draw statement. %v", err)
	}

	mdraws, err := matchDraw(context.TODO(), mStmt, 1, 12, 15, 21, 40, 1, 1)
	if err != nil {
		fmt.Printf("Match not found. %v", err)
	}

	fmt.Printf("Number of input: %d Number of match: %d\n", len(draws), len(mdraws))

	for _, md := range mdraws {
		fmt.Println(md)
	}

	// Output:
	// Number of input: 6 Number of match: 4
	// {2003-01-10 12:00:00 +0000 GMT Friday 1 2 3 4 5 1 2 uk marker 1233}
	// {2003-03-10 12:00:00 +0000 GMT Monday 1 10 20 27 40 2 5 uk marker 1235}
	// {2003-04-10 13:00:00 +0100 BST Thursday 3 27 30 35 40 3 6 uk marker 1236}
	// {2003-05-10 13:00:00 +0100 BST Saturday 1 12 15 21 40 1 1 uk marker 1238}
}

func Example_matchBet() {
	db, err := dbConn()
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
	defer stmt.Close()

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
		{
			DrawDate:  time.Date(2003, time.May, 10, 12, 00, 00, 0, time.UTC),
			DayOfWeek: time.Date(2003, time.May, 10, 12, 00, 00, 0, time.UTC).Weekday(),
			Ball1:     1,
			Ball2:     12,
			Ball3:     15,
			Ball4:     21,
			Ball5:     40,
			LS1:       1,
			LS2:       1,
			UKMarker:  "uk marker",
			DrawNo:    1238,
		},
	}

	for _, d := range draws {
		_, err = insertDraw(context.TODO(), stmt, d)
		if errors.Is(err, dbutil.ErrDBInsertTbl) {
			fmt.Printf("Insert draw. %v", err)
		}
	}

	bets := []Bet{
		{
			Ball1: 1,
			Ball2: 10,
			Ball3: 20,
			Ball4: 40,
			Ball5: 45,
			LS1:   2,
			LS2:   4,
		},
	}

	mbs, err := matchBets(context.TODO(), db, bets)
	if err != nil {
		fmt.Printf("Match bet error. %v", err)
	}

	for _, mb := range mbs {
		fmt.Printf("Bet: %v Draw: %v Balls match: %v Lucky stars: %v\n", mb.Bet, mb.Draw, mb.Balls, mb.LuckyStars)
	}

	// Output:
	// Bet: {1 10 20 40 45 2 4} Draw: {2003-01-10 12:00:00 +0000 GMT Friday 1 2 3 4 5 1 2 uk marker 1233} Balls match: [1] Lucky stars: []
	// Bet: {1 10 20 40 45 2 4} Draw: {2003-02-10 12:00:00 +0000 GMT Monday 2 10 20 27 41 3 6 uk marker 1234} Balls match: [10 20] Lucky stars: []
	// Bet: {1 10 20 40 45 2 4} Draw: {2003-03-10 12:00:00 +0000 GMT Monday 1 10 20 27 40 2 5 uk marker 1235} Balls match: [1 10 20] Lucky stars: [2]
	// Bet: {1 10 20 40 45 2 4} Draw: {2003-05-10 13:00:00 +0100 BST Saturday 1 12 15 21 40 1 1 uk marker 1238} Balls match: [1] Lucky stars: []
}
