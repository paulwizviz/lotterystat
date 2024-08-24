// Package sforl represents data structures for Set For Life draw
package sforl

import (
	"context"
	"database/sql"
	"fmt"
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

type MainCount struct {
	Num   uint8
	Count uint
}

func MainFreq(ctx context.Context, db *sql.DB) ([]MainCount, error) {
	stmt, err := prepCountBallStmt(ctx, db)
	if err != nil {
		return nil, err
	}

	mainCounts := []MainCount{}
	for i := 1; i < 48; i++ {
		var bc MainCount
		bc.Num = uint8(i)
		count, err := countChoice(ctx, stmt, uint8(i))
		if err != nil {
			continue
		}
		bc.Count = count
		mainCounts = append(mainCounts, bc)
	}
	return mainCounts, nil
}

type StarCount struct {
	Num   uint8
	Count uint
}

func StarFreq(ctx context.Context, db *sql.DB) ([]StarCount, error) {
	stmt, err := prepCountLuckyStmt(ctx, db)
	if err != nil {
		return nil, err
	}

	starCounts := []StarCount{}
	for i := 1; i < 11; i++ {
		var sc StarCount
		sc.Num = uint8(i)
		count, err := countChoice(ctx, stmt, uint8(i))
		if err != nil {
			continue
		}
		sc.Count = count
		starCounts = append(starCounts, sc)
	}
	return starCounts, nil
}

type TwoCombo struct {
	Combo [2]uint8
	Count uint
}

func twoMainComboFreqWorker(ctx context.Context, stmt *sql.Stmt, input <-chan [2]uint8, result chan<- TwoCombo) {
	for i := range input {
		tc := TwoCombo{}
		count, err := countTwoMain(ctx, stmt, i[0], i[1])
		if err != nil {
			continue
		}
		tc.Combo = i
		tc.Count = count
		result <- tc
	}
}

func twoMainComboFreq(ctx context.Context, stmt *sql.Stmt, numworkers int) []TwoCombo {

	numOfJob := 1081
	set := []uint8{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47,
	}

	jobs := make(chan [2]uint8, numOfJob)
	output := make(chan TwoCombo, numOfJob)

	for i := 0; i < numworkers; i++ {
		go twoMainComboFreqWorker(ctx, stmt, jobs, output)
	}

	for i := 0; i < len(set); i++ {
		for j := i + 1; j < len(set); j++ {
			d := [2]uint8{}
			d[0], d[1] = set[i], set[j]
			jobs <- d
		}
	}

	results := []TwoCombo{}
	for i := 0; i < numOfJob; i++ {
		results = append(results, <-output)
	}
	return results
}

func TwoMainComboFreq(ctx context.Context, db *sql.DB, numworkers int) []TwoCombo {
	stmt, err := prepTwoMainStmt(context.TODO(), db)
	if err != nil {
		fmt.Printf("Two main statement error: %v", err)
	}
	defer stmt.Close()
	return twoMainComboFreq(ctx, stmt, numworkers)
}

func DuplicateData(ctx context.Context, sqliteDB *sql.DB, psqlDB *sql.DB) error {

	rows, err := selectAllDrawRows(ctx, sqliteDB)
	if err != nil {
		return err
	}

	psqlStmt, err := prepInsertDrawStmt(ctx, psqlDB)
	if err != nil {
		return err
	}

	draws := selectAllDraw(rows)

	for d := range draws {
		_, err := insertDraw(ctx, psqlStmt, d)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
