package csvdata

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

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

func (s *Set4LifeDraw) SQLiteTags() map[string]string {
	return sqliteTags(s)
}

type Set4LifeDrawSig struct {
	Log  map[string]string
	Err  error
	Item Set4LifeDraw
}

func ProcessS4LCVS(r io.Reader) <-chan Set4LifeDrawSig {
	c := make(chan Set4LifeDrawSig)
	go func() {
		cr := csv.NewReader(r)
		cr.Read() // remove titles
		ln := 1
		for {
			ln++
			s4ld := Set4LifeDrawSig{}
			s4ld.Log = map[string]string{
				CSVLogKeyLineNo: fmt.Sprintf("%d", ln),
			}
			rec, err := cr.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				s4ld.Err = fmt.Errorf("%w-%w", err, ErrCSVLine)
				c <- s4ld
				continue
			}
			drawDate, err := parseDateTime(rec[0])
			if err != nil {
				s4ld.Err = fmt.Errorf("%w", err)
				c <- s4ld
				continue
			}
			b1, err := strconv.Atoi(rec[1])
			if err != nil {
				s4ld.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- s4ld
				continue
			}
			b2, err := strconv.Atoi(rec[2])
			if err != nil {
				s4ld.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- s4ld
				continue
			}
			b3, err := strconv.Atoi(rec[3])
			if err != nil {
				s4ld.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- s4ld
				continue
			}
			b4, err := strconv.Atoi(rec[4])
			if err != nil {
				s4ld.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- s4ld
				continue
			}
			b5, err := strconv.Atoi(rec[5])
			if err != nil {
				s4ld.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- s4ld
				continue
			}
			lb, err := strconv.Atoi(rec[6])
			if err != nil {
				s4ld.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- s4ld
				continue
			}
			dn, err := strconv.Atoi(rec[9])
			if err != nil {
				s4ld.Err = fmt.Errorf("%w-%w", err, ErrInvalidDrawDigit)
				c <- s4ld
				continue
			}
			s4ld.Item = Set4LifeDraw{
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
				DrawNo:    uint64(dn),
			}
			c <- s4ld
		}
		close(c)
	}()

	return c
}
