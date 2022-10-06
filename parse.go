package main

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

func Parse(val string) (time.Time, error) {
	if unixTs, err := strconv.ParseInt(val, 10, 64); err == nil {
		result := time.Unix(unixTs, 0)
		if result.Year() > 2070 {
			result = time.Unix(0, unixTs*int64(time.Millisecond))
		}
		if result.Year() < 1970 {
			result = time.Unix(0, unixTs)
		}
		return result, nil
	}

	if result, err := dateparse.ParseStrict(val); err == nil {
		return result, err
	}

	valUpper := strings.ToUpper(val)
	if result, err := dateparse.ParseStrict(valUpper); err != nil {
		return result, err
	}

	var result time.Time
	if err := result.UnmarshalText([]byte(val)); err == nil {
		return result, err
	}

	if err := result.UnmarshalText([]byte(valUpper)); err == nil {
		return result, err
	}

	return time.Time{}, errors.New("failed to parse date string")
}
