package tball

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
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
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,16,4,6,13,28,3,T6,Excalibur 1,3547
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{
						DrawDate:  time.Date(2024, time.August, 28, 0, 0, 0, 0, time.UTC),
						DayOfWeek: time.Wednesday,
						Ball1:     16,
						Ball2:     4,
						Ball3:     6,
						Ball4:     13,
						Ball5:     28,
						TBall:     3,
						BallSet:   "T6",
						Machine:   "Excalibur 1",
						DrawNo:    3547,
					},
					Err: nil,
				},
			},
		},
		// UnhappyPath
		{
			name: fmt.Sprintf("%s-invalid date", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-1-2024,16,4,6,13,28,3,T6,Excalibur 1,3547
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
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,0,4,6,13,28,3,T6,Excalibur 1,3547
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
			name: fmt.Sprintf("%s-invalid ball2", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,16,100,6,13,28,3,T6,Excalibur 1,3547
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrBall2,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid ball3", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,16,4,-1,13,28,3,T6,Excalibur 1,3547
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrBall3,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid ball4", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,16,4,6,100,28,3,T6,Excalibur 1,3547
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrBall4,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid ball5", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,16,4,6,13,0,3,T6,Excalibur 1,3547
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
			name: fmt.Sprintf("%s-invalid TBall", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,16,4,6,13,28,15,T6,Excalibur 1,3547
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrTBall,
				},
			},
		},
		{
			name: fmt.Sprintf("%s-invalid seq", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
28-Aug-2024,16,4,6,13,28,3,T6,Excalibur 1,abc
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
		{
			name: fmt.Sprintf("%s-invalid rec", testUnappyPath),
			input: func() io.Reader {
				b := []byte(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Thunderball,Ball Set,Machine,DrawNumber
Excalibur 1,abc
`)
				return bytes.NewReader(b)
			}(),
			expected: []DrawChan{
				{
					Draw: Draw{},
					Err:  ErrRec,
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

type benchmark struct {
	name       string
	numWorkers int
}

func BenchmarkWorker(b *testing.B) {

	benchmarks := []benchmark{}
	for i := range 8 {
		bm := benchmark{}
		bm.name = fmt.Sprintf("%d workers", i+1)
		bm.numWorkers = i + 1
		benchmarks = append(benchmarks, bm)
	}

	for _, bm := range benchmarks {
		inputFile := "./testdata/perform.csv"
		data, err := os.ReadFile(inputFile)
		if err != nil {
			b.Fatalf("Fail read file: %s", inputFile)
		}
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				reader := bytes.NewReader(data)
				reader.Seek(0, io.SeekStart)
				recs := csvops.ExtractRec(context.TODO(), reader)
				ProcessCSV(recs, bm.numWorkers)
			}
		})
	}
}
