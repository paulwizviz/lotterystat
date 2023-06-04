package euro

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/paulwizviz/lotterystat/internal/csvutil"

	"github.com/stretchr/testify/assert"
)

func TestProcessCSV(t *testing.T) {
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
					DrawDate: func() time.Time {
						dt, _ := csvutil.ParseDateTime("04-Apr-2023")
						return dt
					}(),
					DayOfWeek: func() time.Weekday {
						dt, _ := csvutil.ParseDateTime("04-Apr-2023")
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
				{
					DrawDate: func() time.Time {
						dt, _ := csvutil.ParseDateTime("31-Mar-2023")
						return dt
					}(),
					DayOfWeek: func() time.Weekday {
						dt, _ := csvutil.ParseDateTime("31-Mar-2023")
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
			description: "All valid draws",
		},
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
a-Apr-2023,10,16,31,33,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`),
			expected: []EuroDraw{
				{
					DrawDate: func() time.Time {
						dt, _ := csvutil.ParseDateTime("31-Mar-2023")
						return dt
					}(),
					DayOfWeek: func() time.Weekday {
						dt, _ := csvutil.ParseDateTime("31-Mar-2023")
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
			description: "Invalid day for 1st line",
		},
	}

	for i, tc := range testcases {
		draw := processCSV(context.TODO(), bytes.NewReader(tc.input))
		idx := 0
		for d := range draw {
			assert.Equal(t, tc.expected[idx], d, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
			idx++
		}
	}
}

// This test verify that
func Example_processCSV() {
	input := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,31,33,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`)

	draw := processCSV(context.TODO(), bytes.NewReader(input))
	for d := range draw {
		fmt.Println(d) // All draws will be displayed
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	draw = processCSV(ctx, bytes.NewReader(input))
	// The following step will not be called
	for d := range draw {
		fmt.Println(d)
	}

	// Output:
	// {2023-04-04 00:00:00 +0000 UTC Tuesday 10 16 31 33 50 3 8 XCRG53171  1621}
	// {2023-03-31 00:00:00 +0000 UTC Friday 16 18 28 34 47 5 10 JBQS10867  1620}
}
