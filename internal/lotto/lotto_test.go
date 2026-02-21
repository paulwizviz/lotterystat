package lotto

import (
	"testing"
)

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
			input:       "60",
			expected:    false,
			description: "One invalid ball",
		},
		{
			input:       "1,2,3,4,5,6",
			expected:    true,
			description: "Six valid balls",
		},
		{
			input:       "1,2,3,4,5,60",
			expected:    false,
			description: "Six balls. Sixth invalid",
		},
	}
	for i, tc := range testcases {
		actual := IsValidBall(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Description: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}

func TestIsValidBonus(t *testing.T) {
	testcases := []struct {
		input       string
		expected    bool
		description string
	}{
		{
			input:       "1",
			expected:    true,
			description: "Valid bonus ball",
		},
		{
			input:       "59",
			expected:    true,
			description: "Valid bonus ball",
		},
		{
			input:       "60",
			expected:    false,
			description: "Invalid bonus ball",
		},
	}
	for i, tc := range testcases {
		actual := IsValidBonus(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Description: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}
