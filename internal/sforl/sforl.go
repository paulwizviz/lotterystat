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

func IsValidBet(arg string) bool {
	pattern := `^\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-7])\b,\b([1-9]|10)\b$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}

type BallCount struct {
	Ball  uint8
	Count uint
}

func BallFreq(ctx context.Context, db *sql.DB) ([]BallCount, error) {
	stmt, err := prepCountBallStmt(ctx, db)
	if err != nil {
		return nil, err
	}

	ballcounts := []BallCount{}
	for i := 1; i < 48; i++ {
		var bc BallCount
		bc.Ball = uint8(i)
		count, err := countChoice(ctx, stmt, uint8(i))
		if err != nil {
			continue
		}
		bc.Count = count
		ballcounts = append(ballcounts, bc)
	}
	return ballcounts, nil
}

type StarCount struct {
	Star  uint8
	Count uint
}

func StarFreq(ctx context.Context, db *sql.DB) ([]StarCount, error) {
	stmt, err := prepCountLuckyStmt(ctx, db)
	if err != nil {
		return nil, err
	}

	starCounts := []StarCount{}
	for i := 1; i < 11; i++ {
		var bc StarCount
		bc.Star = uint8(i)
		count, err := countChoice(ctx, stmt, uint8(i))
		if err != nil {
			continue
		}
		bc.Count = count
		starCounts = append(starCounts, bc)
	}
	return starCounts, nil
}
