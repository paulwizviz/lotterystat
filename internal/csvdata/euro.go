package csvdata

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Euro is structure representing a draw of National Lottery Euromillion
type EuroDraw struct {
	Log  map[string]string
	Err  error
	Item struct {
		DrawDate   time.Time
		DayOfWeek  time.Weekday
		Ball1      uint8
		Ball2      uint8
		Ball3      uint8
		Ball4      uint8
		Ball5      uint8
		LS1        uint8
		LS2        uint8
		UKMarker   string
		EuroMarker string
		DrawNo     uint64
	}
}

func ProcessEuroCVS(r io.Reader) chan EuroDraw {
	c := make(chan EuroDraw)
	go func() {
		cr := csv.NewReader(r)
		cr.Read() // remove titles
		ln := 1
		for {
			ln++
			ecs := EuroDraw{}
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
			ecs.Item = struct {
				DrawDate   time.Time
				DayOfWeek  time.Weekday
				Ball1      uint8
				Ball2      uint8
				Ball3      uint8
				Ball4      uint8
				Ball5      uint8
				LS1        uint8
				LS2        uint8
				UKMarker   string
				EuroMarker string
				DrawNo     uint64
			}{
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
