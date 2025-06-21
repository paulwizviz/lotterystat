package csvops

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	dateScenarios = []struct {
		name     string
		input    string
		expected struct {
			dt  time.Time
			err error
		}
	}{
		{
			name:  "Single day number, valid format",
			input: "1-Jan-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Single day number, valid format",
			input: "01-Jan-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Invalid year format",
			input: "29-Dec-abc",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDateFmt,
			},
		},
		{
			name:  "Invalid day format",
			input: "a-Dec-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDateFmt,
			},
		},
		{
			name:  "January fullname",
			input: "21-January-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDateFmt,
			},
		},
		{
			name:  "Invalid month",
			input: "29-D-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDateFmt,
			},
		},
		{
			name:  "Numeric",
			input: "29-01-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDateFmt,
			},
		},
		{
			name:  "Slash separator",
			input: "29/1/2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDateFmt,
			},
		},
	}

	dayScenarios = []struct {
		name     string
		input    string
		expected struct {
			dt  time.Time
			err error
		}
	}{
		{
			name:  "Invalid day in Jan",
			input: "32-Dec-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Feb leap year",
			input: "30-Feb-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Feb non leap year",
			input: "29-Feb-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Mar",
			input: "32-Mar-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Apr",
			input: "31-Apr-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in May",
			input: "32-May-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Jun",
			input: "31-Jun-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Jul",
			input: "32-Jul-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Aug",
			input: "32-Aug-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Sep",
			input: "31-Sep-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Oct",
			input: "32-Oct-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Nov",
			input: "31-Nov-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
		{
			name:  "Invalid day in Dec",
			input: "32-Dec-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
		},
	}

	monthScenarios = []struct {
		name     string
		input    string
		expected struct {
			dt  time.Time
			err error
		}
	}{
		{
			name:  "Properly formatted Jan",
			input: "21-Jan-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.January, 21, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Feb non leap year",
			input: "28-Feb-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Feb yeap year",
			input: "29-Feb-2024",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Mar",
			input: "1-Mar-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Apr",
			input: "29-Apr-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted May",
			input: "29-May-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.May, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Jun",
			input: "29-Jun-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.June, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Jul",
			input: "29-Jul-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.July, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Aug",
			input: "29-Aug-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.August, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Sep",
			input: "29-Sep-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.September, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Oct",
			input: "29-Oct-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.October, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Nov",
			input: "29-Nov-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.November, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
		{
			name:  "Properly formatted Dec",
			input: "29-Dec-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.December, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
		},
	}
)

func TestParseDate(t *testing.T) {
	for i, scenario := range dateScenarios {
		t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), func(t *testing.T) {
			actual, err := ParseDate(scenario.input)
			if assert.ErrorIs(t, err, scenario.expected.err) {
				assert.Equal(t, scenario.expected.dt, actual)
			}
		})
	}
}

func TestParseDate_Day(t *testing.T) {
	for i, scenario := range dayScenarios {
		t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), func(t *testing.T) {
			actual, err := ParseDate(scenario.input)
			if assert.ErrorIs(t, err, scenario.expected.err) {
				assert.Equal(t, scenario.expected.dt, actual)
			}
		})
	}
}

func TestParseDate_month(t *testing.T) {
	for i, scenario := range monthScenarios {
		t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), func(t *testing.T) {
			actual, err := ParseDate(scenario.input)
			if assert.ErrorIs(t, err, scenario.expected.err) {
				assert.Equal(t, scenario.expected.dt, actual)
			}
		})
	}
}
