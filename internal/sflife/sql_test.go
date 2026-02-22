package sflife_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/paulwizviz/lotterystat/internal/sflife"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

func TestCalculateBallFreq(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()
	err = sqlops.CreateTables(ctx, db, sflife.CreateTableFn)
	if err != nil {
		t.Fatal(err)
	}

	// Insert some draws
	draws := []sflife.Draw{
		{
			DrawDate: time.Date(2026, time.February, 19, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 2, Ball3: 3, Ball4: 4, Ball5: 5, LBall: 1,
			DrawNo: 1,
		},
		{
			DrawDate: time.Date(2026, time.February, 20, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 10, Ball3: 20, Ball4: 30, Ball5: 47, LBall: 2,
			DrawNo: 2,
		},
	}

	for _, d := range draws {
		err = sflife.PersistsDraw(ctx, db, d)
		if err != nil {
			t.Fatal(err)
		}
	}

	freqs, err := sflife.CalculateBallFreq(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	if len(freqs) != 47 {
		t.Fatalf("expected 47 frequencies, got %d", len(freqs))
	}

	// Check specific frequencies
	expectedFreqs := map[uint]uint{
		1:  2,
		2:  1,
		3:  1,
		4:  1,
		5:  1,
		10: 1,
		20: 1,
		30: 1,
		47: 1,
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

func TestCalculateLBallFreq(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()
	err = sqlops.CreateTables(ctx, db, sflife.CreateTableFn)
	if err != nil {
		t.Fatal(err)
	}

	// Insert some draws
	draws := []sflife.Draw{
		{
			DrawDate: time.Date(2026, time.February, 19, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 2, Ball3: 3, Ball4: 4, Ball5: 5, LBall: 1,
			DrawNo: 1,
		},
		{
			DrawDate: time.Date(2026, time.February, 20, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 10, Ball3: 20, Ball4: 30, Ball5: 47, LBall: 10,
			DrawNo: 2,
		},
	}

	for _, d := range draws {
		err = sflife.PersistsDraw(ctx, db, d)
		if err != nil {
			t.Fatal(err)
		}
	}

	freqs, err := sflife.CalculateLBallFreq(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	if len(freqs) != 10 {
		t.Fatalf("expected 10 frequencies, got %d", len(freqs))
	}

	// Check specific frequencies
	expectedFreqs := map[uint]uint{
		1:  1,
		10: 1,
	}

	for _, lbf := range freqs {
		if expected, ok := expectedFreqs[lbf.LBall]; ok {
			if lbf.Frequency != expected {
				t.Errorf("lball %d: expected frequency %d, got %d", lbf.LBall, expected, lbf.Frequency)
			}
		} else {
			if lbf.Frequency != 0 {
				t.Errorf("lball %d: expected frequency 0, got %d", lbf.LBall, lbf.Frequency)
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
