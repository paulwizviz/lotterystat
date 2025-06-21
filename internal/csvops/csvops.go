package csvops

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// CSV File
	ErrDownloadFromURL = errors.New("unable to download from url")
	ErrInvalidURL      = errors.New("invalid url")
	// Date
	ErrInvalidDateFmt     = errors.New("invalid date format")
	ErrInvalidDayFmt      = errors.New("invalid day format")
	ErrInvalidDaysInMonth = errors.New("invalid day in the month")
	ErrInvalidMonth       = errors.New("invalid month")
	ErrInvalidYearFmt     = errors.New("invalid year format")
	// Content
	ErrContentMissing   = errors.New("empty csv content")
	ErrLine             = errors.New("unable to process line")
	ErrInvalidDrawDigit = errors.New("invalid draw digit")
	ErrInvalidDrawRange = errors.New("draw out of range")
	ErrInvalidDrawSeq   = errors.New("invalid draw seq")
)

func ParseDate(date string) (time.Time, error) {
	regex := regexp.MustCompile(`(?i)^\d{1,2}-(?:Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)-\d{4}$`)
	if !regex.MatchString(date) {
		return time.Time{}, fmt.Errorf("%w: invalid date format", ErrInvalidDateFmt)
	}
	elm := strings.Split(date, "-")

	day, err := strconv.Atoi(elm[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: improper day format", ErrInvalidDayFmt)
	}

	year, err := strconv.Atoi(elm[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: improper year format", ErrInvalidYearFmt)
	}

	var mth time.Month
	switch strings.ToLower(elm[1]) {
	case "jan":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d January %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.January
	case "feb":
		if year%4 == 0 {
			if day < 1 || day > 29 {
				return time.Time{}, fmt.Errorf("%w: %d February leap year %d", ErrInvalidDaysInMonth, day, year)
			}
		} else {
			if day < 1 || day > 28 {
				return time.Time{}, fmt.Errorf("%w: %d February %d", ErrInvalidDaysInMonth, day, year)
			}
		}
		mth = time.February
	case "mar":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d Feb %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.March
	case "apr":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d April %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.April
	case "may":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d May %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.May
	case "jun":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d June %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.June
	case "jul":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d July %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.July
	case "aug":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d August %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.August
	case "sep":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d September %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.September
	case "oct":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d October %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.October
	case "nov":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d November %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.November
	case "dec":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d December %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.December
	default:
		return time.Time{}, fmt.Errorf("%w: incorrect month", ErrInvalidMonth)
	}

	tm := time.Date(year, mth, day, 0, 0, 0, 0, time.UTC)

	return tm, nil
}

func ParseDrawNum(value string, maxval int) (uint8, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%w-%s", ErrInvalidDrawDigit, err.Error())
	}
	if result < 1 {
		return 0, fmt.Errorf("%w-%s", ErrInvalidDrawRange, fmt.Sprintf("got %v max %v", result, maxval))
	}
	if result > maxval {
		return 0, fmt.Errorf("%w-%s", ErrInvalidDrawRange, fmt.Sprintf("got %v max %v", result, maxval))
	}
	return uint8(result), nil
}

func ParseDrawSeq(value string) (uint64, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%w-%s", ErrInvalidDrawSeq, err.Error())
	}
	if result < 0 {
		return 0, fmt.Errorf("%w-Less than 0", ErrInvalidDrawSeq)
	}
	return uint64(result), nil
}

func DownloadFrom(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

type CSVRec struct {
	Record []string
	Line   uint
	Err    error
}

func ExtractRec(ctx context.Context, r io.Reader) chan CSVRec {
	c := make(chan CSVRec)
	go func(ch chan CSVRec) {
		defer close(ch)
		csvr := csv.NewReader(r)
		csvr.Read() // Remove header
		ln := uint(1)
	loop:
		for {
			select {
			case <-ctx.Done():
				break loop
			default:
				ln++
				rec, err := csvr.Read()
				if errors.Is(err, io.EOF) {
					break loop
				}
				if err != nil {
					ch <- CSVRec{
						Record: rec,
						Line:   ln,
						Err:    fmt.Errorf("%w-%s", ErrLine, err.Error()),
					}
					continue loop
				}
				ch <- CSVRec{
					Record: rec,
					Line:   ln,
					Err:    nil,
				}
			}
		}
	}(c)
	return c
}
