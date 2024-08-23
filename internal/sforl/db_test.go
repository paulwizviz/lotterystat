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
	s := countBallSQL()
	fmt.Println(s)

	// Output:
	// SELECT COUNT(*) FROM set_for_life WHERE ball1=? OR ball2=? OR ball3=? OR ball4=? OR ball5=?;
}

func Example_countChoice() {
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
	stmt1, err := prepInsertDrawStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Insert table error: %v", err)
		return
	}
	defer stmt1.Close()

	d := Draw{
		DrawDate:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local),
		DayOfWeek: time.Monday,
		Ball1:     1,
		Ball2:     2,
		Ball3:     3,
		Ball4:     4,
		Ball5:     5,
		LifeBall:  9,
		BallSet:   "abc",
		Machine:   "efg",
		DrawNo:    1,
	}
	insertDraw(context.TODO(), stmt1, d)

	stmt2, err := prepCountBallStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Ball count statement error: %v", err)
	}
	defer stmt2.Close()

	stmt3, err := prepCountLuckyStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Lucky count statement error: %v", err)
	}
	defer stmt3.Close()

	ballCount, err := countChoice(context.TODO(), stmt2, 1)
	if err != nil {
		fmt.Printf("Ball count error: %v", err)
	}

	luckyCount, err := countChoice(context.TODO(), stmt3, 9)
	if err != nil {
		fmt.Printf("Ball count error: %v", err)
	}

	fmt.Printf("Number of Ball numbered 1: %d\n", ballCount)
	fmt.Printf("Number of Lucky star numbered 9: %d\n", luckyCount)

	// Output:
	// Number of Ball numbered 1: 1
	// Number of Lucky star numbered 9: 1
}

func Example_countTwoMain() {
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
	stmt, err := prepInsertDrawStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Insert table error: %v", err)
		return
	}
	defer stmt.Close()

	d1 := Draw{
		DrawDate:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local),
		DayOfWeek: time.Monday,
		Ball1:     1,
		Ball2:     11,
		Ball3:     22,
		Ball4:     24,
		Ball5:     30,
		LifeBall:  1,
		BallSet:   "abc",
		Machine:   "efg",
		DrawNo:    1,
	}
	insertDraw(context.TODO(), stmt, d1)

	d2 := Draw{
		DrawDate:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.Local),
		DayOfWeek: time.Monday,
		Ball1:     1,
		Ball2:     10,
		Ball3:     20,
		Ball4:     30,
		Ball5:     42,
		LifeBall:  5,
		BallSet:   "hij",
		Machine:   "klm",
		DrawNo:    2,
	}
	insertDraw(context.TODO(), stmt, d2)

	stmt1, err := prepTwoMainStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Two main statement error: %v", err)
	}
	defer stmt1.Close()

	count, err := countTwoMain(context.TODO(), stmt1, 1, 30)
	if err != nil {
		fmt.Printf("Two main count error: %v", err)
	}

	fmt.Printf("Combination of 1,30 count: %d", count)

	// Output:
	// Combination of 1,30 count: 2
}
