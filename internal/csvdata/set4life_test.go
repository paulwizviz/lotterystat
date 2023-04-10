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
		expected    []Set4LifeDrawSig
		description string
	}{
		{
			input: []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Jan-2023,42,30,47,40,15,8,SFL1,Excalibur 5,402
16-Jan-2023,36,10,23,40,32,10,SFL1,Excalibur 5,401`),
			expected: []Set4LifeDrawSig{
				{
					Log: map[string]string{
						CSVLogKeyLineNo: "2",
					},
					Err: nil,
					Item: Set4LifeDraw{
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
					Item: Set4LifeDraw{
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

func TestS4LDrawSQLiteTags(t *testing.T) {
	testcases := []struct {
		input       Set4LifeDraw
		expected    map[string]string
		description string
	}{
		{
			input: Set4LifeDraw{},
			expected: map[string]string{
				"DrawDate":  "draw_date,INTEGER",
				"DayOfWeek": "day_of_week,INTEGER",
				"Ball1":     "ball1,INTEGER",
				"Ball2":     "ball2,INTEGER",
				"Ball3":     "ball3,INTEGER",
				"Ball4":     "ball4,INTEGER",
				"Ball5":     "ball5,INTEGER",
				"LifeBall":  "life_ball,INTEGER",
				"BallSet":   "ball_set,TEXT",
				"Machine":   "machine,TEXT",
				"DrawNo":    "draw_no,TEXT",
			},
			description: "Valid tags",
		},
	}
	for i, tc := range testcases {
		actual := tc.input.SQLiteTags()
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
	}
}
