package sforl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func dbConn() (*sql.DB, error) {
	return sql.Open("sqlite3", ":memory:")
}

func listSQLiteTables(db *sql.DB) ([]string, error) {
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

func Example_listSQLiteTable() {
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
	tbls, err := listSQLiteTables(db)
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

func Example_freqBallSQLiteSQL() {
	s := freqBallSQLiteSQL()
	fmt.Println(s)

	// Output:
	// SELECT COUNT(*) FROM set_for_life WHERE ball1=? OR ball2=? OR ball3=? OR ball4=? OR ball5=?;
}

func Example_ballCount() {
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
	stmt1, err := prepInsertSQLiteDrawStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Insert table error: %v", err)
		return
	}
	d := Draw{
		DrawDate:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local),
		DayOfWeek: time.Monday,
		Ball1:     1,
		Ball2:     2,
		Ball3:     3,
		Ball4:     4,
		Ball5:     5,
		LifeBall:  1,
		BallSet:   "abc",
		Machine:   "efg",
		DrawNo:    1,
	}
	insertSQLiteDraw(context.TODO(), stmt1, d)
	stmt2, err := prepSQLiteBallCountStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Ball count statement error: %v", err)
	}
	result, err := ballCount(context.TODO(), stmt2, 1)
	if err != nil {
		fmt.Printf("Ball count error: %v", err)
	}
	fmt.Println(result)

	// Output:
	// 1
}
