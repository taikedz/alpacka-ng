package pakang

import (
	"fmt"
	"os"
	"os/user"
	"regexp"
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
	return "", fmt.Errorf("required parameter '%s' not found", key)
}

func IsRootUser() bool {
	if os.Getenv("PAF_TEST_PMAN") != "" {
		// test mode - we'll always say we're not root, to catch sudo detection
		return false
	}
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
		return false, fmt.Errorf("not on Windows")
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
func ArrIntsGt_b(reference, held_data []int, equalok bool) bool {

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
	// they are equal at this point
	return equalok
}

func ArrIntsGte(reference, held_data []int) bool {
	return ArrIntsGt_b(reference, held_data, true)
}

func ArrIntsGt(reference, held_data []int) bool {
	return ArrIntsGt_b(reference, held_data, false)
}

// Check that held_data <= reference
func ArrIntsLt_b(reference, held_data []int, equalok bool) bool {

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
	// they are equal by now
	return equalok
}

func ArrIntsLte(reference, held_data []int) bool {
	return ArrIntsLt_b(reference, held_data, true)
}

func ArrIntsLt(reference, held_data []int) bool {
	return ArrIntsLt_b(reference, held_data, false)
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

func extractSection(section, text string) string {
	lines := strings.Split(text, "\n")
	var desc_lines []string
	extracting := false
	pat, err := regexp.Compile(fmt.Sprintf("%s(-[^:]+?):", section))
	FailIf(err, 1, "Dev error: Invalid section pattern")

	for _, line := range lines {
		if pat.MatchString(line) {
			extracting = true
			desc_lines = append(desc_lines, line)
		} else if extracting && len(line) > 0 && line[0] == ' ' {
			desc_lines = append(desc_lines, line)
		} else {
			extracting = false
		}
	}
	return strings.Join(desc_lines, "\n")
}

func parapend(prefix string, data []string, suffix string) []string {
	var new_strings []string
	for _, item := range data {
		new_strings = append(new_strings, fmt.Sprintf("%s%s%s", prefix, item, suffix))
	}
	return new_strings
}
