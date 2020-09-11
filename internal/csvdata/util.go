package csvdata

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseDateTime(dt string) (time.Time, error) {
	elm := strings.Split(dt, "-")

	day, err := strconv.Atoi(elm[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("improper day string - %w", err)
	}

	year, err := strconv.Atoi(elm[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("improper year string - %w", err)
	}

	var mth time.Month
	switch elm[1] {
	case "Jan":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d January %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.January
	case "Feb":
		if year%4 == 0 {
			if day < 1 || day > 29 {
				return time.Time{}, fmt.Errorf("%d February %d - %w", day, year, errInvalidDaysInMonth)
			}
		} else {
			if day < 1 || day > 28 {
				return time.Time{}, fmt.Errorf("%d February %d - %w", day, year, errInvalidDaysInMonth)
			}
		}
		mth = time.February
	case "Mar":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d Feb %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.March
	case "Apr":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%d April %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.April
	case "May":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d May %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.May
	case "Jun":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%d June %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.June
	case "Jul":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d July %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.July
	case "Aug":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d August %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.August
	case "Sep":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%d September %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.September
	case "Oct":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d October %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.October
	case "Nov":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%d November %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.November
	case "Dec":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d December %d - %w", day, year, errInvalidDaysInMonth)
		}
		mth = time.December
	default:
		return time.Time{}, fmt.Errorf("incorrect month - %w", errInvalidMonth)
	}

	tm := time.Date(year, mth, day, 0, 0, 0, 0, time.UTC)

	return tm, nil
}
