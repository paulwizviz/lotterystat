package sflife

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
			input:       "48",
			expected:    false,
			description: "One invalid ball",
		},
		{
			input:       "1,2,3,4,5",
			expected:    true,
			description: "Five valid balls",
		},
		{
			input:       "1,2,3,4,48",
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

func TestIsValidLifeBall(t *testing.T) {
	testcases := []struct {
		input       string
		expected    bool
		description string
	}{
		{
			input:       "1",
			expected:    true,
			description: "Valid life ball",
		},
		{
			input:       "10",
			expected:    true,
			description: "Valid life ball",
		},
		{
			input:       "11",
			expected:    false,
			description: "Invalid life ball",
		},
	}
	for i, tc := range testcases {
		actual := IsValidLifeBall(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Description: %s Expected: %v Actual: %v", i, tc.description, tc.expected, actual)
		}
	}
}
