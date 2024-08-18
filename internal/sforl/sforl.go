// Package sforl represents data structures for Set For Life draw
package sforl

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

const (
	CSVUrl = "https://www.national-lottery.co.uk/results/set-for-life/draw-history/csv"
)

const (
	tblName   = "set_for_life"
	drawDate  = "draw_date"
	dayOfWeek = "day_of_week"
	ball1     = "ball1"
	ball2     = "ball2"
	ball3     = "ball3"
	ball4     = "ball4"
	ball5     = "ball5"
	luckyBall = "lb"
	ballset   = "ball_set"
	machine   = "machine"
	drawNo    = "draw_no"
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

func freqBallSQL(b uint8) string {
	return fmt.Sprintf("SELECT COUNT(*) FROM %[1]s WHERE %[2]s=%[7]d AND %[3]s=%[7]d AND %[4]s=%[7]d AND %[5]s=%[7]d AND %[6]s=%[7]d", tblName, ball1, ball2, ball3, ball4, ball5, b)
}
