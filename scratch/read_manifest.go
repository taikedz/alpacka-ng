package main

import (
    "os"
    "fmt"
    "github.com/taikedz/alpacka-ng/pakang"
)

func main() {
    manifest := pakang.LoadManifest(os.Args[1])
    fmt.Printf("%#v\n", manifest)

    osr := pakang.LoadOsRelease()
    fmt.Printf("%v\n", osr.ParamGteValueInts("VERSION_ID", "22.04"))
    fmt.Printf("%v\n", osr.ParamGteValueInts("VERSION_ID", "24.04"))
    fmt.Printf("%v\n", osr.ParamGteValueInts("VERSION_ID", "24.09"))

    if osr.Param("ID") == "ubuntu" {
        println("is actually ubuntu")
    }
    if osr.ParamContains("ID_LIKE", "ubuntu") {
        println("is like ubuntu")
    }
    if osr.ParamContains("ID_LIKE", "debian") {
        println("is like debian")
    }
    if osr.ParamGteValueInts("VERSION_ID", "26.04") {
        println("version is a lie")
    }
    if osr.ParamGteValueInts("VERSION_ID", "24.04") {
        println("modern LTS")
    }
    if osr.ParamLteValueInts("VERSION_ID", "18.04") {
        println("outdated LTS")
    }

}
