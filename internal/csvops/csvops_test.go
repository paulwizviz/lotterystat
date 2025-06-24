package csvops

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type extractRecScenario struct {
	name     string
	data     []byte
	expected []CSVRec
}

func (e extractRecScenario) noCancel(t *testing.T) {
	r := bytes.NewReader(e.data)
	actual := []CSVRec{}
	recs := ExtractRec(context.TODO(), r)
	for rec := range recs {
		actual = append(actual, rec)
	}
	assert.Equal(t, e.expected, actual)
}

func (e extractRecScenario) withCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	r := bytes.NewReader(e.data)
	go func(context.CancelFunc) {
		time.Sleep(1 * time.Millisecond)
		cancel()
	}(cancel)
	actual := []CSVRec{}
	recs := ExtractRec(ctx, r)
	for rec := range recs {
		actual = append(actual, rec)
		time.Sleep(10 * time.Millisecond)
	}
	assert.Equal(t, e.expected, actual)
}

var (
	noCancel   = "No cancel"
	withCancel = "With cancel"

	scenarios = []extractRecScenario{
		// No cancel
		{
			name: fmt.Sprintf("%s-Valid csv content", noCancel),
			data: []byte(`d1,d2,d3
1,a,c
2,2,7`),
			expected: []CSVRec{
				{Header: []string{"d1", "d2", "d3"}, Record: []string{"1", "a", "c"}, Line: uint(1), Err: nil},
				{Header: []string{"d1", "d2", "d3"}, Record: []string{"2", "2", "7"}, Line: uint(2), Err: nil},
			},
		},
		{
			name: fmt.Sprintf("%s-Invalid csv line", noCancel),
			data: []byte(`d1,d2,d3
1,a,c
2,2`),
			expected: []CSVRec{
				{Header: []string{"d1", "d2", "d3"}, Record: []string{"1", "a", "c"}, Line: uint(1), Err: nil},
				{Header: []string{"d1", "d2", "d3"}, Record: []string{"2", "2"}, Line: uint(2), Err: fmt.Errorf("%w-%s", ErrLine, "record on line 3: wrong number of fields")},
			},
		},
		{
			name: fmt.Sprintf("%s-Missmatched header", noCancel),
			data: []byte(`d1,d2
1,a,c
2,2,b`),
			expected: []CSVRec{
				{Header: []string{"d1", "d2"}, Record: []string{"1", "a", "c"}, Line: uint(1), Err: fmt.Errorf("%w-%s", ErrLine, "record on line 2: wrong number of fields")},
				{Header: []string{"d1", "d2"}, Record: []string{"2", "2", "b"}, Line: uint(2), Err: fmt.Errorf("%w-%s", ErrLine, "record on line 3: wrong number of fields")},
			},
		},
		// Context cancel called
		{
			name: fmt.Sprintf("%s-cancel called after 1 ms", withCancel),
			data: []byte(`d1,d2,d3
1,a,c
2,2,7
3,z,d
4,x,1
5,x7,100`),
			expected: []CSVRec{
				{Header: []string{"d1", "d2", "d3"}, Record: []string{"1", "a", "c"}, Line: 0x1, Err: error(nil)},
				{Header: []string{"d1", "d2", "d3"}, Record: []string{"2", "2", "7"}, Line: 0x2, Err: error(nil)},
			},
		},
	}
)

func TestExtractRec(t *testing.T) {
	for i, scenario := range scenarios {
		if strings.Contains(scenario.name, noCancel) {
			t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), scenario.noCancel)
		} else {
			t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), scenario.withCancel)
		}
	}
}
