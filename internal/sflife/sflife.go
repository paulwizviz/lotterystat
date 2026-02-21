package sflife

import (
	"errors"
	"log"
	"regexp"
	"time"
)

const (
	CSVUrl = "https://www.national-lottery.co.uk/results/set-for-life/draw-history/csv"
)

var (
	ErrDrawDate = errors.New("invalid draw date")
	ErrBall1    = errors.New("invalid ball 1")
	ErrBall2    = errors.New("invalid ball 2")
	ErrBall3    = errors.New("invalid ball 3")
	ErrBall4    = errors.New("invalid ball 4")
	ErrBall5    = errors.New("invalid ball 5")
	ErrLBall    = errors.New("invalid life ball")
	ErrSeq      = errors.New("invalid seq")
	ErrRec      = errors.New("invalid record")
)

// Draw represents a line from set for life draw results
type Draw struct {
	DrawDate  time.Time    `json:"draw_date"`
	DayOfWeek time.Weekday `json:"day_of_week"`
	Ball1     uint8        `json:"ball1"`
	Ball2     uint8        `json:"ball2"`
	Ball3     uint8        `json:"ball3"`
	Ball4     uint8        `json:"ball4"`
	Ball5     uint8        `json:"ball5"`
	LBall     uint8        `json:"lball"`
	BallSet   string       `json:"ball_set"`
	Machine   string       `json:"machine"`
	DrawNo    uint64       `json:"draw_no"`
}

type DrawChan struct {
	Draw Draw
	Err  error
}

func IsValidBall(arg string) bool {
	pattern := `^\b([1-9]|[1-3][0-9]|4[0-7])\b(,\b([1-9]|[1-3][0-9]|4[0-7])\b)*$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
		return matched
	}
	return matched
}

func IsValidLifeBall(arg string) bool {
	pattern := `^\b([1-9]|10)\b$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}
