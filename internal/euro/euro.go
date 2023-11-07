// Package euro contains abstractions of the Euro Millions draw from UK National lottery and
// all related data analytics.
package euro

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
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

func IsValidBet(arg string) bool {
	pattern := `^\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|50)\b(,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|50)\b){4},\b([1-9]|1[0-2])\b,\b([1-9]|1[0-2])\b$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}

func ProcessBetArg(arg string) (Bet, error) {
	elems := strings.Split(arg, ",")
	var bArray []uint8
	b1, err := strconv.ParseInt(elems[0], 0, 0)
	if err != nil {
		return Bet{}, err
	}
	b2, err := strconv.ParseInt(elems[1], 0, 0)
	if err != nil {
		return Bet{}, err
	}
	b3, err := strconv.ParseInt(elems[2], 0, 0)
	if err != nil {
		return Bet{}, err
	}
	b4, err := strconv.ParseInt(elems[3], 0, 0)
	if err != nil {
		return Bet{}, err
	}
	b5, err := strconv.ParseInt(elems[4], 0, 0)
	if err != nil {
		return Bet{}, err
	}
	bArray = append(bArray, uint8(b1), uint8(b2), uint8(b3), uint8(b4), uint8(b5))
	sort.Slice(bArray, func(i, j int) bool {
		return bArray[i] < bArray[j]
	})

	var lsArray []uint8
	ls1, err := strconv.ParseInt(elems[5], 0, 0)
	if err != nil {
		return Bet{}, err
	}
	ls2, err := strconv.ParseInt(elems[6], 0, 0)
	if err != nil {
		return Bet{}, err
	}
	lsArray = append(lsArray, uint8(ls1), uint8(ls2))
	sort.Slice(lsArray, func(i, j int) bool {
		return lsArray[i] < lsArray[j]
	})
	b := Bet{
		Ball1: bArray[0],
		Ball2: bArray[1],
		Ball3: bArray[2],
		Ball4: bArray[3],
		Ball5: bArray[4],
		LS1:   lsArray[0],
		LS2:   lsArray[1],
	}
	return b, nil
}

func IsValidBall(arg string) bool {
	pattern := `^\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|50)\b(,\b([1-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|50)\b){0,4}$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}

func IsValidStars(arg string) bool {
	pattern := `^\b([1-9]|1[0-2])\b(,\b([1-9]|1[0-2])\b){0,1}$`
	matched, err := regexp.MatchString(pattern, arg)
	if err != nil {
		log.Println(err)
	}
	return matched
}
