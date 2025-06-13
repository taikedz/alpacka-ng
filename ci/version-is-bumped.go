package main

import (
    "os"
    "fmt"
    "strings"
    re "regexp"
    "os/exec"
    
    "github.com/taikedz/alpacka-ng/pakang"
)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "show" {
        fmt.Printf("%s\n", pakang.GetVersionString())
        return
    }

    latest_tag_nums := getLatestGitTag()
    version_nums := pakang.GetVersionInts()

    if !pakang.ArrIntsGt(latest_tag_nums, version_nums) {
        fmt.Printf("Current version %v not greater than latest tag %v\n", version_nums, latest_tag_nums)
        println("Need to bump version.")
        os.Exit(1)
    }
}

func getLatestGitTag() []int {
    tokens := []string{"log", "--oneline", "--decorate=short", "--color=never"}
    cmd := exec.Command("git", tokens...)

    output, err := cmd.Output()
    pakang.FailIf(err, 1, "Could not run 'git log'")

    verstr, err := find("tag: v(\\d\\.\\d\\.\\d)", strings.Split(string(output), "\n"))
    pakang.FailIf(err, 1, "Could not find tag")

    res, err := pakang.ExtractInts(verstr)
    pakang.FailIf(err, 1, "Could not parse supplied version string")

    return res
}

func find(expr string, lines []string) (string, error) {
    pat, err := re.Compile(expr)
    pakang.FailIf(err, 1, "Could not compile expression", expr)
    for _,line := range lines {
        res := pat.FindStringSubmatch(line)
        if len(res) > 0 {
            return res[1], nil
        }
    }
    return "", fmt.Errorf("Could not find pattern")
}

