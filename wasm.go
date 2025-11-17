//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/xeonx/timeago"
)

// WebTimeAgo is the web-specific timeago config with no maximum duration
var WebTimeAgo = timeago.Config{
	PastPrefix:   "",
	PastSuffix:   " ago",
	FuturePrefix: "in ",
	FutureSuffix: "",

	Periods: []timeago.FormatPeriod{
		{D: time.Second, One: "about a second", Many: "%d seconds"},
		{D: time.Minute, One: "about a minute", Many: "%d minutes"},
		{D: time.Hour, One: "about an hour", Many: "%d hours"},
		{D: timeago.Day, One: "one day", Many: "%d days"},
		{D: timeago.Month, One: "one month", Many: "%d months"},
		{D: timeago.Year, One: "one year", Many: "%d years"},
	},

	Zero: "about a second",
	Max:  100 * timeago.Year, // Effectively unlimited - always show relative time
}

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

	// Add relative time using web-specific config (no 13-month limit)
	relativeTime := WebTimeAgo.Format(parsed)
	result["relative"] = map[string]interface{}{
		"formatted": relativeTime,
	}

	return result
}

func main() {
	// Keep the program running
	c := make(chan struct{})

	// Register the parseDate function to be callable from JavaScript
	js.Global().Set("pdateParse", js.FuncOf(parseDate))

	fmt.Println("pdate WASM module loaded")

	<-c
}
