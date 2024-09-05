package tball

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"sync"
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

func DuplicateData(ctx context.Context, sqliteDB *sql.DB, psqlDB *sql.DB, numworkers int) error {

	rows, err := selectAllDrawRows(ctx, sqliteDB)
	if err != nil {
		return err
	}

	psqlStmt, err := prepInsertDrawStmt(ctx, psqlDB)
	if err != nil {
		return err
	}

	draws := selectAllDraw(rows)

	var wg sync.WaitGroup
	for w := 0; w < numworkers; w++ {
		wg.Add(1)
		go func() {
			for d := range draws {
				_, err := insertDraw(ctx, psqlStmt, d)
				if err != nil {
					log.Println(err)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return nil
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
	for i := 1; i < 40; i++ {
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

type TballCount struct {
	Num   uint8
	Count uint
}

func TballFreq(ctx context.Context, db *sql.DB) ([]TballCount, error) {
	stmt, err := prepCountTBallStmt(ctx, db)
	if err != nil {
		return nil, err
	}

	tballCounts := []TballCount{}
	for i := 1; i < 15; i++ {
		var lc TballCount
		lc.Num = uint8(i)
		count, err := countChoice(ctx, stmt, uint8(i))
		if err != nil {
			continue
		}
		lc.Count = count
		tballCounts = append(tballCounts, lc)
	}
	return tballCounts, nil
}
