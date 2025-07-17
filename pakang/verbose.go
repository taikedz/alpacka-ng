package pakang

import (
	"fmt"
	"os"
)

var VERBOSE bool = false

func printVerbose(template string, items ...any) {
	if VERBOSE {
		setColor("30")
		fmt.Printf("$> ")
		fmt.Printf(template, items...)
		setColor("0")
	}
}

func setColor(val string) {
	// https://no-color.org/
	/*
		Command-line software which adds ANSI color to its output by default should check
		for a NO_COLOR environment variable that, when present and not an empty string
		(regardless of its value), prevents the addition of ANSI color.
	*/
	if val, isset := os.LookupEnv("NO_COLOR"); isset && len(val) > 0 {
		return
	}
	fmt.Printf("\033[%sm", val)
}
