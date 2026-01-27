package tball_test

import (
	"context"
	"fmt"
	"time"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/paulwizviz/lotterystat/internal/tball"
)

func Example_insertDraw() {

	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = sqlops.CreateTables(context.TODO(), db, tball.CreateTableFn)
	if err != nil {
		fmt.Println(err)
	}

	d := tball.Draw{
		DrawDate:  time.Date(2024, time.August, 28, 0, 0, 0, 0, time.UTC),
		DayOfWeek: time.Date(2024, time.August, 28, 0, 0, 0, 0, time.UTC).Weekday(),
		Ball1:     1,
		Ball2:     2,
		Ball3:     3,
		Ball4:     4,
		Ball5:     5,
		TBall:     1,
		BallSet:   "ball set",
		Machine:   "machine",
		DrawNo:    1,
	}

	err = tball.PersistsDraw(context.TODO(), db, d)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM tball")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var drawDate string
		var dayOfWeek time.Weekday
		var ball1 int
		var ball2 int
		var ball3 int
		var ball4 int
		var ball5 int
		var tBall int
		var ballSet string
		var machine string
		var drawNo uint64
		err := rows.Scan(&drawDate, &dayOfWeek, &ball1, &ball2, &ball3, &ball4, &ball5, &tBall, &ballSet, &machine, &drawNo)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(drawDate, dayOfWeek, ball1, ball2, ball3, ball4, ball5, tBall, machine, drawNo)
	}

	// Output:
}
