// Package euro contains abstractions of the Euro Millions draw from UK National lottery and
// all related data analytics.
package euro

import (
	"context"
	"database/sql"
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

type Bet struct {
	Ball1 uint8 `json:"ball1"`
	Ball2 uint8 `json:"ball2"`
	Ball3 uint8 `json:"ball3"`
	Ball4 uint8 `json:"ball4"`
	Ball5 uint8 `json:"ball5"`
	LS1   uint8 `json:"ls1"`
	LS2   uint8 `json:"ls2"`
}

type MatchedDraw struct {
	Bet        Bet     `json:"bet"`
	Draw       Draw    `json:"draw"`
	Balls      []uint8 `json:"balls"`
	LuckyStars []uint8 `json:"luck_stars"`
}

type DrawChan struct {
	Draw Draw
	Err  error
}

func CreateTable(ctx context.Context, db *sql.DB) error {
	return createTable(ctx, db)
}

func MatchBets(ctx context.Context, db *sql.DB, bets []Bet) ([]MatchedDraw, error) {
	return matchBets(ctx, db, bets)
}

func PersistsCSV(ctx context.Context, db *sql.DB, nworkers int) error {
	return persistsCSV(ctx, db, nworkers)
}
