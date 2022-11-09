package main

import (
	"time"

	"github.com/xeonx/timeago"
)

// CustomEnglishTimeAgo is the application's configuration for displaying "time ago" strings.
var CustomEnglishTimeAgo = timeago.Config{
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

	Max:           13 * timeago.Month,
	DefaultLayout: "2006-01-02 15:04:05 MST",
}
