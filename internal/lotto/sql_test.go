package lotto_test

import (
	"context"
	"fmt"
	"time"

	"github.com/paulwizviz/lotterystat/internal/lotto"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

func Example_insertListDraw() {

	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = sqlops.CreateTables(context.TODO(), db, lotto.CreateTableFn)
	if err != nil {
		fmt.Println(err)
	}

	d := lotto.Draw{
		DrawDate:  time.Date(2026, time.February, 18, 0, 0, 0, 0, time.UTC),
		DayOfWeek: time.Wednesday,
		Ball1:     1,
		Ball2:     11,
		Ball3:     12,
		Ball4:     13,
		Ball5:     18,
		Ball6:     49,
		BonusBall: 33,
		BallSet:   "L10",
		Machine:   "Lotto4",
		DrawNo:    3147,
	}

	err = lotto.PersistsDraw(context.TODO(), db, d)
	if err != nil {
		fmt.Println(err)
	}

	results, err := lotto.ListAllDraws(context.TODO(), db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output:
	// [{2026-02-18 00:00:00 +0000 UTC Wednesday 1 11 12 13 18 49 33 L10 Lotto4 3147}]
}
