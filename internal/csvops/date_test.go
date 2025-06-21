package csvops

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var dateScenarios = []struct {
	name     string
	input    string
	expected struct {
		dt  time.Time
		err error
	}
}{
	// Formats
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
	// January
	{
		name:  "Valid Jan 1",
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
		name:  "Valid Jan 01",
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
		name:  "Valid Jan 31",
		input: "31-Jan-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.January, 31, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Jan 0",
		input: "0-Jan-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid Jan 32",
		input: "32-Jan-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// February
	{
		name:  "Valid Feb 1",
		input: "1-Feb-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Feb 0",
		input: "0-Feb-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Valid Feb 28, non leap year",
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
		name:  "Valid Feb 29, leap year",
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
		name:  "Valid Feb 29, leap year",
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
		name:  "Invalid Feb 30, leap year",
		input: "30-Feb-2024",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// March
	{
		name:  "Valid Mar 1",
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
		name:  "Valid Mar 31",
		input: "31-Mar-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.March, 31, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Mar 0",
		input: "0-Mar-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid Mar 32",
		input: "32-Mar-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// April
	{
		name:  "Valid Apr 1",
		input: "1-Apr-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Valid Apr 30",
		input: "30-Apr-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.April, 30, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Apr 0",
		input: "0-Apr-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid Apr 31",
		input: "31-Apr-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// May
	{
		name:  "Valid May 1",
		input: "1-May-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.May, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Valid May 31",
		input: "31-May-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.May, 31, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid May 0",
		input: "0-May-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid May 32",
		input: "32-May-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// June
	{
		name:  "Invalid Jun 1",
		input: "1-Jun-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.June, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Jun 30",
		input: "30-Jun-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.June, 30, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Jun 0",
		input: "0-Jun-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid Jun 31",
		input: "31-Jun-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// July
	{
		name:  "Valid Jul 1",
		input: "1-Jul-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Valid Jul 31",
		input: "31-Jul-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.July, 31, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Jul 0",
		input: "0-Jul-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid Jul 32",
		input: "32-Jul-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// August
	{
		name:  "Valid Aug 1",
		input: "1-Aug-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Valid Aug 31",
		input: "31-Aug-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.August, 31, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Aug 0",
		input: "0-Aug-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid Aug 32",
		input: "32-Aug-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// September
	{
		name:  "Valid Sep 1",
		input: "1-Sep-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Valid Sep 30",
		input: "30-Sep-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.September, 30, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Sep 0",
		input: "0-Sep-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid Sep 31",
		input: "31-Sep-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// October
	{
		name:  "Valid Oct 1",
		input: "1-Oct-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.October, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Valid Oct 31",
		input: "31-Oct-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.October, 31, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid Oct 0",
		input: "0-Oct-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid Oct 32",
		input: "32-Oct-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	// November
	{
		name:  "Valid day in Nov 1st",
		input: "1-Nov-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.November, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Valid day in Nov 30th",
		input: "30-Nov-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.November, 30, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid day in Nov 0",
		input: "0-Nov-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid day in Nov 31st",
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
		name:  "Valid December 1st",
		input: "1-Dec-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	// December
	{
		name:  "Valid December 1",
		input: "1-Dec-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Valid December 31st",
		input: "31-Dec-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Date(2022, time.December, 31, 0, 0, 0, 0, time.UTC),
			err: nil,
		},
	},
	{
		name:  "Invalid 0 December",
		input: "0-Dec-2022",
		expected: struct {
			dt  time.Time
			err error
		}{
			dt:  time.Time{},
			err: ErrInvalidDaysInMonth,
		},
	},
	{
		name:  "Invalid 32nd December",
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
