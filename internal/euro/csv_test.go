package euro

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/paulwizviz/lotterystat/internal/csvops"
	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name     string
	input    io.Reader
	expected []DrawChan
}

func (tc testcase) happyPath(t *testing.T) {
	csvRecs := csvops.ExtractRec(context.TODO(), tc.input)
	actual := ProcessCSV(csvRecs, 1)
	assert.Equal(t, tc.expected, actual)
}

func (tc testcase) unhappyPath(t *testing.T) {
	csvRecs := csvops.ExtractRec(context.TODO(), tc.input)
	actual := ProcessCSV(csvRecs, 1)
	assert.ErrorIs(t, actual[0].Err, tc.expected[0].Err)
}

var (
	testHappyPath  = "happy path"
	testUnappyPath = "unhappy path"

	testcases = []testcase{
		// Happy Path
		{
			name: testHappyPath,
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,Ball Set,Machine,DrawNumber
20-Feb-2026,13,24,28,33,35,5,9,ZDTF34718,,21,13,1922
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{
						DrawDate:  time.Date(2026, time.February, 20, 0, 0, 0, 0, time.UTC),
						DayOfWeek: time.Friday,
						Ball1:     13,
						Ball2:     24,
						Ball3:     28,
						Ball4:     33,
						Ball5:     35,
						Star1:     5,
						Star2:     9,
						UKMaker:   "ZDTF34718",
						EUMaker:   "",
						BallSet:   "21",
						Machine:   "13",
						DrawNo:    1922,
					},
					Err: nil,
				},
			},
		},
		// UnhappyPath
		{
			name: fmt.Sprintf("%s-invalid date", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,Ball Set,Machine,DrawNumber
20-1-2026,13,24,28,33,35,5,9,ZDTF34718,,21,13,1922
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrDrawDate,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid ball1", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,Ball Set,Machine,DrawNumber
20-Feb-2026,0,24,28,33,35,5,9,ZDTF34718,,21,13,1922
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrBall1,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid star1", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,Ball Set,Machine,DrawNumber
20-Feb-2026,13,24,28,33,35,13,9,ZDTF34718,,21,13,1922
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrStar1,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid seq", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,Ball Set,Machine,DrawNumber
20-Feb-2026,13,24,28,33,35,5,9,ZDTF34718,,21,13,abc
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrSeq,
				},
			},
		},
	}
)

func TestProcessCSV(t *testing.T) {
	for _, tc := range testcases {
		if strings.Contains(tc.name, testUnappyPath) {
			t.Run(tc.name, tc.unhappyPath)
		} else {
			t.Run(tc.name, tc.happyPath)
		}
	}
}
