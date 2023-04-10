package csvdata

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

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

func (e *EuroDraw) SQLiteTags() map[string]string {
	return sqliteTags(e)
}

// Euro is structure representing a draw of National Lottery Euromillion
type EuroDrawSig struct {
	Log  map[string]string
	Err  error
	Item EuroDraw
}

func ProcessEuroCVS(r io.Reader) <-chan EuroDrawSig {
	c := make(chan EuroDrawSig)
	go func() {
		cr := csv.NewReader(r)
		cr.Read() // remove titles
		ln := 1
		for {
			ln++
			ecs := EuroDrawSig{}
			ecs.Log = map[string]string{
				CSVLogKeyLineNo: fmt.Sprintf("%d", ln),
			}
			rec, err := cr.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrCSVLine)
				c <- ecs
				continue
			}
			drawDate, err := parseDateTime(rec[0])
			if err != nil {
				ecs.Err = fmt.Errorf("%w", err)
				c <- ecs
				continue
			}
			b1, err := strconv.Atoi(rec[1])
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- ecs
				continue
			}
			b2, err := strconv.Atoi(rec[2])
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- ecs
				continue
			}
			b3, err := strconv.Atoi(rec[3])
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- ecs
				continue
			}
			b4, err := strconv.Atoi(rec[4])
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- ecs
				continue
			}
			b5, err := strconv.Atoi(rec[5])
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- ecs
				continue
			}
			ls1, err := strconv.Atoi(rec[6])
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- ecs
				continue
			}
			ls2, err := strconv.Atoi(rec[7])
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- ecs
				continue
			}
			dn, err := strconv.Atoi(rec[10])
			if err != nil {
				ecs.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- ecs
				continue
			}
			ecs.Item = EuroDraw{
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
				DrawNo:     uint64(dn),
			}
			c <- ecs
		}
		close(c)
	}()

	return c
}
