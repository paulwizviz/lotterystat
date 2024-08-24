package euro

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"paulwizviz/lotterystat/internal/dbutil"
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

func Example_listTable() {
	db, err := dbConn()
	if err != nil {
		fmt.Printf("DB not found: %v", err)
		return
	}
	defer db.Close()
	err = createSQLiteTable(context.TODO(), db)
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

func Example_insertDraw() {
	db, err := dbConn()
	if err != nil {
		fmt.Printf("Unable to connect to DB: %v", err)
		return
	}
	defer db.Close()

	err = CreateSQLiteTable(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBCreateTbl) {
		fmt.Printf("Unable to create table: %v", err)
		return
	}

	stmt, err := prepInsertDrawStmt(context.TODO(), db)
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

	// Output:
	// draw: {2023-01-20 12:00:00 +0000 GMT Friday 1 2 3 4 5 1 2 uk marker 1234}
}

func Example_insertDuplicateDraw() {
	db, err := dbConn()
	if err != nil {
		fmt.Printf("Unable to connect to DB: %v", err)
		return
	}
	defer db.Close()

	err = CreateSQLiteTable(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBCreateTbl) {
		fmt.Printf("Unable to create table: %v", err)
		return
	}

	stmt, err := prepInsertDrawStmt(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBPrepareStmt) {
		fmt.Printf("Prepare insert statement: %v ", err)
		return
	}
	defer stmt.Close()

	draws := []Draw{
		{
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
		},
		{
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
		},
	}

	for _, d := range draws {
		_, err = insertDraw(context.TODO(), stmt, d)
		if errors.Is(err, dbutil.ErrDBInsertTbl) {
			fmt.Printf("Error-insert draw: %v", err)
		}
	}

	// Output:
	// Error-insert draw: unable to write to table-UNIQUE constraint failed: euro.draw_no[{2023-01-20 12:00:00 +0000 GMT Friday 1 2 3 4 5 1 2 uk marker 1234}]
}

func Example_selectAllDraw() {

	db, err := dbConn()
	if err != nil {
		fmt.Printf("Unable to connect to DB: %v", err)
		return
	}
	defer db.Close()

	err = CreateSQLiteTable(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBCreateTbl) {
		fmt.Printf("Unable to create table: %v", err)
		return
	}

	stmt, err := prepInsertDrawStmt(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBPrepareStmt) {
		fmt.Printf("Prepare insert statement: %v", err)
		return
	}
	defer stmt.Close()

	draws := []Draw{
		{
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
		},
	}

	for _, d := range draws {
		_, err = insertDraw(context.TODO(), stmt, d)
		if errors.Is(err, dbutil.ErrDBInsertTbl) {
			fmt.Printf("Error-insert draw: %v\n", err)
		}
	}

	rows, err := selectAllDrawRows(context.TODO(), db)
	if err != nil {
		fmt.Printf("Select all draws error: %v", err)
	}

	draw := selectAllDraw(rows)
	for d := range draw {
		fmt.Println(d)
	}

	// Output:
	// {2023-01-20 12:00:00 +0000 GMT Friday 1 2 3 4 5 1 2 uk marker 1234}
}
