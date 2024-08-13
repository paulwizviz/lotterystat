// Package sforl represents data structures for Set For Life draw
package sforl

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"time"
)

const (
	CSVUrl = "https://www.national-lottery.co.uk/results/set-for-life/draw-history/csv"
)

// Draw represents a draw from Set for Life
type Draw struct {
	DrawDate  time.Time    `json:"draw_date"`
	DayOfWeek time.Weekday `json:"day_of_week"`
	Ball1     uint8        `json:"ball1"`
	Ball2     uint8        `json:"ball2"`
	Ball3     uint8        `json:"ball3"`
	Ball4     uint8        `json:"ball4"`
	Ball5     uint8        `json:"ball5"`
	LifeBall  uint8        `json:"life_ball"`
	BallSet   string       `json:"ball_set"`
	Machine   string       `json:"machine"`
	DrawNo    uint64       `json:"draw_no"`
}

type DrawChan struct {
	Draw Draw
	Err  error
}

func CreateSQLiteTable(ctx context.Context, db *sql.DB) error {
	return createSQLiteTable(ctx, db)
}

func PersistsCSV(ctx context.Context, db *sql.DB, nworkers int) error {
	return persistsSQLiteCSV(ctx, db, nworkers)
}

func IsValidBet(arg string) bool {
	pattern := `^\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|10)\b$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}
