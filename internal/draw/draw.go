// Package euro implements datasvc draws
package draw

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidDayFmt      = errors.New("invalid day format")
	ErrInvalidYearFmt     = errors.New("invalid year format")
	ErrInvalidDaysInMonth = errors.New("invalid day in the month")
	ErrInvalidMonth       = errors.New("invalid month")
	ErrCSVLine            = errors.New("unable to process line")
	ErrInvalidDrawDigit   = errors.New("invalid draw digit")
	ErrInvalidDrawRange   = errors.New("draw out of range")
	ErrInvalidDrawSeq     = errors.New("invalid draw seq")
	ErrDownloadFromURL    = errors.New("unable to download from url")
	ErrContentMissing     = errors.New("empty csv content")
)

func parseDateTime(dt string) (time.Time, error) {
	elm := strings.Split(dt, "-")

	day, err := strconv.Atoi(elm[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: improper day format", ErrInvalidDayFmt)
	}

	year, err := strconv.Atoi(elm[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: improper year format", ErrInvalidYearFmt)
	}

	var mth time.Month
	switch elm[1] {
	case "Jan":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d January %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.January
	case "Feb":
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
	case "Mar":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d Feb %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.March
	case "Apr":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d April %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.April
	case "May":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d May %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.May
	case "Jun":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d June %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.June
	case "Jul":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d July %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.July
	case "Aug":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d August %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.August
	case "Sep":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d September %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.September
	case "Oct":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%w: %d October %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.October
	case "Nov":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%w: %d November %d", ErrInvalidDaysInMonth, day, year)
		}
		mth = time.November
	case "Dec":
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

func parseDrawNum(value string, maxval int) (uint8, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrInvalidDrawDigit, err.Error())
	}
	if result < 1 {
		return 0, fmt.Errorf("%w: %s", ErrInvalidDrawRange, fmt.Sprintf("got %v max %v", result, maxval))
	}
	if result > maxval {
		return 0, fmt.Errorf("%w: %s", ErrInvalidDrawRange, fmt.Sprintf("got %v max %v", result, maxval))
	}
	return uint8(result), nil
}

func parseDrawSeq(value string) (uint64, error) {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrInvalidDrawSeq, err.Error())
	}
	return uint64(result), nil
}
