package tball

import (
	"log"
	"regexp"
	"time"
)

const (
	CSVUrl = "https://www.national-lottery.co.uk/results/thunderball/draw-history/csv"
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
	TBall     uint8        `json:"tball"`
	BallSet   string       `json:"ball_set"`
	Machine   string       `json:"machine"`
	DrawNo    uint64       `json:"draw_no"`
}

type DrawChan struct {
	Draw Draw
	Err  error
}

func IsValidBall(arg string) bool {
	pattern := `^\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|50)\b(,\b([1-9]|1[0-9]|2[0-9]|3[0-9])\b)*$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
		return matched
	}
	return matched
}

func IsValidStars(arg string) bool {
	pattern := `^\b([1-9]|1[0-4]\b)*$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}
