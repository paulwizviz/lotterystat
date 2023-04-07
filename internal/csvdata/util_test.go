package csvdata

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseProperDate(t *testing.T) {

	tcases := []struct {
		input       string
		expected    time.Time
		description string
	}{
		{
			input:       "21-Jan-2022",
			expected:    time.Date(2022, time.January, 21, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "28-Feb-2022",
			expected:    time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Feb-2024",
			expected:    time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "1-Mar-2022",
			expected:    time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Apr-2022",
			expected:    time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-May-2022",
			expected:    time.Date(2022, time.May, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Jun-2022",
			expected:    time.Date(2022, time.June, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Jul-2022",
			expected:    time.Date(2022, time.July, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Aug-2022",
			expected:    time.Date(2022, time.August, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Sep-2022",
			expected:    time.Date(2022, time.September, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Oct-2022",
			expected:    time.Date(2022, time.October, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Nov-2022",
			expected:    time.Date(2022, time.November, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
		{
			input:       "29-Dec-2022",
			expected:    time.Date(2022, time.December, 29, 0, 0, 0, 0, time.UTC),
			description: "Properly formatted date time",
		},
	}

	for i, tc := range tcases {
		actual, err := parseDateTime(tc.input)
		if assert.Nil(t, err) {
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}

func TestParseInvalidDate(t *testing.T) {
	testcases := []struct {
		input       string
		expected    time.Time
		expectedErr error
		description string
	}{
		{
			input:       "a-Jan-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidNumFmt,
			description: "Invalid day format",
		},
		{
			input:       "32-Jan-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Jan",
		},
		{
			input:       "29-Feb-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Feb",
		},
		{
			input:       "30-Feb-2024",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Feb leap year",
		},
		{
			input:       "32-Mar-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Mar",
		},
		{
			input:       "31-Apr-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Apr",
		},
		{
			input:       "32-May-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Apr",
		},
		{
			input:       "31-Jun-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Jun",
		},
		{
			input:       "32-Jul-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Jul",
		},
		{
			input:       "32-Aug-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Aug",
		},
		{
			input:       "31-Sep-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Sep",
		},
		{
			input:       "32-Oct-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Oct",
		},
		{
			input:       "31-Nov-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Nov",
		},
		{
			input:       "32-Dec-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidDaysInMonth,
			description: "Day out of range in Dec",
		},
		{
			input:       "1-January-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Jan",
		},
		{
			input:       "1-February-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Feb",
		},
		{
			input:       "1-March-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Mar",
		},
		{
			input:       "1-April-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Apr",
		},
		{
			input:       "1-June-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month June",
		},
		{
			input:       "1-July-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month July",
		},
		{
			input:       "1-August-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Aug",
		},
		{
			input:       "1-September-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Sep",
		},
		{
			input:       "1-October-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Oct",
		},
		{
			input:       "1-November-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Jan",
		},
		{
			input:       "1-December-2022",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidMonth,
			description: "Invalid month Dec",
		},
		{
			input:       "1-Jan-a",
			expected:    time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrInvalidNumFmt,
			description: "Invalid month Jan",
		},
	}

	for i, tc := range testcases {
		actual, actualErr := parseDateTime(tc.input)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		assert.True(t, errors.Is(actualErr, tc.expectedErr))
	}
}
