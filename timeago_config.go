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
		{time.Second, "about a second", "%d seconds"},
		{time.Minute, "about a minute", "%d minutes"},
		{time.Hour, "about an hour", "%d hours"},
		{timeago.Day, "one day", "%d days"},
		{timeago.Month, "one month", "%d months"},
		{timeago.Year, "one year", "%d years"},
	},

	Zero: "about a second",

	Max:           13 * timeago.Month,
	DefaultLayout: "2006-01-02 15:04:05 MST",
}
