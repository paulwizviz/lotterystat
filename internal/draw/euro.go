package draw

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Euro represents a line from euro draw results
type Euro struct {
	DrawDate   time.Time    `json:"draw_date" sqlite:"draw_date,INTEGER"`
	DayOfWeek  time.Weekday `json:"day_of_week" sqlite:"day_of_week,INTEGER"`
	Ball1      uint8        `json:"ball1" sqlite:"ball1,INTEGER"`
	Ball2      uint8        `json:"ball2" sqlite:"ball2,INTEGER"`
	Ball3      uint8        `json:"ball3" sqlite:"ball3,INTEGER"`
	Ball4      uint8        `json:"ball4" sqlite:"ball4,INTEGER"`
	Ball5      uint8        `json:"ball5" sqlite:"ball5,INTEGER"`
	LS1        uint8        `json:"ls1" sqlite:"ls1,INTEGER"`
	LS2        uint8        `json:"ls2" sqlite:"ls2,INTEGER"`
	UKMarker   string       `json:"uk_marker" sqlite:"uk_marker,TEXT"`
	EuroMarker string       `json:"euro_marker" sqlite:"euro_marker,TEXT"`
	DrawNo     uint64       `json:"draw_no" sqlite:"draw_no,INTEGER"`
}

// EuroCSVFromURL implements function to download draw results from source url
func EuroCSVFromURL(ctx context.Context) (<-chan Euro, error) {
	url := "https://www.national-lottery.co.uk/results/euromillions/draw-history/csv"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w: %s %s", ErrDownloadFromURL, url, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrContentMissing, err.Error())
	}
	c := make(chan Euro)
	ec := processEuroCSV(ctx, bytes.NewReader(body))
	go func() {
		for e := range ec {
			if e.Err != nil {
				log.Println(e.Err)
			} else {
				c <- e.Draw
			}
		}
		close(c)
	}()
	return c, nil
}

type euroChan struct {
	Draw Euro
	Err  error
}

func processEuroCSV(ctx context.Context, r io.Reader) <-chan euroChan {
	c := make(chan euroChan)
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
					c <- euroChan{
						Draw: Euro{},
						Err:  err,
					}
					continue loop
				}
				drawDate, err := parseDateTime(rec[0])
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b1, err := parseDrawNum(rec[1], 50)
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b2, err := parseDrawNum(rec[2], 50)
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b3, err := parseDrawNum(rec[3], 50)
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b4, err := parseDrawNum(rec[4], 50)
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b5, err := parseDrawNum(rec[5], 50)
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				ls1, err := parseDrawNum(rec[6], 12)
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				ls2, err := parseDrawNum(rec[7], 12)
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				dn, err := parseDrawSeq(rec[10])
				if err != nil {
					c <- euroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				c <- euroChan{
					Draw: Euro{
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
					},
					Err: nil,
				}
			}
		}
	}()
	return c
}