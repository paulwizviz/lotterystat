package euro

import (
	"reflect"
	"testing"
)

func TestIsValidBet(t *testing.T) {
	testcases := []struct {
		input       string
		expected    bool
		description string
	}{
		{
			input:       "1,10,20,30,50,1,12",
			expected:    true,
			description: "Valid entry",
		},
		{
			input:       "11,21,25,34,44,1,12",
			expected:    true,
			description: "Valid entry - mid range",
		},
		{
			input:       "11,21,25,34,1,12",
			expected:    false,
			description: "Insufficient number of entries",
		},
		{
			input:       "a,21,25,34,1,12",
			expected:    false,
			description: "Invalid bet item",
		},
		{
			input:       "0,10,20,30,50,1,12",
			expected:    false,
			description: "1st ball is zero",
		},
		{
			input:       "1,0,20,30,50,1,12",
			expected:    false,
			description: "2nd ball is zero",
		},
		{
			input:       "1,10,0,30,50,1,12",
			expected:    false,
			description: "3rd ball is zero",
		},
		{
			input:       "1,10,20,0,50,1,12",
			expected:    false,
			description: "4th ball is zero",
		},
		{
			input:       "1,10,20,30,0,1,12",
			expected:    false,
			description: "5th ball is zero",
		},
		{
			input:       "1,10,20,30,50,0,12",
			expected:    false,
			description: "LS1 is zero",
		},
		{
			input:       "1,10,20,30,50,1,0",
			expected:    false,
			description: "LS2 is zero",
		},
		{
			input:       "51,10,20,30,50,1,12",
			expected:    false,
			description: "1st ball is over 50",
		},
		{
			input:       "1,51,20,30,50,1,12",
			expected:    false,
			description: "2nd ball is over 50",
		},
		{
			input:       "1,10,51,30,50,1,12",
			expected:    false,
			description: "3rd ball is over 50",
		},
		{
			input:       "1,10,20,51,50,1,12",
			expected:    false,
			description: "4th ball is over 50",
		},
		{
			input:       "1,10,20,30,51,1,12",
			expected:    false,
			description: "5th ball is over 50",
		},
		{
			input:       "1,10,20,30,50,13,12",
			expected:    false,
			description: "LS1 is over 12",
		},
		{
			input:       "1,10,20,30,50,1,13",
			expected:    false,
			description: "LS2 is over 12",
		},
	}
	for i, tc := range testcases {
		actual := IsValidBet(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Description: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}

func TestProcessBetArg(t *testing.T) {
	testcases := []struct {
		input       string
		expected    Bet
		description string
	}{
		{
			input: "1,10,20,30,45,1,12",
			expected: Bet{
				Ball1: 1,
				Ball2: 10,
				Ball3: 20,
				Ball4: 30,
				Ball5: 45,
				LS1:   1,
				LS2:   12,
			},
			description: "Valid argument",
		},
		{
			input: "10,1,25,20,45,10,1",
			expected: Bet{
				Ball1: 1,
				Ball2: 10,
				Ball3: 20,
				Ball4: 25,
				Ball5: 45,
				LS1:   1,
				LS2:   10,
			},
			description: "Invalid arrangement",
		},
	}
	for i, tc := range testcases {
		actual, _ := ProcessBetArg(tc.input)
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Fatalf("Case: %d Description: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}

func TestIsValidBalls(t *testing.T) {
	testcases := []struct {
		input       string
		expected    bool
		description string
	}{
		{
			input:       "1",
			expected:    true,
			description: "One valid ball",
		},
		{
			input:       "51",
			expected:    false,
			description: "One invalid ball",
		},
		{
			input:       "1,2",
			expected:    true,
			description: "Two valid balls",
		},
		{
			input:       "51,2",
			expected:    false,
			description: "Two valid balls. First invalid",
		},
		{
			input:       "1,51",
			expected:    false,
			description: "Two valid balls. Second invalid",
		},
		{
			input:       "1,2,3",
			expected:    true,
			description: "Three valid balls",
		},
		{
			input:       "51,2,3",
			expected:    false,
			description: "Three balls. First invalid",
		},
		{
			input:       "1,51,3",
			expected:    false,
			description: "Three balls. Second invalid",
		},
		{
			input:       "1,2,51",
			expected:    false,
			description: "Three balls. First invalid",
		},
		{
			input:       "1,2,3,4",
			expected:    true,
			description: "Four valid balls",
		},
		{
			input:       "1,2,3,51",
			expected:    false,
			description: "Four balls. Fourth invalid",
		},
		{
			input:       "1,2,3,4,5",
			expected:    true,
			description: "Five valid balls",
		},
		{
			input:       "1,2,3,4,51",
			expected:    false,
			description: "Five balls. Fifth invalid",
		},
	}
	for i, tc := range testcases {
		actual := IsValidBall(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Descrption: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}

func TestIsValidStars(t *testing.T) {
	testcases := []struct {
		input       string
		expected    bool
		description string
	}{
		{
			input:       "1",
			expected:    true,
			description: "One valid star",
		},
		{
			input:       "13",
			expected:    false,
			description: "One invalid star",
		},
		{
			input:       "1,1",
			expected:    true,
			description: "Two valid stars",
		},
		{
			input:       "13,1",
			expected:    false,
			description: "Two stars. First invalid.",
		},
		{
			input:       "1,13",
			expected:    false,
			description: "Two stars. Second invalid.",
		},
	}
	for i, tc := range testcases {
		actual := IsValidStars(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Descrption: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}
