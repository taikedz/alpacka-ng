package pakang

import (
	"fmt"
	"os"
)

func Fail(code int, message string, err error) {
    if err == nil {
        fmt.Println(message)
    } else {
        fmt.Printf("%s : %v\n", message, err)
    }
    os.Exit(code)
}
