package csvdata

import (
	"errors"
	"time"
)

const (
	EuroURL     = "https://www.national-lottery.co.uk/results/euromillions/draw-history/csv"
	Set4LifeURL = "https://www.national-lottery.co.uk/results/set-for-life/draw-history/csv"
)

var InvalidTime = time.Date(0001, 1, 1, 0, 0, 0, 0, time.Local)

var (
	errInvalidDaysInMonth = errors.New("invalid days in a month")
	errInvalidMonth       = errors.New("invalid month")
)
