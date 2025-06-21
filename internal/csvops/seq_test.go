package csvops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var seqScenarios = []struct {
	name     string
	input    string
	expected struct {
		result uint64
		err    error
	}
}{
	{
		name:  "Valid sequence number",
		input: "1000",
		expected: struct {
			result uint64
			err    error
		}{
			result: 1000,
			err:    nil,
		},
	},
	{
		name:  "Mixed numeric and akphabet",
		input: "1a",
		expected: struct {
			result uint64
			err    error
		}{
			result: 0,
			err:    ErrInvalidDrawSeq,
		},
	},
	{
		name:  "Negative number",
		input: "-1",
		expected: struct {
			result uint64
			err    error
		}{
			result: 0,
			err:    ErrInvalidDrawSeq,
		},
	},
}

func TestParseSeq(t *testing.T) {
	for i, scenario := range seqScenarios {
		t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), func(t *testing.T) {
			actual, err := ParseDrawSeq(scenario.input)
			if assert.ErrorIs(t, err, scenario.expected.err) {
				assert.Equal(t, scenario.expected.result, actual)
			}
		})
	}
}
