package lotto_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/paulwizviz/lotterystat/internal/lotto"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

func TestCalculateBallFreq(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()
	err = sqlops.CreateTables(ctx, db, lotto.CreateTableFn)
	if err != nil {
		t.Fatal(err)
	}

	// Insert some draws
	draws := []lotto.Draw{
		{
			DrawDate: time.Date(2026, time.February, 18, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 2, Ball3: 3, Ball4: 4, Ball5: 5, Ball6: 6, BonusBall: 7,
			DrawNo: 1,
		},
		{
			DrawDate: time.Date(2026, time.February, 19, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 10, Ball3: 20, Ball4: 30, Ball5: 40, Ball6: 59, BonusBall: 50,
			DrawNo: 2,
		},
	}

	for _, d := range draws {
		err = lotto.PersistsDraw(ctx, db, d)
		if err != nil {
			t.Fatal(err)
		}
	}

	freqs, err := lotto.CalculateBallFreq(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	if len(freqs) != 59 {
		t.Fatalf("expected 59 frequencies, got %d", len(freqs))
	}

	// Check specific frequencies
	expectedFreqs := map[uint]uint{
		1:  2,
		2:  1,
		3:  1,
		4:  1,
		5:  1,
		6:  1,
		10: 1,
		20: 1,
		30: 1,
		40: 1,
		59: 1,
	}

	for _, bf := range freqs {
		if expected, ok := expectedFreqs[bf.Ball]; ok {
			if bf.Frequency != expected {
				t.Errorf("ball %d: expected frequency %d, got %d", bf.Ball, expected, bf.Frequency)
			}
		} else {
			if bf.Frequency != 0 {
				t.Errorf("ball %d: expected frequency 0, got %d", bf.Ball, bf.Frequency)
			}
		}
	}
}

func TestCalculateBonusFreq(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()
	err = sqlops.CreateTables(ctx, db, lotto.CreateTableFn)
	if err != nil {
		t.Fatal(err)
	}

	// Insert some draws
	draws := []lotto.Draw{
		{
			DrawDate: time.Date(2026, time.February, 18, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 2, Ball3: 3, Ball4: 4, Ball5: 5, Ball6: 6, BonusBall: 1,
			DrawNo: 1,
		},
		{
			DrawDate: time.Date(2026, time.February, 19, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 10, Ball3: 20, Ball4: 30, Ball5: 40, Ball6: 59, BonusBall: 59,
			DrawNo: 2,
		},
	}

	for _, d := range draws {
		err = lotto.PersistsDraw(ctx, db, d)
		if err != nil {
			t.Fatal(err)
		}
	}

	freqs, err := lotto.CalculateBonusFreq(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	if len(freqs) != 59 {
		t.Fatalf("expected 59 frequencies, got %d", len(freqs))
	}

	// Check specific frequencies
	expectedFreqs := map[uint]uint{
		1:  1,
		59: 1,
	}

	for _, bf := range freqs {
		if expected, ok := expectedFreqs[bf.Ball]; ok {
			if bf.Frequency != expected {
				t.Errorf("bonus ball %d: expected frequency %d, got %d", bf.Ball, expected, bf.Frequency)
			}
		} else {
			if bf.Frequency != 0 {
				t.Errorf("bonus ball %d: expected frequency 0, got %d", bf.Ball, bf.Frequency)
			}
		}
	}
}

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
