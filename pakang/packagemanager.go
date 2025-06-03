package pakang

/*
import (
    "fmt"
)
//*/

type PackageManager interface {
    Update()
    Extra(terms []string)
    Search(terms []string)
    Show(pkg string)
    Install(yes bool, packages []string)
    Upgrade(yes bool)
    Remove(packages []string)
    Help() []string
}

var found_pm PackageManager = nil

func GetPackageManager(extra []string) PackageManager {
    if status, err := RunCmdOut(false, 0, "which", "which"); status < 0 {
        Fail(1, "Could not run 'which' : %v", err)
    }

    if status, _ := RunCmdOut(false, 0, "which", "apt-get"); status == 0 {
        return NewAptPM(extra)
    } else if status, _ := RunCmdOut(false, 0, "which", "dnf"); status == 0 {
        return NewDnfPM("dnf", extra)
    } else if status, _ := RunCmdOut(false, 0, "which", "yum"); status == 0 {
        return NewDnfPM("yum", extra)
    }

    Fail(10, "No package manager found.", nil)
    return nil // for the compiler ...
}
