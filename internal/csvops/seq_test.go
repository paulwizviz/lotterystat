package csvops

var (
	drawSeqs = []struct {
		input    string
		expected struct {
			result uint64
			err    error
		}
	}{
		{
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
			input: "1a",
			expected: struct {
				result uint64
				err    error
			}{
				result: 0,
				err:    ErrCSVInvalidDrawSeq,
			},
		},
		{
			input: "-1",
			expected: struct {
				result uint64
				err    error
			}{
				result: 0,
				err:    ErrCSVInvalidDrawSeq,
			},
		},
	}
)
