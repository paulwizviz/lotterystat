package sflife_test

import (
	"context"
	"fmt"
	"time"

	"github.com/paulwizviz/lotterystat/internal/sflife"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

func Example_insertListDraw() {

	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = sqlops.CreateTables(context.TODO(), db, sflife.CreateTableFn)
	if err != nil {
		fmt.Println(err)
	}

	d := sflife.Draw{
		DrawDate:  time.Date(2026, time.February, 19, 0, 0, 0, 0, time.UTC),
		DayOfWeek: time.Thursday,
		Ball1:     5,
		Ball2:     9,
		Ball3:     13,
		Ball4:     34,
		Ball5:     45,
		LBall:     8,
		BallSet:   "SFL3",
		Machine:   "Excalibur6",
		DrawNo:    724,
	}

	err = sflife.PersistsDraw(context.TODO(), db, d)
	if err != nil {
		fmt.Println(err)
	}

	results, err := sflife.ListAllDraws(context.TODO(), db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output:
	// [{2026-02-19 00:00:00 +0000 UTC Thursday 5 9 13 34 45 8 SFL3 Excalibur6 724}]
}
