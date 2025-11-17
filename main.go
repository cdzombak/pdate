//go:build !js && !wasm
// +build !js,!wasm

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

var version = "<dev>"

func main() {
	var val string
	var parsed time.Time
	var err error

	whiteColor := color.New(color.FgHiWhite, color.Bold)

	if len(os.Args) < 2 {
		parsed = time.Now()
	} else {
		val = os.Args[1]

		if val == "-h" || val == "-v" || strings.Contains(val, "help") || strings.Contains(val, "version") {
			_, _ = whiteColor.Printf("pdate %s\n", version)
			fmt.Println("usage: pdate [datetime|ULID string]")
			os.Exit(0)
		}

		if val == "-i" || val == "--interactive" {
			fmt.Printf("Enter datetime string to parse: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			fmt.Println("")
			if err != nil {
				panic(err)
			}
			val = strings.TrimSuffix(input, "\n")
		}

		val = strings.TrimSpace(val)
		parsed, err = Parse(val)
		if err != nil {
			fmt.Println(color.RedString("failed to parse input"))
			os.Exit(1)
		}
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

	if val != "" {
		fmt.Printf(" %s\t%s\n", whiteColor.Sprint("input:"), val)
		fmt.Printf("%s\t%s\n", whiteColor.Sprint("parsed:"), parsed.Format("2006-01-02 15:04:05 MST"))
		fmt.Println("       \t(verify this matches your input/expectations)")
	} else {
		fmt.Printf(" %s\tnow\n", whiteColor.Sprint("input:"))
	}
	fmt.Println("")

	fmt.Printf("   %s\t%s\n", whiteColor.Sprint("UTC:"), parsedUTC.Format("2006-01-02 3:04:05 PM"))
	fmt.Printf("  %s\t%s\n", whiteColor.Sprint("3339:"), parsedUTC.Format(time.RFC3339))
	fmt.Printf(" %s\t%s\n", whiteColor.Sprint("1123Z:"), parsedUTC.Format(time.RFC1123Z))
	fmt.Printf("  %s\t%d\n", whiteColor.Sprint("Unix:"), parsedUTC.Unix())
	fmt.Printf("       \t%d (millis)\n", parsedUTC.UnixMilli())
	fmt.Printf("       \t%d (nanos)\n", parsedUTC.UnixNano())
	fmt.Println("")
	fmt.Printf(" %s\t%s\n", whiteColor.Sprint("local:"), parsedLocal.Format("2006-01-02 3:04:05 PM MST"))
	if val != "" {
		fmt.Printf("\t(%s)\n", CustomEnglishTimeAgo.Format(parsed))
	}
}
