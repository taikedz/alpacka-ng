package main

import (
    "os"
    "fmt"
    "github.com/taikedz/alpacka-ng/pakang"
)

func main() {
    manifest := pakang.LoadManifest(os.Args[1])
    fmt.Printf("%#v\n", manifest)
}
