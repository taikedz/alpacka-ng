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
	if val, isset := os.LookupEnv("NOCOLOR"); isset && val != "false" {
		return
	}
	fmt.Printf("\033[%sm", val)
}
