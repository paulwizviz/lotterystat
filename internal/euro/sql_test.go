package euro_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/paulwizviz/lotterystat/internal/euro"
	"github.com/paulwizviz/lotterystat/internal/sqlops"
)

func TestCalculateBallFreq(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()
	err = sqlops.CreateTables(ctx, db, euro.CreateTableFn)
	if err != nil {
		t.Fatal(err)
	}

	// Insert some draws
	draws := []euro.Draw{
		{
			DrawDate: time.Date(2026, time.February, 20, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 2, Ball3: 3, Ball4: 4, Ball5: 5, Star1: 1, Star2: 2,
			DrawNo: 1,
		},
		{
			DrawDate: time.Date(2026, time.February, 21, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 10, Ball3: 20, Ball4: 30, Ball5: 50, Star1: 1, Star2: 12,
			DrawNo: 2,
		},
	}

	for _, d := range draws {
		err = euro.PersistsDraw(ctx, db, d)
		if err != nil {
			t.Fatal(err)
		}
	}

	freqs, err := euro.CalculateBallFreq(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	if len(freqs) != 50 {
		t.Fatalf("expected 50 frequencies, got %d", len(freqs))
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
		50: 1,
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

func TestCalculateStarFreq(t *testing.T) {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()
	err = sqlops.CreateTables(ctx, db, euro.CreateTableFn)
	if err != nil {
		t.Fatal(err)
	}

	// Insert some draws
	draws := []euro.Draw{
		{
			DrawDate: time.Date(2026, time.February, 20, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 2, Ball3: 3, Ball4: 4, Ball5: 5, Star1: 1, Star2: 2,
			DrawNo: 1,
		},
		{
			DrawDate: time.Date(2026, time.February, 21, 0, 0, 0, 0, time.UTC),
			Ball1:    1, Ball2: 10, Ball3: 20, Ball4: 30, Ball5: 50, Star1: 1, Star2: 12,
			DrawNo: 2,
		},
	}

	for _, d := range draws {
		err = euro.PersistsDraw(ctx, db, d)
		if err != nil {
			t.Fatal(err)
		}
	}

	freqs, err := euro.CalculateStarFreq(ctx, db)
	if err != nil {
		t.Fatal(err)
	}

	if len(freqs) != 12 {
		t.Fatalf("expected 12 frequencies, got %d", len(freqs))
	}

	// Check specific frequencies
	expectedFreqs := map[uint]uint{
		1:  2,
		2:  1,
		12: 1,
	}

	for _, sf := range freqs {
		if expected, ok := expectedFreqs[sf.Star]; ok {
			if sf.Frequency != expected {
				t.Errorf("star %d: expected frequency %d, got %d", sf.Star, expected, sf.Frequency)
			}
		} else {
			if sf.Frequency != 0 {
				t.Errorf("star %d: expected frequency 0, got %d", sf.Star, sf.Frequency)
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
