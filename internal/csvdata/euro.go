package csvdata

import (
	"encoding/csv"
	"io"
	"time"
)

// Euro is structure representing a draw of National Lottery Euromillion
type Euro struct {
	DrawDate  time.Time
	DayOfWeek time.Weekday
	Ball1     uint8
	Ball2     uint8
	Ball3     uint8
	Ball4     uint8
	Ball5     uint8
	LS1       uint8
	LS2       uint8
	Marker    string
}

type euroChanSignal struct {
	Err     error
	Content Euro
}

func processEuroCVS(r io.Reader) chan euroChanSignal {
	c := make(chan euroChanSignal)
	go func() {
		cr := csv.NewReader(r)
		for {
			ecs := euroChanSignal{}
			_, err := cr.Read()
			if err != nil {
				ecs.Err = err
			}
			if err == io.EOF {
				break
			}
			ecs.Content = Euro{}
			c <- ecs
		}
		close(c)
	}()

	return c
}
