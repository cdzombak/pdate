//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"math"
	"syscall/js"
	"time"
)

// parseDate is the WASM-exported function that parses a datetime string
func parseDate(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return map[string]interface{}{
			"error": "expected 1 argument: datetime string",
		}
	}

	input := args[0].String()
	parsed, err := Parse(input)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	localLoc, err := time.LoadLocation("Local")
	if err != nil {
		localLoc = time.UTC
	}
	parsedLocal := parsed.In(localLoc)

	// Get timezone name and offset
	zoneName, offset := parsedLocal.Zone()
	offsetHours := offset / 3600
	offsetMins := (offset % 3600) / 60
	timezoneStr := fmt.Sprintf("%s (UTC%+d:%02d)", zoneName, offsetHours, offsetMins)

	result := map[string]interface{}{
		"success": true,
		"input":   input,
		"parsed":  parsed.Format("2006-01-02 15:04:05 MST"),
		"utc": map[string]interface{}{
			"formatted":     parsed.Format("2006-01-02 3:04:05 PM"),
			"rfc3339":       parsed.Format(time.RFC3339),
			"rfc1123z":      parsed.Format(time.RFC1123Z),
			"unix":          parsed.Unix(),
			"unixMilli":     parsed.UnixMilli(),
			"unixNano":      parsed.UnixNano(),
			"iso8601":       parsed.Format("2006-01-02T15:04:05Z"),
		},
		"local": map[string]interface{}{
			"formatted": parsedLocal.Format("2006-01-02 3:04:05 PM MST"),
			"timezone":  timezoneStr,
		},
	}

	// Add relative time
	relativeTime := formatRelativeTime(parsed)
	result["relative"] = map[string]interface{}{
		"formatted": relativeTime,
	}

	return result
}

// formatRelativeTime formats a time as a relative string (e.g., "5 minutes ago")
// This matches the logic from timeago_config.go
func formatRelativeTime(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	const (
		second = time.Second
		minute = time.Minute
		hour   = time.Hour
		day    = 24 * time.Hour
		month  = 30 * day
		year   = 365 * day
	)

	// Maximum relative time is 13 months
	maxDuration := 13 * month
	if diff > maxDuration || diff < -maxDuration {
		return t.Format("2006-01-02 15:04:05 MST")
	}

	isPast := diff >= 0
	if !isPast {
		diff = -diff
	}

	var value int
	var unit string

	switch {
	case diff < minute:
		value = int(diff / second)
		if value <= 1 {
			unit = "about a second"
		} else {
			return formatTimeString(value, "seconds", isPast)
		}
	case diff < hour:
		value = int(diff / minute)
		if value == 1 {
			unit = "about a minute"
		} else {
			return formatTimeString(value, "minutes", isPast)
		}
	case diff < day:
		value = int(diff / hour)
		if value == 1 {
			unit = "about an hour"
		} else {
			return formatTimeString(value, "hours", isPast)
		}
	case diff < month:
		value = int(diff / day)
		if value == 1 {
			unit = "one day"
		} else {
			return formatTimeString(value, "days", isPast)
		}
	case diff < year:
		value = int(math.Round(float64(diff) / float64(month)))
		if value == 1 {
			unit = "one month"
		} else {
			return formatTimeString(value, "months", isPast)
		}
	default:
		value = int(math.Round(float64(diff) / float64(year)))
		if value == 1 {
			unit = "one year"
		} else {
			return formatTimeString(value, "years", isPast)
		}
	}

	if isPast {
		return unit + " ago"
	}
	return "in " + unit
}

func formatTimeString(value int, unit string, isPast bool) string {
	result := fmt.Sprintf("%d %s", value, unit)
	if isPast {
		return result + " ago"
	}
	return "in " + result
}

func main() {
	// Keep the program running
	c := make(chan struct{})

	// Register the parseDate function to be callable from JavaScript
	js.Global().Set("pdateParse", js.FuncOf(parseDate))

	fmt.Println("pdate WASM module loaded")

	<-c
}
