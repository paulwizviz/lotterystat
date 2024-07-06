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
