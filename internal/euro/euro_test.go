package euro

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
			input:       "51",
			expected:    false,
			description: "One invalid ball",
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
			t.Fatalf("Case: %d Description: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
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
			input:       "1,12",
			expected:    true,
			description: "Two valid stars",
		},
	}
	for i, tc := range testcases {
		actual := IsValidStars(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Description: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}
