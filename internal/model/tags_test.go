package model

import (
	"fmt"
	"reflect"
)

func Example_euroDrawTags() {
	e := EuroDraw{}

	got := SqliteTags(&e)
	want := make(map[string]string)

	want["DrawDate"] = "draw_date,INTEGER"
	want["DayOfWeek"] = "day_of_week,INTEGER"
	want["Ball1"] = "ball1,INTEGER"
	want["Ball2"] = "ball2,INTEGER"
	want["Ball3"] = "ball3,INTEGER"
	want["Ball4"] = "ball4,INTEGER"
	want["Ball5"] = "ball5,INTEGER"
	want["LS1"] = "ls1,INTEGER"
	want["LS2"] = "ls2,INTEGER"
	want["UKMarker"] = "uk_marker,TEXT"
	want["EuroMarker"] = "euro_marker,TEXT"
	want["DrawNo"] = "draw_no,INTEGER"

	result := reflect.DeepEqual(got, want)
	fmt.Println(result)

	// output:
	// true
}
