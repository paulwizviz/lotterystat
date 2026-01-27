package tball_test

import (
	"context"
	"fmt"
	"time"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/paulwizviz/lotterystat/internal/tball"
)

func Example_insertListDraw() {

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

	d1 := tball.Draw{
		DrawDate:  time.Date(2024, time.August, 28, 0, 0, 0, 0, time.UTC),
		DayOfWeek: time.Date(2024, time.August, 28, 0, 0, 0, 0, time.UTC).Weekday(),
		Ball1:     10,
		Ball2:     20,
		Ball3:     30,
		Ball4:     40,
		Ball5:     50,
		TBall:     11,
		BallSet:   "ball set",
		Machine:   "machine",
		DrawNo:    2,
	}

	err = tball.PersistsDraw(context.TODO(), db, d1)
	if err != nil {
		fmt.Println(err)
	}

	results, err := tball.ListAllDraws(context.TODO(), db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output:
	// [{2024-08-28 00:00:00 +0000 UTC Wednesday 1 2 3 4 5 1 ball set machine 1} {2024-08-28 00:00:00 +0000 UTC Wednesday 10 20 30 40 50 11 ball set machine 2}]
}
