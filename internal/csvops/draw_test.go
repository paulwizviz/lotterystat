package csvops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	drawScenarios = []struct {
		name  string
		input struct {
			value  string
			maxval int
		}
		expected struct {
			result uint8
			err    error
		}
	}{
		{
			name: "valid string less than max",
			input: struct {
				value  string
				maxval int
			}{
				value:  "1",
				maxval: 2,
			},
			expected: struct {
				result uint8
				err    error
			}{
				result: 1,
				err:    nil,
			},
		},
		{
			name: "valid value greater than max",
			input: struct {
				value  string
				maxval int
			}{
				value:  "2",
				maxval: 1,
			},
			expected: struct {
				result uint8
				err    error
			}{
				result: 0,
				err:    ErrInvalidDrawRange,
			},
		},
		{
			name: "zero value",
			input: struct {
				value  string
				maxval int
			}{
				value:  "0",
				maxval: 1,
			},
			expected: struct {
				result uint8
				err    error
			}{
				result: 0,
				err:    ErrInvalidDrawRange,
			},
		},
		{
			name: "invalid value format",
			input: struct {
				value  string
				maxval int
			}{
				value:  "a",
				maxval: 1,
			},
			expected: struct {
				result uint8
				err    error
			}{
				result: 0,
				err:    ErrInvalidDrawDigit,
			},
		},
	}
)

func TestParseDrawNum(t *testing.T) {
	for i, scenario := range drawScenarios {
		t.Run(fmt.Sprintf("case %d-%s", i, scenario.name), func(t *testing.T) {
			actual, err := ParseDrawNum(scenario.input.value, scenario.input.maxval)
			if assert.ErrorIs(t, err, scenario.expected.err) {
				assert.Equal(t, scenario.expected.result, actual)
			}
		})
	}
}
