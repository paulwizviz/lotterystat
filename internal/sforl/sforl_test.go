package sforl

import (
	"testing"
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
