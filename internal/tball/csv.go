package tball

import (
	"errors"
	"fmt"
	"sync"

	"github.com/paulwizviz/lotterystat/internal/csvops"
)

func ProcessCSV(recs chan csvops.CSVRec, numWorkers int) []DrawChan {
	result := make(chan DrawChan)
	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go func() {
			csvWorker(recs, result)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	drawChans := []DrawChan{}
	for r := range result {
		drawChans = append(drawChans, r)
	}
	return drawChans
}

func csvWorker(jobs chan csvops.CSVRec, results chan DrawChan) {
	for j := range jobs {
		drawChan := DrawChan{}
		if errors.Is(j.Err, csvops.ErrLine) {
			drawChan.Draw = Draw{}
			drawChan.Err = ErrRec
			results <- drawChan
			continue
		}
		draw, err := processRecord(j.Record)
		if err != nil {
			drawChan.Draw = Draw{}
			drawChan.Err = err
			results <- drawChan
			continue
		}
		drawChan.Draw = draw
		results <- drawChan
	}
}

func processRecord(rec []string) (Draw, error) {
	maxValue := 39
	draw := Draw{}

	dt, err := csvops.ParseDate(rec[0])
	if err != nil {
		return draw, fmt.Errorf("%w-%v", ErrDrawDate, err)
	}
	draw.DrawDate = dt
	draw.DayOfWeek = dt.Weekday()

	ball1, err := csvops.ParseDrawNum(rec[1], maxValue)
	if err != nil {
		return draw, fmt.Errorf("%w-%v", ErrBall1, err)
	}
	draw.Ball1 = ball1

	ball2, err := csvops.ParseDrawNum(rec[2], maxValue)
	if err != nil {
		return draw, fmt.Errorf("%w-%v", ErrBall2, err)
	}
	draw.Ball2 = ball2

	ball3, err := csvops.ParseDrawNum(rec[3], maxValue)
	if err != nil {
		return draw, fmt.Errorf("%w-%v", ErrBall3, err)
	}
	draw.Ball3 = ball3

	ball4, err := csvops.ParseDrawNum(rec[4], maxValue)
	if err != nil {
		return draw, fmt.Errorf("%w-%v", ErrBall4, err)
	}
	draw.Ball4 = ball4

	ball5, err := csvops.ParseDrawNum(rec[5], maxValue)
	if err != nil {
		return draw, fmt.Errorf("%w-%v", ErrBall5, err)
	}
	draw.Ball5 = ball5

	tball, err := csvops.ParseDrawNum(rec[6], 14)
	if err != nil {
		return draw, fmt.Errorf("%w-%v", ErrTBall, err)
	}
	draw.TBall = tball

	draw.BallSet = rec[7]
	draw.Machine = rec[8]

	seq, err := csvops.ParseDrawSeq(rec[9])
	if err != nil {
		return draw, fmt.Errorf("%w-%v", ErrSeq, err)
	}
	draw.DrawNo = seq
	return draw, nil
}
