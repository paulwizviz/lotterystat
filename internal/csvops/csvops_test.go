package csvops

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var csvScenarios = []struct {
	name     string
	data     []byte
	expected []CSVRec
}{
	{
		name: "Valid csv content",
		data: []byte(`d1,d2,d3
1,a,c
2,2,7`),
		expected: []CSVRec{
			{Record: []string{"1", "a", "c"}, Line: uint(1), Err: nil},
			{Record: []string{"2", "2", "7"}, Line: uint(2), Err: nil},
		},
	},
	{
		name: "Invalid csv line",
		data: []byte(`d1,d2,d3
1,a,c
2,2`),
		expected: []CSVRec{
			{Record: []string{"1", "a", "c"}, Line: uint(1), Err: nil},
			{Record: []string{"2", "2"}, Line: uint(2), Err: fmt.Errorf("%w-%s", ErrLine, "record on line 3: wrong number of fields")},
		},
	},
}

func TestExtractRec(t *testing.T) {
	for i, scenario := range csvScenarios {
		t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), func(t *testing.T) {
			r := bytes.NewReader(scenario.data)
			actual := []CSVRec{}
			recs := ExtractRec(context.TODO(), r)
			for rec := range recs {
				actual = append(actual, rec)
			}
			assert.Equal(t, scenario.expected, actual)
		})
	}
}

var csvCancelScenarios = []struct {
	name     string
	ctxFn    func(context.CancelFunc)
	data     []byte
	expected []CSVRec
}{
	{
		name: "Context cancer",
		ctxFn: func(cf context.CancelFunc) {
			go func(cancel context.CancelFunc) {
				time.Sleep(1 * time.Millisecond)
				cancel()
			}(cf)
		},
		data: []byte(`d1,d2,d3
1,a,c
2,2,7
3,z,d
4,x,1
5,x7,100`),
		expected: []CSVRec{
			{Record: []string{"1", "a", "c"}, Line: 0x1, Err: error(nil)},
			{Record: []string{"2", "2", "7"}, Line: 0x2, Err: error(nil)},
		},
	},
	{
		name: "Context cancer",
		ctxFn: func(cf context.CancelFunc) {
			go func(cancel context.CancelFunc) {
				time.Sleep(1 * time.Millisecond)
				cancel()
			}(cf)
		},
		data: []byte(`d1,d2,d3
1,a,c
2,2,7
3,z,d
4,x,1
5,a,c
6,2,7
7,z,d
8,x,1
9,x7,100`),
		expected: []CSVRec{
			{Record: []string{"1", "a", "c"}, Line: 0x1, Err: error(nil)},
			{Record: []string{"2", "2", "7"}, Line: 0x2, Err: error(nil)},
		},
	},
}

func TestExtractRec_cancel(t *testing.T) {
	for i, scenario := range csvCancelScenarios {
		t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			scenario.ctxFn(cancel)
			r := bytes.NewReader(scenario.data)
			actual := []CSVRec{}
			recs := ExtractRec(ctx, r)
			for rec := range recs {
				actual = append(actual, rec)
				time.Sleep(10 * time.Millisecond)
			}
			assert.Equal(t, scenario.expected, actual)
		})
	}
}
