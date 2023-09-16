package draw

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	EuroURL = "https://www.national-lottery.co.uk/results/euromillions/draw-history/csv"
)

var (
	ErrCSVContentMissing     = errors.New("empty csv content")
	ErrCSVDownloadFromURL    = errors.New("unable to download from url")
	ErrCSVInvalidDayFmt      = errors.New("invalid day format")
	ErrCSVInvalidYearFmt     = errors.New("invalid year format")
	ErrCSVInvalidDaysInMonth = errors.New("invalid day in the month")
	ErrInvalidMonth          = errors.New("invalid month")
	ErrCSVLine               = errors.New("unable to process line")
	ErrCSVInvalidDrawDigit   = errors.New("invalid draw digit")
	ErrCSVInvalidDrawRange   = errors.New("draw out of range")
	ErrCSVInvalidDrawSeq     = errors.New("invalid draw seq")
	ErrCSVInvalidURL         = errors.New("invalid url")
)

func parseDateTime(dt string) (time.Time, error) {
	elm := strings.Split(dt, "-")

	day, err := strconv.Atoi(elm[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: improper day format", ErrCSVInvalidDayFmt)
	}

	year, err := strconv.Atoi(elm[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: improper year format", ErrCSVInvalidYearFmt)
	}

	var mth time.Month
	switch elm[1] {
	case "Jan":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d January %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.January
	case "Feb":
		if year%4 == 0 {
			if day < 1 || day > 29 {
				return time.Time{}, fmt.Errorf("%w: %d February leap year %d", ErrCSVInvalidDaysInMonth, day, year)
			}
		} else {
			if day < 1 || day > 28 {
				return time.Time{}, fmt.Errorf("%w: %d February %d", ErrCSVInvalidDaysInMonth, day, year)
			}
		}
		mth = time.February
	case "Mar":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d Feb %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.March
	case "Apr":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d April %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.April
	case "May":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d May %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.May
	case "Jun":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d June %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.June
	case "Jul":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d July %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.July
	case "Aug":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d August %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.August
	case "Sep":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d September %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.September
	case "Oct":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d October %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.October
	case "Nov":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d November %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.November
	case "Dec":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d December %d", ErrCSVInvalidDaysInMonth, day, year)
		}
		mth = time.December
	default:
		return time.Time{}, fmt.Errorf("%w: incorrect month", ErrInvalidMonth)
	}

	tm := time.Date(year, mth, day, 0, 0, 0, 0, time.UTC)

	return tm, nil
}

func parseDrawNum(value string, maxval int) (uint8, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrCSVInvalidDrawDigit, err.Error())
	}
	if result < 1 {
		return 0, fmt.Errorf("%w: %s", ErrCSVInvalidDrawRange, fmt.Sprintf("got %v max %v", result, maxval))
	}
	if result > maxval {
		return 0, fmt.Errorf("%w: %s", ErrCSVInvalidDrawRange, fmt.Sprintf("got %v max %v", result, maxval))
	}
	return uint8(result), nil
}

func parseDrawSeq(value string) (uint64, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrCSVInvalidDrawSeq, err.Error())
	}
	return uint64(result), nil
}

func ProcessEuroCSV(ctx context.Context, r io.Reader) <-chan EuroChan {
	c := make(chan EuroChan)
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
					c <- EuroChan{
						Draw: Euro{},
						Err:  err,
					}
					continue loop
				}
				drawDate, err := parseDateTime(rec[0])
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b1, err := parseDrawNum(rec[1], 50)
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b2, err := parseDrawNum(rec[2], 50)
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b3, err := parseDrawNum(rec[3], 50)
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b4, err := parseDrawNum(rec[4], 50)
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				b5, err := parseDrawNum(rec[5], 50)
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				ls1, err := parseDrawNum(rec[6], 12)
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				ls2, err := parseDrawNum(rec[7], 12)
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				dn, err := parseDrawSeq(rec[10])
				if err != nil {
					c <- EuroChan{
						Draw: Euro{},
						Err:  fmt.Errorf("record on line: %d: %w", ln, err),
					}
					continue loop
				}
				c <- EuroChan{
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
