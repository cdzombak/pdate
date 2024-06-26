package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/itlightning/dateparse"
	"github.com/oklog/ulid/v2"
)

// Parse attempts to parse the given string into a UTC time.Time.
func Parse(val string) (time.Time, error) {
	if ui, err := ulid.Parse(val); err == nil {
		return ulid.Time(ui.Time()).UTC(), nil
	}

	if unixTs, err := strconv.ParseInt(val, 10, 64); err == nil {
		result := time.Unix(unixTs, 0)
		if result.Year() > 2070 {
			result = time.Unix(0, unixTs*int64(time.Millisecond))
		}
		if result.Year() < 1970 {
			result = time.Unix(0, unixTs)
		}
		return result.UTC(), nil
	}

	if _, err := strconv.ParseFloat(val, 64); err == nil {
		split := strings.Split(val, ".")
		if len(split) != 2 {
			// should never happen, as we know this is parseable as a float
			panic(err)
		}
		secondsStr := split[0]
		subsecondsStr := split[1]
		if len(subsecondsStr) > 9 {
			subsecondsStr = subsecondsStr[:9]
		} else if len(subsecondsStr) < 9 {
			subsecondsStr = fmt.Sprintf("%s%s", subsecondsStr, strings.Repeat("0", 9-len(subsecondsStr)))
		}
		seconds, err := strconv.ParseInt(secondsStr, 10, 64)
		if err != nil {
			// should never happen, as we know this is parseable as a float
			panic(err)
		}
		nanos, err := strconv.ParseInt(subsecondsStr, 10, 64)
		if err != nil {
			// should never happen, as we know this is parseable as a float
			panic(err)
		}

		result := time.Unix(seconds, nanos)
		return result.UTC(), nil
	}

	if result, err := dateparse.ParseStrict(val); err == nil {
		return result.UTC(), err
	}

	valUpper := strings.ToUpper(val)
	if result, err := dateparse.ParseStrict(valUpper); err == nil {
		return result.UTC(), err
	}

	var result time.Time
	if err := result.UnmarshalText([]byte(val)); err == nil {
		return result.UTC(), err
	}

	if err := result.UnmarshalText([]byte(valUpper)); err == nil {
		return result.UTC(), err
	}

	return time.Time{}, errors.New("failed to parse date string")
}
