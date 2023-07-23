package draw

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDateTime(t *testing.T) {

	tcases := []struct {
		input    string
		expected struct {
			dt  time.Time
			err error
		}
		description string
	}{
		{
			input: "21-Jan-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.January, 21, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Jan",
		},
		{
			input: "28-Feb-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.February, 28, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Feb non leap year",
		},
		{
			input: "29-Feb-2024",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Feb yeap year",
		},
		{
			input: "1-Mar-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Mar",
		},
		{
			input: "29-Apr-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Apr",
		},
		{
			input: "29-May-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.May, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted May",
		},
		{
			input: "29-Jun-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.June, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Jun",
		},
		{
			input: "29-Jul-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.July, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Jul",
		},
		{
			input: "29-Aug-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.August, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Aug",
		},
		{
			input: "29-Sep-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.September, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Sep",
		},
		{
			input: "29-Oct-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.October, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Oct",
		},
		{
			input: "29-Nov-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.November, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Nov",
		},
		{
			input: "29-Dec-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Date(2022, time.December, 29, 0, 0, 0, 0, time.UTC),
				err: nil,
			},
			description: "Properly formatted Dec",
		},
		{
			input: "a-Dec-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDayFmt,
			},
			description: "Invalid day format",
		},
		{
			input: "29-D-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidMonth,
			},
			description: "Invalid month format",
		},
		{
			input: "29-Dec-abc",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidYearFmt,
			},
			description: "Invalid year format",
		},
		{
			input: "32-Dec-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Jan",
		},
		{
			input: "30-Feb-2024",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Feb leap year",
		},
		{
			input: "29-Feb-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Feb non leap year",
		},
		{
			input: "32-Mar-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Mar",
		},
		{
			input: "31-Apr-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Apr",
		},
		{
			input: "32-May-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in May",
		},
		{
			input: "31-Jun-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Jun",
		},
		{
			input: "32-Jul-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Ju",
		},
		{
			input: "32-Aug-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Aug",
		},
		{
			input: "31-Sep-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Sep",
		},
		{
			input: "32-Oct-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Oct",
		},
		{
			input: "31-Nov-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Nov",
		},
		{
			input: "32-Dec-2022",
			expected: struct {
				dt  time.Time
				err error
			}{
				dt:  time.Time{},
				err: ErrInvalidDaysInMonth,
			},
			description: "Invalid day in Dec",
		},
	}

	for i, tc := range tcases {
		actual, err := parseDateTime(tc.input)
		if assert.True(t, errors.Is(err, tc.expected.err), fmt.Sprintf("Case: %d Description: %s", i, tc.description)) {
			assert.Equal(t, tc.expected.dt, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}

func Example_parseDrawNum() {
	result, err := parseDrawNum("10", 10)
	fmt.Printf("Result: %v Error: %v\n", result, err)

	result, err = parseDrawNum("1", 10)
	fmt.Printf("Result: %v Error: %v\n", result, err)

	result, err = parseDrawNum("1a", 10)
	fmt.Printf("Result: %v Error: %v\n", result, err)

	result, err = parseDrawNum("0", 10)
	fmt.Printf("Result: %v Error: %v\n", result, err)

	result, err = parseDrawNum("11", 10)
	fmt.Printf("Result: %v Error: %v\n", result, err)

	// Output:
	// Result: 10 Error: <nil>
	// Result: 1 Error: <nil>
	// Result: 0 Error: invalid draw digit: strconv.Atoi: parsing "1a": invalid syntax
	// Result: 0 Error: draw out of range: got 0 max 10
	// Result: 0 Error: draw out of range: got 11 max 10
}

func Example_parseDrawSeq() {
	num, err := parseDrawSeq("10000")
	fmt.Println(num, err)

	num, err = parseDrawSeq("1a")
	fmt.Println(num, err)

	// Output:
	// 10000 <nil>
	// 0 invalid draw seq: strconv.Atoi: parsing "1a": invalid syntax
}
