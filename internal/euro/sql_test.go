package euro_test

import (
	"context"
	"fmt"
	"time"

	"github.com/paulwizviz/lotterystat/internal/euro"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

func Example_insertListDraw() {

	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = sqlops.CreateTables(context.TODO(), db, euro.CreateTableFn)
	if err != nil {
		fmt.Println(err)
	}

	d := euro.Draw{
		DrawDate:  time.Date(2026, time.February, 20, 0, 0, 0, 0, time.UTC),
		DayOfWeek: time.Friday,
		Ball1:     13,
		Ball2:     24,
		Ball3:     28,
		Ball4:     33,
		Ball5:     35,
		Star1:     5,
		Star2:     9,
		UKMaker:   "ZDTF34718",
		EUMaker:   "",
		BallSet:   "21",
		Machine:   "13",
		DrawNo:    1922,
	}

	err = euro.PersistsDraw(context.TODO(), db, d)
	if err != nil {
		fmt.Println(err)
	}

	results, err := euro.ListAllDraws(context.TODO(), db)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(results)

	// Output:
	// [{2026-02-20 00:00:00 +0000 UTC Friday 13 24 28 33 35 5 9 ZDTF34718  21 13 1922}]
}
