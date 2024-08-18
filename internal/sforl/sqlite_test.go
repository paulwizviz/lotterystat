package sforl

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
	// set_for_life
}

func Example_listAllDraw() {
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

	stmt, err := prepInsertSQLiteDrawStmt(context.TODO(), db)
	if errors.Is(err, dbutil.ErrDBPrepareStmt) {
		fmt.Printf("Prepare insert statement: %v ", err)
		return
	}
	defer stmt.Close()

	d := Draw{
		DrawDate:  time.Date(2003, time.January, 10, 12, 00, 00, 0, time.UTC),
		DayOfWeek: time.Date(2003, time.January, 10, 12, 00, 00, 0, time.UTC).Weekday(),
		Ball1:     1,
		Ball2:     2,
		Ball3:     3,
		Ball4:     4,
		Ball5:     5,
		LifeBall:  1,
		BallSet:   "ball set",
		Machine:   "machine",
		DrawNo:    1234,
	}
	_, err = insertSQLiteDraw(context.TODO(), stmt, d)

	draws, err := listSQLiteAllDraw(context.TODO(), db)
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
	// draw: {2003-01-10 12:00:00 +0000 GMT Friday 1 2 3 4 5 1 ball set machine 1234}
}
