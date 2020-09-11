package csvdata

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcEuroCSV(t *testing.T) {
	testcases := []struct {
		input       bytes.Buffer
		expected    euroChanSignal
		description string
	}{
		{
			input: func() bytes.Buffer {
				var buf bytes.Buffer
				buf.WriteString(`DrawDate,Ball 1,Ball 2,Ball 3,Ball 4,Ball 5,Lucky Star 1,Lucky Star 2,UK Millionaire Maker,European Millionaire Maker,DrawNumber
04-Apr-2023,10,16,31,33,50,3,8,"XCRG53171","",1621
31-Mar-2023,16,18,28,34,47,5,10,"JBQS10867","",1620`)
				return buf
			}(),
			expected: euroChanSignal{
				Err: nil,
			},
			description: "Valid csv",
		},
	}

	for i, tc := range testcases {

		actualChanSig := processEuroCVS(&tc.input)

		for c := range actualChanSig {
			assert.Equal(t, tc.expected.Err, c.Err, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}

	}
}
