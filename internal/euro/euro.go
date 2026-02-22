package euro

import (
	"errors"
	"log"
	"regexp"
	"time"
)

const (
	CSVUrl = "https://www.national-lottery.co.uk/results/euromillions/draw-history/csv"
)

var (
	ErrDrawDate = errors.New("invalid draw date")
	ErrBall1    = errors.New("invalid ball 1")
	ErrBall2    = errors.New("invalid ball 2")
	ErrBall3    = errors.New("invalid ball 3")
	ErrBall4    = errors.New("invalid ball 4")
	ErrBall5    = errors.New("invalid ball 5")
	ErrStar1    = errors.New("invalid lucky star 1")
	ErrStar2    = errors.New("invalid lucky star 2")
	ErrSeq      = errors.New("invalid seq")
	ErrRec      = errors.New("invalid record")
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
	Star1     uint8        `json:"star1"`
	Star2     uint8        `json:"star2"`
	UKMaker   string       `json:"uk_maker"`
	EUMaker   string       `json:"eu_maker"`
	BallSet   string       `json:"ball_set"`
	Machine   string       `json:"machine"`
	DrawNo    uint64       `json:"draw_no"`
}

type DrawChan struct {
	Draw Draw
	Err  error
}

func IsValidBall(arg string) bool {
	pattern := `^\b([1-9]|[1-4][0-9]|50)\b(,\b([1-9]|[1-4][0-9]|50)\b)*$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
		return matched
	}
	return matched
}

func IsValidStars(arg string) bool {
	pattern := `^\b([1-9]|1[0-3])\b(,\b([1-9]|1[0-3])\b)*$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}
