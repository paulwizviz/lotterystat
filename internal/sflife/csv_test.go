package sflife

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
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Feb-2026,5,9,13,34,45,8,SFL3,Excalibur6,724
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{
						DrawDate:  time.Date(2026, time.February, 19, 0, 0, 0, 0, time.UTC),
						DayOfWeek: time.Thursday,
						Ball1:     5,
						Ball2:     9,
						Ball3:     13,
						Ball4:     34,
						Ball5:     45,
						LBall:     8,
						BallSet:   "SFL3",
						Machine:   "Excalibur6",
						DrawNo:    724,
					},
					Err: nil,
				},
			},
		},
		// UnhappyPath
		{
			name: fmt.Sprintf("%s-invalid date", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-1-2026,5,9,13,34,45,8,SFL3,Excalibur6,724
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
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Feb-2026,0,9,13,34,45,8,SFL3,Excalibur6,724
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
			name: fmt.Sprintf("%s-invalid ball5", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Feb-2026,5,9,13,34,48,8,SFL3,Excalibur6,724
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrBall5,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid LBall", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Feb-2026,5,9,13,34,45,11,SFL3,Excalibur6,724
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrLBall,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid seq", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Life Ball,Ball Set,Machine,DrawNumber
19-Feb-2026,5,9,13,34,45,8,SFL3,Excalibur6,abc
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
