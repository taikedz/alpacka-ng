package main

import (
    "os"
    "github.com/taikedz/alpacka-ng/pakang"
)

func main() {
    pakang.LoadManifest(os.Args[1])
}
