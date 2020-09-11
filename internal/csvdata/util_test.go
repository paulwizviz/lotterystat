package csvdata

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDateTime(t *testing.T) {

	tcases := []struct {
		input       string
		expected    time.Time
		description string
	}{
		{
			input: "1-Jan-2022",
			expected: func() time.Time {
				return time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
			}(),
			description: "Passing operations",
		},
		{
			input: "32-Jai-2022",
			expected: func() time.Time {
				return time.Time{}
			}(),
			description: "Incorrect month",
		},
		{
			input: "32-Jan-2022",
			expected: func() time.Time {
				return time.Time{}
			}(),
			description: "32 day in January",
		},
		{
			input: "0-Jan-2022",
			expected: func() time.Time {
				return time.Time{}
			}(),
			description: "0 day in January",
		},
		{
			input: "a-Jan-2022",
			expected: func() time.Time {
				return time.Time{}
			}(),
			description: "Wrong day type",
		},
		{
			input: "1-Jan-202a",
			expected: func() time.Time {
				return time.Time{}
			}(),
			description: "Wrong year type",
		},
	}

	for i, tc := range tcases {
		actual, err := parseDateTime(tc.input)

		var e *strconv.NumError
		if errors.Is(err, errInvalidMonth) {
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Test: %d Description: %s", i, tc.description))
		} else if errors.Is(err, errInvalidDaysInMonth) {
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Test: %d Description: %s", i, tc.description))
		} else if errors.As(err, &e) {
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Test: %d Description: %s", i, tc.description))
		} else {
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Test: %d Description: %s", i, tc.description))
		}
	}

}
