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

// CSV Processing Errors
var (
	ErrCSVURL             = errors.New("unable to access url")
	ErrResponseBody       = errors.New("unable to extract response body")
	ErrCSVLine            = errors.New("invalid csv line")
	ErrInvalidNumFmt      = errors.New("invalid number format")
	ErrInvalidDaysInMonth = errors.New("invalid days in a month")
	ErrInvalidMonth       = errors.New("invalid month")
	ErrInvalidDrawDigit   = errors.New("invalid draw digit")
)

// DB related errors
var (
	ErrUnableToCreateTable = errors.New("unable to create table")
	ErrDuplicateEntry      = errors.New("duplicate data entry")
)

// Logging key
const (
	CSVLogKeyLineNo = "CSV line no"
)

// Struct tag represents Go struct tag information obtained from
// reflection of a struct
type StructTag struct {
	FieldName string
	Tag       string
}

// DrawType represents a collection of type constraints
type DrawType interface {
	EuroDraw | Set4LifeDraw
}
