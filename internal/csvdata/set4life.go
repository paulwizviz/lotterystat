package csvdata

import "time"

// Set4Life is a structure representing a draw for Set For Life Game
type Set4Life struct {
	DrawDate   time.Time
	DayOfWeek  time.Weekday
	Ball1      uint8
	Ball2      uint8
	Ball3      uint8
	Ball4      uint8
	Ball5      uint8
	LifeBall   uint8
	BallSet    string
	Machine    string
	DrawNumber int64
}
