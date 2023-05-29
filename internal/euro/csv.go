package euro

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/paulwizviz/lotterystat/internal/csvutil"
	"golang.org/x/exp/slog"
)

// DrawFromURL implements function to download draw results from source url
func DrawFromURL(ctx context.Context) (<-chan Draw, error) {
	url := "https://www.national-lottery.co.uk/results/euromillions/draw-history/csv"
	slog.Info("Dowloading draws csv from url", "url", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w: %s %s", csvutil.ErrDownloadFromURL, url, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", csvutil.ErrContentMissing, err.Error())
	}
	c := processCSV(ctx, bytes.NewReader(body))
	return c, nil
}

func processCSV(ctx context.Context, r io.Reader) <-chan Draw {
	const logMsg = "Processing CSV"
	c := make(chan Draw)
	go func() {
		cr := csv.NewReader(r)
		cr.Read() // remove titles
		ln := 1
		defer close(c)
	loop:
		for {
			select {
			case <-ctx.Done():
				slog.Info(logMsg, "context", "done")
				break loop
			default:
				ln++
				rec, err := cr.Read()
				if errors.Is(err, io.EOF) {
					break loop
				}
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				drawDate, err := csvutil.ParseDateTime(rec[0])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				b1, err := csvutil.ParseDrawNum(rec[1])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				b2, err := csvutil.ParseDrawNum(rec[2])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				b3, err := csvutil.ParseDrawNum(rec[3])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				b4, err := csvutil.ParseDrawNum(rec[4])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				b5, err := csvutil.ParseDrawNum(rec[5])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				ls1, err := csvutil.ParseDrawNum(rec[6])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				ls2, err := csvutil.ParseDrawNum(rec[7])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				dn, err := csvutil.ParseDrawSeq(rec[10])
				if err != nil {
					slog.Error(logMsg, "error", err.Error())
					continue loop
				}
				c <- Draw{
					DrawDate:   drawDate,
					DayOfWeek:  drawDate.Weekday(),
					Ball1:      uint8(b1),
					Ball2:      uint8(b2),
					Ball3:      uint8(b3),
					Ball4:      uint8(b4),
					Ball5:      uint8(b5),
					LS1:        uint8(ls1),
					LS2:        uint8(ls2),
					UKMarker:   rec[8],
					EuroMarker: rec[9],
					DrawNo:     dn,
				}
			}
		}
	}()
	return c
}
