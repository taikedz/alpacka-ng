package pakang

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

var VERBOSE bool = false
var COLOR bool = true

func printVerbose(template string, items ...any) {
	if VERBOSE {
		setAnsi("30")
		fmt.Printf("$> ")
		fmt.Printf(template, items...)
		setAnsi("0")
	}
}

func loadColorizationMode() {
	val, isset := os.LookupEnv("NO_COLOR")
	COLOR = !(isset && len(val) > 0)

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		COLOR = false
	}
}

func setAnsi(ansi_code string) {
	// https://no-color.org/
	/*
		Command-line software which adds ANSI color to its output by default should check
		for a NO_COLOR environment variable that, when present and not an empty string
		(regardless of its value), prevents the addition of ANSI color.
	*/
	if COLOR {
		fmt.Printf("\033[%sm", ansi_code)
	}
}
