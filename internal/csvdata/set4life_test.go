package csvdata

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProcessS4LCVS(t *testing.T) {
	testcases := []struct {
		input       []byte
		expected    []Set4LifeDraw
		description string
	}{
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Jan-2023,42,30,47,40,15,8,SFL1,Excalibur 5,402
16-Jan-2023,36,10,23,40,32,10,SFL1,Excalibur 5,401`),
			expected: []Set4LifeDraw{
				{
					Log: map[string]string{
						CSVLogKeyLineNo: "2",
					},
					Err: nil,
					Item: struct {
						DrawDate  time.Time
						DayOfWeek time.Weekday
						Ball1     uint8
						Ball2     uint8
						Ball3     uint8
						Ball4     uint8
						Ball5     uint8
						LifeBall  uint8
						BallSet   string
						Machine   string
						DrawNo    uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("19-Jan-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("19-Jan-2023")
							return dt.Weekday()
						}(),
						Ball1:    uint8(42),
						Ball2:    uint8(30),
						Ball3:    uint8(47),
						Ball4:    uint8(40),
						Ball5:    uint8(15),
						LifeBall: uint8(8),
						BallSet:  "SFL1",
						Machine:  "Excalibur 5",
						DrawNo:   uint64(402),
					},
				},
				{
					Log: map[string]string{
						CSVLogKeyLineNo: "3",
					},
					Err: nil,
					Item: struct {
						DrawDate  time.Time
						DayOfWeek time.Weekday
						Ball1     uint8
						Ball2     uint8
						Ball3     uint8
						Ball4     uint8
						Ball5     uint8
						LifeBall  uint8
						BallSet   string
						Machine   string
						DrawNo    uint64
					}{
						DrawDate: func() time.Time {
							dt, _ := parseDateTime("16-Jan-2023")
							return dt
						}(),
						DayOfWeek: func() time.Weekday {
							dt, _ := parseDateTime("16-Jan-2023")
							return dt.Weekday()
						}(),
						Ball1:    uint8(36),
						Ball2:    uint8(10),
						Ball3:    uint8(23),
						Ball4:    uint8(40),
						Ball5:    uint8(32),
						LifeBall: uint8(10),
						BallSet:  "SFL1",
						Machine:  "Excalibur 5",
						DrawNo:   uint64(401),
					},
				},
			},
			description: "Valid records",
		},
	}

	for i, tc := range testcases {
		sig := ProcessS4LCVS(bytes.NewReader(tc.input))
		idx := 0
		for s := range sig {
			if assert.True(t, errors.Is(s.Err, tc.expected[idx].Err), fmt.Sprintf("Case: %d Description: %s", i, tc.description)) {
				assert.Equal(t, tc.expected[idx].Log, s.Log, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
				assert.Equal(t, tc.expected[idx].Item, s.Item, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
			}
			idx++
		}
	}
}
