package lotto

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
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Ball 6,Bonus Ball,Ball Set,Machine,DrawNumber
18-Feb-2026,1,11,12,13,18,49,33,L10,Lotto4,3147
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{
						DrawDate:  time.Date(2026, time.February, 18, 0, 0, 0, 0, time.UTC),
						DayOfWeek: time.Wednesday,
						Ball1:     1,
						Ball2:     11,
						Ball3:     12,
						Ball4:     13,
						Ball5:     18,
						Ball6:     49,
						BonusBall: 33,
						BallSet:   "L10",
						Machine:   "Lotto4",
						DrawNo:    3147,
					},
					Err: nil,
				},
			},
		},
		// UnhappyPath
		{
			name: fmt.Sprintf("%s-invalid date", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Ball 6,Bonus Ball,Ball Set,Machine,DrawNumber
18-1-2026,1,11,12,13,18,49,33,L10,Lotto4,3147
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
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Ball 6,Bonus Ball,Ball Set,Machine,DrawNumber
18-Feb-2026,0,11,12,13,18,49,33,L10,Lotto4,3147
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
			name: fmt.Sprintf("%s-invalid ball6", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Ball 6,Bonus Ball,Ball Set,Machine,DrawNumber
18-Feb-2026,1,11,12,13,18,60,33,L10,Lotto4,3147
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrBall6,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid bonus", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Ball 6,Bonus Ball,Ball Set,Machine,DrawNumber
18-Feb-2026,1,11,12,13,18,49,60,L10,Lotto4,3147
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrBonus,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid seq", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Ball 6,Bonus Ball,Ball Set,Machine,DrawNumber
18-Feb-2026,1,11,12,13,18,49,33,L10,Lotto4,abc
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
