package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/xeonx/timeago"
)

var version = "<dev>"

func main() {
	var val string

	if len(os.Args) < 2 {
		fmt.Printf("Enter datetime string to parse: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		fmt.Println("")
		if err != nil {
			panic(err)
		}
		val = strings.TrimSuffix(input, "\n")
	} else {
		val = os.Args[1]

		if val == "-h" || val == "help" || val == "--help" || val == "-v" || val == "version" || val == "--version" {
			fmt.Printf("dateutil %s\n", version)
			fmt.Println("usage: dateutil [datetime string]")
			os.Exit(0)
		}
	}

	parsed, err := dateparse.ParseStrict(val)
	if err != nil {
		fmt.Printf("failed to parse input: %s\n", err)
		os.Exit(1)
	}

	localLoc, err := time.LoadLocation("Local")
	if err != nil {
		panic(fmt.Sprintf("failed to load location 'Local': %s", err))
	}
	parsedLocal := parsed.In(localLoc)

	utcLoc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(fmt.Sprintf("failed to load location 'UTC': %s", err))
	}
	parsedUTC := parsed.In(utcLoc)

	fmt.Printf(" input:\t%s\n", val)
	fmt.Printf("parsed:\t%s\n", parsed.Format("2006-01-02 15:04:05.1234 MST"))
	fmt.Println("       \t(verify this matches your input)")
	fmt.Println("")
	fmt.Printf("   UTC:\t%s\n", parsedUTC.Format("2006-01-02 3:04:05 PM"))
	fmt.Printf("       \t%s\n", parsedUTC.Format("2006-01-02T15:04:05.1234Z"))
	fmt.Println("")
	fmt.Printf(" local:\t%s\n", parsedLocal.Format("2006-01-02 3:04:05 PM MST"))
	fmt.Printf("\t(%s)\n", timeago.English.Format(parsed))
}
