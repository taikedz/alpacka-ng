package pakang

import (
	"fmt"
	"strings"
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