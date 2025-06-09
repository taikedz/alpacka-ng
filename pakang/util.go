package pakang

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

func ArrayHas(term string, stuff []string) bool {
	for _, thing := range stuff {
		if term == thing {
			return true
		}
	}
	return false
}

func ExcludeStr(input []string, exclude []string) []string {
	var retained []string

	for _, s := range input {
		if !ArrayHas(s, exclude) {
			retained = append(retained, s)
		}
	}

	return retained
}

func ExtractValueOfKey(key string, items []string) (string, error) {
	// assume an array of "key=value" strings
	// locate key , split on '=', return the value
	key_eq := fmt.Sprintf("%s=", key)
	for _, item := range items {
		if strings.Index(item, key_eq) == 0 {
			return item[len(key_eq):], nil
		}
	}
	return "", fmt.Errorf("Requred parameter '%s' not found", key)
}

func IsRootUser() bool {
	u, e := user.Current()
	FailIf(e, 98, "Fatal - Could not get current user!")
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

func SplitStringMultichar(data string, chars string) []string {
	tokens := []string{data}

	for _, c := range chars {
		tokens = SplitStringsChar(tokens, string(c))
	}
	return tokens
}

func SplitStringsChar(data []string, char string) []string {
	var tokens []string
	for _, piece := range data {
		tokens = append(tokens, strings.Split(piece, string(char))...)
	}
	return tokens
}

// Check that held_data >= reference
func ArrIntsGte(reference, held_data []int) bool {

	z := max(len(reference), len(held_data))
	for i := 0; i < z; i++ {
		held_data_v := 0
		if i < len(held_data) {
			held_data_v = held_data[i]
		}
		reference_v := 0
		if i < len(reference) {
			reference_v = reference[i]
		}
		if held_data_v > reference_v {
			return true
		}
		if held_data_v < reference_v {
			return false
		}
	}
	return true // they are equal by now
}

// Check that held_data <= reference
func ArrIntsLte(reference, held_data []int) bool {

	z := max(len(reference), len(held_data))
	for i := 0; i < z; i++ {
		held_data_v := 0
		if i < len(held_data) {
			held_data_v = held_data[i]
		}
		reference_v := 0
		if i < len(reference) {
			reference_v = reference[i]
		}
		if held_data_v < reference_v {
			return true
		}
		if held_data_v > reference_v {
			return false
		}
	}
	return true // they are equal by now
}

func ExtractInts(data string) ([]int, error) {
	tokens := SplitStringMultichar(data, ".,: ")
	var nums []int
	for _, t := range tokens {
		n, e := strconv.Atoi(strings.Trim(t, " "))
		if e != nil {
			return nil, e
		}
		nums = append(nums, n)
	}

	return nums, nil
}
