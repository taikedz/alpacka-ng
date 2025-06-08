package pakang

import (
	"os"
	"strings"
	"strconv"
)

type OsRelease struct {
	data map[string]string
}

func (self OsRelease) ParamContains(param, subvalue string) bool {
	for key, val := range self.data {
		if param != key { continue }

		return strings.Index(subvalue, val) >= 0
	}
	return false
}

func (self OsRelease) ParamGteValueInts(param string, expect string) bool {
	if self.data[param] == "" { return false }

	tokens, err := self.extractInts(self.data[param])
	FailIf(err, 1, "Failed number parsing")

	target, err := self.extractInts(expect)
	FailIf(err, 1, "Failed number parsing")

	z := max(len(target), len(tokens))
	for i:=0 ; i++ ; i<z {
		token_v := 0
		if len(tokens) < i {
			token_v = tokens[i]
		}
		target_v := 0
		if len(target) < i {
			target_v = target[i]
		}
		if token_v > target_v { return true }
		if token_v < target_v { return false }
	}
	return true // they are equal by now
}

func (self OsRelease) extractInts(data string) ([]int, error) {
	tokens := strings.Split(data, ",")
	nums []int
	for _,t := range tokens {
		n, e := strconv.Atoi(strings.Trim(t, " "))
		if e != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func LoadOsRelease() OsRelease {
	data, err := os.ReadFile("/etc/os-release")
	FailIf(err, 1, "Could not read /etc/os-release file")

	m := make(map[string]string)
	os_release := OsRelease{m}
	lines := strings.Split(data, "\n")
	for _,line := range lines {
		i := strings.Index("=", line)
		if i < 0 { continue }

		os_release[ line[:i] ] = strings.Trim(line[i+1,], "\"")
	}

	return os_release
}