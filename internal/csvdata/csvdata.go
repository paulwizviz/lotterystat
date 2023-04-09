// Package csvdata contains implementations to support operations related to the processing
// of CSV data source.
package csvdata

import (
	"errors"
)

// URLs to UK National lottery csv data sources
const (
	EuroURL     = "https://www.national-lottery.co.uk/results/euromillions/draw-history/csv"
	Set4LifeURL = "https://www.national-lottery.co.uk/results/set-for-life/draw-history/csv"
)

// Error types
var (
	ErrCSVURL             = errors.New("unable to access url")
	ErrResponseBody       = errors.New("unable to extract response body")
	ErrCSVLine            = errors.New("invalid csv line")
	ErrInvalidNumFmt      = errors.New("invalid number format")
	ErrInvalidDaysInMonth = errors.New("invalid days in a month")
	ErrInvalidMonth       = errors.New("invalid month")
	ErrInvalidDrawDigit   = errors.New("invalid draw digit")
)

// Logging key
const (
	CSVLogKeyLineNo = "CSV line no"
)
