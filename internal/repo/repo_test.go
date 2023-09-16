package repo

import (
	"fmt"
	"paulwizviz/lotterystat/internal/euro"
	"paulwizviz/lotterystat/internal/sforl"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqliteTags(t *testing.T) {

	testcases := []struct {
		input       any
		expected    []structTag
		description string
	}{
		{
			input: &euro.Draw{},
			expected: []structTag{
				{
					FieldName: "DrawDate",
					Tag:       "draw_date,INTEGER",
				},
				{
					FieldName: "DayOfWeek",
					Tag:       "day_of_week,INTEGER",
				},
				{
					FieldName: "Ball1",
					Tag:       "ball1,INTEGER",
				},
				{
					FieldName: "Ball2",
					Tag:       "ball2,INTEGER",
				},
				{
					FieldName: "Ball3",
					Tag:       "ball3,INTEGER",
				},
				{
					FieldName: "Ball4",
					Tag:       "ball4,INTEGER",
				},
				{
					FieldName: "Ball5",
					Tag:       "ball5,INTEGER",
				},
				{
					FieldName: "LS1",
					Tag:       "ls1,INTEGER",
				},
				{
					FieldName: "LS2",
					Tag:       "ls2,INTEGER",
				},
				{
					FieldName: "UKMarker",
					Tag:       "uk_marker,TEXT",
				},
				{
					FieldName: "EuroMarker",
					Tag:       "euro_marker,TEXT",
				},
				{
					FieldName: "DrawNo",
					Tag:       "draw_no,INTEGER",
				},
			},
			description: "EuroDraw tags",
		},
	}

	for i, tc := range testcases {
		switch v := tc.input.(type) {
		case *euro.Draw:
			actual := sqliteTags(v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}

func TestCreateTblStmt(t *testing.T) {

	testcases := []struct {
		input       any
		expected    string
		description string
	}{
		{
			input:       &euro.Draw{},
			expected:    "CREATE TABLE IF NOT EXISTS Draw ( draw_date INTEGER, day_of_week INTEGER, ball1 INTEGER, ball2 INTEGER, ball3 INTEGER, ball4 INTEGER, ball5 INTEGER, ls1 INTEGER, ls2 INTEGER, uk_marker TEXT, euro_marker TEXT, draw_no INTEGER PRIMARY KEY )",
			description: "Euro Draw Table",
		},
		{
			input:       &euro.Draw{},
			expected:    "CREATE TABLE IF NOT EXISTS Draw ( draw_date INTEGER, day_of_week INTEGER, ball1 INTEGER, ball2 INTEGER, ball3 INTEGER, ball4 INTEGER, ball5 INTEGER, ls1 INTEGER, ls2 INTEGER, uk_marker TEXT, euro_marker TEXT, draw_no INTEGER PRIMARY KEY )",
			description: "Set for Life Draw Table",
		},
	}

	for i, tc := range testcases {
		switch v := tc.input.(type) {
		case *euro.Draw:
			actual := CreateTblStmt(v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		case *sforl.Draw:
			actual := CreateTblStmt(v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}