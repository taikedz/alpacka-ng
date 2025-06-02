package pakang

import (
    "fmt"
)

type PackageManager interface {
    Update()
    Show(pkg string)
    Install(yes *bool, packages []string)
    Remove(packages []string)
    Upgrade(yes *bool)
    Manifest(path string)
}

type NoopManager struct {
}

func (self NoopManager) Update() {
    fmt.Printf("apt-get update\n")
}

func (self NoopManager) Install(yes *bool, packages []string) {
    fmt.Printf("apt-get install -y=%v %v\n", *yes, packages)
}

func (self NoopManager) Upgrade(yes *bool) {
    fmt.Printf("apt-get upgrade -y=%v\n", *yes)
}

func (self NoopManager) Remove(packages []string) {
    fmt.Printf("apt-get remove %v\n", packages)
}

func (self NoopManager) Manifest(manifestfile string) {
    fmt.Printf("Use manifest file %v\n", manifestfile)
}

func (self NoopManager) Show(pkg string) {
    fmt.Printf("Show %v\n", pkg)
}

func GetPackageManager(extra []string) PackageManager {
    return NoopManager{}
}
