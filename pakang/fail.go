package pakang

import (
	"fmt"
	"os"
)

/* Immediately exit the program with the specified error code.
*/
func Fail(code int, message string, err error) {
	if err == nil {
		fmt.Println(message)
	} else {
		fmt.Printf("%s : %v\n", message, err)
	}
	os.Exit(code)
}

// Conditionally
func FailIf(err error, code int, message string, items ...any) {
	if err != nil {
		Fail(code, fmt.Sprintf(message, items...), err)
	}
}
