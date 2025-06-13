package pakang

import (
	"os"
	"strings"
)

type OsRelease struct {
	data map[string]string
}

func (self *OsRelease) Set(key string, value string) {
	self.data[key] = value
}

func (self OsRelease) Param(name string) string {
	return self.data[name]
}

func (self OsRelease) ParamContains(param, subvalue string) bool {
	for key, val := range self.data {
		if param != key {
			continue
		}

		return strings.Index(val, subvalue) >= 0
	}
	return false
}

func (self OsRelease) ParamGteValueInts(param string, expect string) bool {
	if self.data[param] == "" {
		return false
	}
	sys_data, reference := extractInts(self.data[param], expect)

	return ArrIntsGte(reference, sys_data)
}

func (self OsRelease) ParamLteValueInts(param string, expect string) bool {
	if self.data[param] == "" {
		return false
	}
	sys_data, reference := extractInts(self.data[param], expect)

	return ArrIntsLte(reference, sys_data)
}

func (self OsRelease) ParamGtValueInts(param string, expect string) bool {
	if self.data[param] == "" {
		return false
	}
	sys_data, reference := extractInts(self.data[param], expect)

	return ArrIntsGt(reference, sys_data)
}

func (self OsRelease) ParamLtValueInts(param string, expect string) bool {
	if self.data[param] == "" {
		return false
	}
	sys_data, reference := extractInts(self.data[param], expect)

	return ArrIntsLt(reference, sys_data)
}

func extractInts(data, expect string) ([]int, []int) {
	sys_data, err := ExtractInts(data)
	FailIf(err, 1, "Failed number parsing on OSR")

	reference, err := ExtractInts(expect)
	FailIf(err, 1, "Failed number parsing on query")

	return sys_data, reference
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LoadOsRelease() OsRelease {
	data, err := os.ReadFile("/etc/os-release")
	FailIf(err, 1, "Could not read /etc/os-release file")

	m := make(map[string]string)
	os_release := OsRelease{m}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		i := strings.Index(line, "=")
		if i < 0 {
			continue
		}

		(&os_release).Set(line[:i], strings.Trim(line[i+1:], "\""))
	}

	return os_release
}
