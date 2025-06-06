package pakang

import (
	"fmt"
	"strings"
	"runtime"
	"os/user"
	"os"
)

func ArrayHas(term string, stuff []string) bool {
	for _, thing := range(stuff) {
		if term == thing {
			return true
		}
	}
	return false
}

func ExtractValueOfKey(key string, items []string) (string, error) {
	// assume an array of "key=value" strings
	// locate key , split on '=', return the value
	key_eq := fmt.Sprintf("%s=", key)
	for _, item := range(items) {
		if strings.Index(key_eq, item) == 0 {
			return item[len(key_eq):], nil
		}
	}
	return "", fmt.Errorf("Requred parameter '%s' not found", key)
}

func IsRootUser() bool {
    u, e := user.Current()
    if e != nil {
        Fail(98, "Fatal - Could not get current user!", e)
    }
    return u.Uid == "0" // posix only!
}

func IsWinAdmin() (bool, error) {
	/* This is apparently the way to handle Windows.
	For now, not supporting windows choco/winget
	But this could be a target for future
	*/

	// https://stackoverflow.com/a/19847868/2703818
	if runtime.GOOS != "windows" {
		return false, fmt.Errorf("Not on Windows")
	}

	// https://stackoverflow.com/a/59147866/2703818
    _, err := os.Open("\\\\.\\PHYSICALDRIVE0")
    if err != nil {
        return false, nil
    }
    return true, nil
}