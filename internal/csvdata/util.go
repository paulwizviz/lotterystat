package csvdata

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func parseDateTime(dt string) (time.Time, error) {
	elm := strings.Split(dt, "-")

	day, err := strconv.Atoi(elm[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("improper day format - %w", ErrInvalidNumFmt)
	}

	year, err := strconv.Atoi(elm[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("improper year format - %w", ErrInvalidNumFmt)
	}

	var mth time.Month
	switch elm[1] {
	case "Jan":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d January %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.January
	case "Feb":
		if year%4 == 0 {
			if day < 1 || day > 29 {
				return time.Time{}, fmt.Errorf("%d February leap year %d - %w", day, year, ErrInvalidDaysInMonth)
			}
		} else {
			if day < 1 || day > 28 {
				return time.Time{}, fmt.Errorf("%d February %d - %w", day, year, ErrInvalidDaysInMonth)
			}
		}
		mth = time.February
	case "Mar":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d Feb %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.March
	case "Apr":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%d April %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.April
	case "May":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d May %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.May
	case "Jun":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%d June %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.June
	case "Jul":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d July %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.July
	case "Aug":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d August %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.August
	case "Sep":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%d September %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.September
	case "Oct":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d October %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.October
	case "Nov":
		if day < 1 || day > 30 {
			return time.Time{}, fmt.Errorf("%d November %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.November
	case "Dec":
		if day < 1 || day > 31 {
			return time.Time{}, fmt.Errorf("%d December %d - %w", day, year, ErrInvalidDaysInMonth)
		}
		mth = time.December
	default:
		return time.Time{}, fmt.Errorf("incorrect month - %w", ErrInvalidMonth)
	}

	tm := time.Date(year, mth, day, 0, 0, 0, 0, time.UTC)

	return tm, nil
}

type drawTypeConstraint interface {
	EuroDraw | Set4LifeDraw
}

func sqliteTags[T drawTypeConstraint](typ *T) map[string]string {
	ev := reflect.Indirect(reflect.ValueOf(typ))
	tags := make(map[string]string)
	for i := 0; i < ev.Type().NumField(); i++ {
		fn := ev.Type().Field(i).Name
		tag := ev.Type().Field(i).Tag
		tItems := strings.Split(string(tag), " ")
		for _, ti := range tItems {
			if strings.Contains(ti, "sqlite") {
				si := strings.Split(ti, ":")
				tags[fn] = si[1][1 : len(si[1])-1]
			}
		}
	}
	return tags
}
