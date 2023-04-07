package csvdata

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// This test is intended to validate the ability to process
// EuroCSV properly
func TestProcessEuroCSV(t *testing.T) {
	testcases := []struct {
		input       []byte
		expected    []EuroDraw
		description string
	}{
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,31,33,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("04-Apr-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("04-Apr-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(10),
						Ball2:      uint8(16),
						Ball3:      uint8(31),
						Ball4:      uint8(33),
						Ball5:      uint8(50),
						LS1:        uint8(3),
						LS2:        uint8(8),
						UKMarker:   "XCRG53171",
						EuroMarker: "",
						DrawNo:     uint64(1621),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "All valid draws",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
a-Apr-2023,10,16,31,33,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    ErrInvalidNumFmt,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							return time.Time{}
						}(),
						DayOfWeek: func() time.Weekday {
							return time.Sunday
						}(),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "Invalid date format",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,a,16,31,33,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    ErrInvalidDrawDigit,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							return time.Time{}
						}(),
						DayOfWeek: func() time.Weekday {
							return time.Sunday
						}(),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "Invalid Ball 1",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,a,31,33,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    ErrInvalidDrawDigit,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							return time.Time{}
						}(),
						DayOfWeek: func() time.Weekday {
							return time.Sunday
						}(),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "Invalid Ball 2",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,a,33,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    ErrInvalidDrawDigit,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							return time.Time{}
						}(),
						DayOfWeek: func() time.Weekday {
							return time.Sunday
						}(),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "Invalid Ball 3",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,10,a,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    ErrInvalidDrawDigit,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							return time.Time{}
						}(),
						DayOfWeek: func() time.Weekday {
							return time.Sunday
						}(),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "Invalid Ball 4",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,10,1,a,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    ErrInvalidDrawDigit,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							return time.Time{}
						}(),
						DayOfWeek: func() time.Weekday {
							return time.Sunday
						}(),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "Invalid Ball 5",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,10,1,1,a,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    ErrInvalidDrawDigit,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							return time.Time{}
						}(),
						DayOfWeek: func() time.Weekday {
							return time.Sunday
						}(),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "Invalid LS 1",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,10,1,1,1,a,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					LineNo: 2,
					Err:    ErrInvalidDrawDigit,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							return time.Time{}
						}(),
						DayOfWeek: func() time.Weekday {
							return time.Sunday
						}(),
					},
				},
				{
					LineNo: 2,
					Err:    nil,
					Item: struct {
						DrawDate   time.Time
						DayOfWeek  time.Weekday
						Ball1      uint8
						Ball2      uint8
						Ball3      uint8
						Ball4      uint8
						Ball5      uint8
						LS1        uint8
						LS2        uint8
						UKMarker   string
						EuroMarker string
						DrawNo     uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("31-Mar-2023")
							return dt.Weekday()
						}(),
						Ball1:      uint8(16),
						Ball2:      uint8(18),
						Ball3:      uint8(28),
						Ball4:      uint8(34),
						Ball5:      uint8(47),
						LS1:        uint8(5),
						LS2:        uint8(10),
						UKMarker:   "JBQS10867",
						EuroMarker: "",
						DrawNo:     uint64(1620),
					},
				},
			},
			description: "Invalid LS 2",
		},
	}

	for i, tc := range testcases {
		sig := ProcessEuroCVS(bytes.NewReader(tc.input))
		idx := 0
		for s := range sig {
			if assert.True(t, errors.Is(s.Err, tc.expected[idx].Err), fmt.Sprintf("Case: %d.1 Description: %s", i, tc.description)) {
				assert.Equal(t, s.LineNo, tc.expected[idx].LineNo, fmt.Sprintf("Case: %d.1 Description: %s", i, tc.description))
				assert.Equal(t, s.Item, tc.expected[idx].Item, fmt.Sprintf("Case: %d.1 Description: %s", i, tc.description))
			}
			idx++
		}
	}
}
