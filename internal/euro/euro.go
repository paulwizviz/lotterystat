// Package euro contains abstractions of the Euro Millions draw from UK National lottery and
// all related data analytics.
package euro

import (
	"context"
	"database/sql"
	"log"
	"paulwizviz/lotterystat/internal/csvutil"
	"sync"
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

var (
	CreateTable func(ctx context.Context, db *sql.DB) error                                                                                           = createTable
	MatchDraw   func(ctx context.Context, db *sql.DB, ball1 uint, ball2 uint, ball3 uint, ball4 uint, ball5 uint, ls1 uint, ls2 uint) ([]Draw, error) = matchDraw
)

func PersistsCSV(ctx context.Context, db *sql.DB, nworkers int) error {
	r, err := csvutil.DownloadFrom(CSVUrl)
	if err != nil {
		return err
	}
	ch := processCSV(ctx, r)
	var wg sync.WaitGroup
	wg.Add(nworkers)
	for i := 0; i < nworkers; i++ {
		go func() {
			err := persistsDrawChan(ctx, db, ch)
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
