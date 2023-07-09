// Package euro implements datasvc draws
package result

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	ErrInvalidDayFmt      = errors.New("invalid day format")
	ErrInvalidYearFmt     = errors.New("invalid year format")
	ErrInvalidDaysInMonth = errors.New("invalid day in the month")
	ErrInvalidMonth       = errors.New("invalid month")
	ErrCSVLine            = errors.New("unable to process line")
	ErrInvalidDrawDigit   = errors.New("invalid draw digit")
	ErrInvalidDrawSeq     = errors.New("invalid draw seq")
	ErrDownloadFromURL    = errors.New("unable to download from url")
	ErrContentMissing     = errors.New("empty csv content")
)

// EuroDraw represents a line from euro draw results
type EuroDraw struct {
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
func EuroCSVFromURL(ctx context.Context) (<-chan EuroDraw, error) {
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
	c := make(chan EuroDraw)
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

// Set4LifeDraw represents a draw from Set for Life
type Set4LifeDraw struct {
	DrawDate  time.Time    `json:"draw_date" sqlite:"draw_date,INTEGER"`
	DayOfWeek time.Weekday `json:"day_of_week" sqlite:"day_of_week,INTEGER"`
	Ball1     uint8        `json:"ball1" sqlite:"ball1,INTEGER"`
	Ball2     uint8        `json:"ball2" sqlite:"ball2,INTEGER"`
	Ball3     uint8        `json:"ball3" sqlite:"ball3,INTEGER"`
	Ball4     uint8        `json:"ball4" sqlite:"ball4,INTEGER"`
	Ball5     uint8        `json:"ball5" sqlite:"ball5,INTEGER"`
	LifeBall  uint8        `json:"life_ball" sqlite:"life_ball,INTEGER"`
	BallSet   string       `json:"ball_set" sqlite:"ball_set,TEXT"`
	Machine   string       `json:"machine" sqlite:"machine,TEXT"`
	DrawNo    uint64       `json:"draw_no" sqlite:"draw_no,TEXT"`
}
