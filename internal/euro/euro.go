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

func PersistsCSV(ctx context.Context, db *sql.DB, nworkers int) error {
	r, err := csvutil.DownloadFrom(CSVUrl)
	if err != nil {
		return err
	}
	ch := ProcessCSV(ctx, r)
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

func persistsDrawChan(ctx context.Context, db *sql.DB, dc <-chan DrawChan) error {
	for c := range dc {
		stmt, err := prepareInsertDrawStmt(ctx, db)
		if err != nil {
			return err
		}
		_, err = insertDraw(ctx, stmt, c.Draw)
		if err != nil {
			return err
		}
	}
	return nil
}
