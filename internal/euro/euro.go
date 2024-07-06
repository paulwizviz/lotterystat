package euro

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"time"
)

const (
	CSVUrl = "https://www.national-lottery.co.uk/results/euromillions/draw-history/csv"
)

// Draw represents a line from euro draw results
type Draw struct {
	DrawDate  time.Time    `json:"draw_date"`
	DayOfWeek time.Weekday `json:"day_of_week"`
	Ball1     uint8        `json:"ball1"`
	Ball2     uint8        `json:"ball2"`
	Ball3     uint8        `json:"ball3"`
	Ball4     uint8        `json:"ball4"`
	Ball5     uint8        `json:"ball5"`
	LS1       uint8        `json:"ls1"`
	LS2       uint8        `json:"ls2"`
	UKMarker  string       `json:"uk_marker"`
	DrawNo    uint64       `json:"draw_no"`
}

type DrawChan struct {
	Draw Draw
	Err  error
}

func CreateTable(ctx context.Context, db *sql.DB) error {
	return createTable(ctx, db)
}

// func PersistsCSV(ctx context.Context, db *sql.DB, nworkers int) error {
// 	return persistsCSV(ctx, db, nworkers)
// }

func IsValidBall(arg string) bool {
	pattern := `^\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|50)\b(,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|50)\b)*$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
		return matched
	}
	return matched
}

func IsValidStars(arg string) bool {
	pattern := `^\b([1-9]|1[0-2])\b(,\b([1-9]|1[0-2])\b)*$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}
