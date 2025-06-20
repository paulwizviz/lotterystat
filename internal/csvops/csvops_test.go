package csvops

import (
	"bytes"
	"context"
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
				err: ErrCSVInvalidDayFmt,
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
				err: ErrCSVInvalidYearFmt,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
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
				err: ErrCSVInvalidDaysInMonth,
			},
			description: "Invalid day in Dec",
		},
	}

	for i, tc := range tcases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			actual, err := ParseDateTime(tc.input)
			if assert.True(t, errors.Is(err, tc.expected.err)) {
				assert.Equal(t, tc.expected.dt, actual)
			}
		})
	}
}

func TestParseDrawNum(t *testing.T) {
	testcases := []struct {
		input    string
		expected struct {
			result uint8
			err    error
		}
	}{
		{
			input: "1",
			expected: struct {
				result uint8
				err    error
			}{
				result: 1,
				err:    nil,
			},
		},
		{
			input: "10",
			expected: struct {
				result uint8
				err    error
			}{
				result: 10,
				err:    nil,
			},
		},
		{
			input: "0",
			expected: struct {
				result uint8
				err    error
			}{
				result: 0,
				err:    ErrCSVInvalidDrawRange,
			},
		},
		{
			input: "11",
			expected: struct {
				result uint8
				err    error
			}{
				result: 0,
				err:    ErrCSVInvalidDrawRange,
			},
		},
		{
			input: "1a",
			expected: struct {
				result uint8
				err    error
			}{
				result: 0,
				err:    ErrCSVInvalidDrawDigit,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			actual, err := ParseDrawNum(tc.input, 10)
			if assert.True(t, errors.Is(err, tc.expected.err)) {
				assert.Equal(t, int(tc.expected.result), int(actual))
			}
		})
	}
}

func TestParseDrawSeq(t *testing.T) {
	testcases := []struct {
		input    string
		expected struct {
			result uint64
			err    error
		}
	}{
		{
			input: "1000",
			expected: struct {
				result uint64
				err    error
			}{
				result: 1000,
				err:    nil,
			},
		},
		{
			input: "1a",
			expected: struct {
				result uint64
				err    error
			}{
				result: 0,
				err:    ErrCSVInvalidDrawSeq,
			},
		},
		{
			input: "-1",
			expected: struct {
				result uint64
				err    error
			}{
				result: 0,
				err:    ErrCSVInvalidDrawSeq,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			actual, err := ParseDrawSeq(tc.input)
			if assert.True(t, errors.Is(err, tc.expected.err)) {
				assert.Equal(t, int(tc.expected.result), int(actual))
			}
		})
	}
}

func TestExtractRec(t *testing.T) {
	testcases := []struct {
		input    []byte
		expected CSVRec
	}{
		{
			input: []byte(`Index,Value,Date
1,abc,29-Sep-2023`),
			expected: CSVRec{
				Record: []string{"1", "abc", "29-Sep-2023"},
				Err:    nil,
			},
		},
		{
			input: []byte(`Index,Value,Date
"abc","29-Sep-2023"`),
			expected: CSVRec{
				Record: []string{"abc", "29-Sep-2023"},
				Err:    ErrCSVLine,
			},
		},
	}

	for i, tc := range testcases {
		recs := ExtractRec(context.TODO(), bytes.NewReader(tc.input))
		for actual := range recs {
			if assert.True(t, errors.Is(actual.Err, tc.expected.Err), fmt.Sprintf("Case: %d Error: %v", i, actual.Err)) {
				assert.Equal(t, tc.expected.Record, actual.Record, fmt.Sprintf("Case: %d", i))
			}
		}
	}
}

func Example_extractCSV_cancel() {
	data := []byte(`d1,d2,d3
1,a,c
2,2,7
3,z,d
4,x,1
5,x7,100`)
	r := bytes.NewReader(data)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Millisecond)
		cancel()
	}()
	records := ExtractRec(ctx, r)
	for rec := range records {
		fmt.Println(rec)
		time.Sleep(10 * time.Millisecond)
	}
	// Output:
	// {[1 a c] 2 <nil>}
	// {[2 2 7] 3 <nil>}
}
