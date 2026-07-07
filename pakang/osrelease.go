package pakang

import (
	"os"
	"strings"
)

type OsRelease struct {
	data map[string]string
}

func (osr *OsRelease) Set(key string, value string) {
	osr.data[key] = value
}

func (osr OsRelease) Param(name string) string {
	return osr.data[name]
}

func (osr OsRelease) ParamContains(param, subvalue string) bool {
	for key, val := range osr.data {
		if param != key {
			continue
		}

		return strings.Contains(val, subvalue)
	}
	return false
}

func (osr OsRelease) ParamGteValueInts(param string, expect string) bool {
	if osr.data[param] == "" {
		return false
	}
	sys_data, reference := extractInts(osr.data[param], expect)

	return ArrIntsGte(reference, sys_data)
}

func (osr OsRelease) ParamLteValueInts(param string, expect string) bool {
	if osr.data[param] == "" {
		return false
	}
	sys_data, reference := extractInts(osr.data[param], expect)

	return ArrIntsLte(reference, sys_data)
}

func (osr OsRelease) ParamGtValueInts(param string, expect string) bool {
	if osr.data[param] == "" {
		return false
	}
	sys_data, reference := extractInts(osr.data[param], expect)

	return ArrIntsGt(reference, sys_data)
}

func (osr OsRelease) ParamLtValueInts(param string, expect string) bool {
	if osr.data[param] == "" {
		return false
	}
	sys_data, reference := extractInts(osr.data[param], expect)

	return ArrIntsLt(reference, sys_data)
}

func extractInts(data, expect string) ([]int, []int) {
	sys_data, err := ExtractInts(data)
	FailIf(err, 1, "Failed number parsing on OSR")

	reference, err := ExtractInts(expect)
	FailIf(err, 1, "Failed number parsing on query")

	return sys_data, reference
}

func LoadOsRelease() OsRelease {
	data, err := os.ReadFile("/etc/os-release")
	FailIf(err, 1, "Could not read /etc/os-release file")

	m := make(map[string]string)
	os_release := OsRelease{m}
	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		before, after, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		(&os_release).Set(before, strings.Trim(after, "\""))
	}

	return os_release
}
