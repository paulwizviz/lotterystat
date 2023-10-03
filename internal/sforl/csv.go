package sforl

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"paulwizviz/lotterystat/internal/csvutil"
)

type DrawChan struct {
	Draw Draw
	Err  error
}

func ProcessCSV(ctx context.Context, r io.Reader) <-chan DrawChan {
	c := make(chan DrawChan)
	go func() {
		cr := csv.NewReader(r)
		cr.Read() // remove titles
		ln := 1
		defer close(c)

	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
				ln++
				rec, err := cr.Read()
				if errors.Is(err, io.EOF) {
					break loop
				}
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  err,
					}
					continue loop
				}
				drawDate, err := csvutil.ParseDateTime(rec[0])
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b1, err := csvutil.ParseDrawNum(rec[1], 50)
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b2, err := csvutil.ParseDrawNum(rec[2], 50)
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b3, err := csvutil.ParseDrawNum(rec[3], 50)
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b4, err := csvutil.ParseDrawNum(rec[4], 50)
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b5, err := csvutil.ParseDrawNum(rec[5], 50)
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				lb, err := csvutil.ParseDrawNum(rec[6], 12)
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				dn, err := csvutil.ParseDrawSeq(rec[9])
				if err != nil {
					c <- DrawChan{
						Draw: Draw{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				c <- DrawChan{
					Draw: Draw{
						DrawDate:  drawDate,
						DayOfWeek: drawDate.Weekday(),
						Ball1:     uint8(b1),
						Ball2:     uint8(b2),
						Ball3:     uint8(b3),
						Ball4:     uint8(b4),
						Ball5:     uint8(b5),
						LifeBall:  uint8(lb),
						BallSet:   rec[7],
						Machine:   rec[8],
						DrawNo:    dn,
					},
					Err: nil,
				}
			}
		}
	}()
	return c
}
