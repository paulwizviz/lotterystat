package tball_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/paulwizviz/lotterystat/internal/sqlops"
	"github.com/paulwizviz/lotterystat/internal/tball"
)

func TestCalculateBallFreq(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()
	err = sqlops.CreateTables(ctx, db, tball.CreateTableFn)
	if err != nil {
		t.Fatal(err)
	}

	// Insert some draws
	// Ball 1: appears twice
	// Ball 2: appears once
	// Ball 38: appears once
	draws := []tball.Draw{
		{
			DrawDate: time.Date(2024, time.August, 28, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 2, Ball3: 3, Ball4: 4, Ball5: 5, TBall: 1,
			DrawNo: 1,
		},
		{
			DrawDate: time.Date(2024, time.August, 29, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 10, Ball3: 20, Ball4: 30, Ball5: 38, TBall: 2,
			DrawNo: 2,
		},
	}

	for _, d := range draws {
		err = tball.PersistsDraw(ctx, db, d)
		if err != nil {
			t.Fatal(err)
		}
	}

	freqs, err := tball.CalculateBallFreq(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	if len(freqs) != 38 {
		t.Fatalf("expected 38 frequencies, got %d", len(freqs))
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
		38: 1,
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

	// Test with canceled context
	canceledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = tball.CalculateBallFreq(canceledCtx, db)
	if err == nil {
		t.Error("expected error for canceled context, got nil")
	}
}

func TestCalculateTBallFreq(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()
	err = sqlops.CreateTables(ctx, db, tball.CreateTableFn)
	if err != nil {
		t.Fatal(err)
	}

	// Insert some draws
	// TBall 1: appears once
	// TBall 13: appears twice
	draws := []tball.Draw{
		{
			DrawDate: time.Date(2024, time.August, 28, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 2, Ball3: 3, Ball4: 4, Ball5: 5, TBall: 1,
			DrawNo: 1,
		},
		{
			DrawDate: time.Date(2024, time.August, 29, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 10, Ball3: 20, Ball4: 30, Ball5: 38, TBall: 13,
			DrawNo: 2,
		},
		{
			DrawDate: time.Date(2024, time.August, 30, 0, 0, 0, 0, time.UTC),
			Ball1:    2, Ball2: 11, Ball3: 21, Ball4: 31, Ball5: 37, TBall: 13,
			DrawNo: 3,
		},
	}

	for _, d := range draws {
		err = tball.PersistsDraw(ctx, db, d)
		if err != nil {
			t.Fatal(err)
		}
	}

	freqs, err := tball.CalculateTBallFreq(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	if len(freqs) != 13 {
		t.Fatalf("expected 13 frequencies, got %d", len(freqs))
	}

	// Check specific frequencies
	expectedFreqs := map[uint]uint{
		1:  1,
		13: 2,
	}

	for _, tbf := range freqs {
		if expected, ok := expectedFreqs[tbf.TBall]; ok {
			if tbf.Frequency != expected {
				t.Errorf("tball %d: expected frequency %d, got %d", tbf.TBall, expected, tbf.Frequency)
			}
		} else {
			if tbf.Frequency != 0 {
				t.Errorf("tball %d: expected frequency 0, got %d", tbf.TBall, tbf.Frequency)
			}
		}
	}

	// Test with canceled context
	canceledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = tball.CalculateTBallFreq(canceledCtx, db)
	if err == nil {
		t.Error("expected error for canceled context, got nil")
	}
}

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
