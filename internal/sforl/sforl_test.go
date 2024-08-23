package sforl

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIsValidBet(t *testing.T) {
	testcases := []struct {
		input       string
		expected    bool
		description string
	}{
		{
			input:       "1,10,20,30,47,1",
			expected:    true,
			description: "Valid entry",
		},
		{
			input:       "11,21,25,34,44,10",
			expected:    true,
			description: "Valid entry - mid range",
		},
		{
			input:       "11,21,25,34,1",
			expected:    false,
			description: "Insufficient number of entries",
		},
		{
			input:       "a,21,25,34,1",
			expected:    false,
			description: "Invalid bet item",
		},
		{
			input:       "0,10,20,30,47,1",
			expected:    false,
			description: "1st ball is zero",
		},
		{
			input:       "1,0,20,30,47,1",
			expected:    false,
			description: "2nd ball is zero",
		},
		{
			input:       "1,10,0,30,50,1",
			expected:    false,
			description: "3rd ball is zero",
		},
		{
			input:       "1,10,20,0,50,1",
			expected:    false,
			description: "4th ball is zero",
		},
		{
			input:       "1,10,20,30,0,1",
			expected:    false,
			description: "5th ball is zero",
		},
		{
			input:       "1,10,20,30,47,0",
			expected:    false,
			description: "Life ball is zero",
		},
		{
			input:       "48,10,20,30,47,1",
			expected:    false,
			description: "1st ball is over 47",
		},
		{
			input:       "1,48,20,30,47,1",
			expected:    false,
			description: "2nd ball is over 47",
		},
		{
			input:       "1,10,48,30,47,1",
			expected:    false,
			description: "3rd ball is over 47",
		},
		{
			input:       "1,10,20,48,47,1",
			expected:    false,
			description: "4th ball is over 47",
		},
		{
			input:       "1,10,20,30,48,1",
			expected:    false,
			description: "5th ball is over 47",
		},
		{
			input:       "1,10,20,30,47,11",
			expected:    false,
			description: "Life ball is over 10",
		},
	}
	for i, tc := range testcases {
		actual := IsValidBet(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Description: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}

func Example_twoCombination() {
	set := []uint8{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47,
	}

	combination := [][2]uint8{}
	for i := 0; i < len(set); i++ {
		for j := i + 1; j < len(set); j++ {
			d := [2]uint8{}
			d[0], d[1] = set[i], set[j]
			combination = append(combination, d)
		}
	}

	fmt.Printf("Size of combination: %d", len(combination))

	// Output:
	// Size of combination: 1081
}

func Example_twoMainComboFreq() {
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

	twoMainComboFreq(context.TODO(), stmt1)

	// Output:

}
