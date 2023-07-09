package result

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqliteTagsEuro(t *testing.T) {

	input := &EuroDraw{}
	expected := []structTag{
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
	}

	actual := sqliteTags[EuroDraw](input)
	assert.Equal(t, expected, actual, "Extracting EuroDraw sqlite tags")

}
