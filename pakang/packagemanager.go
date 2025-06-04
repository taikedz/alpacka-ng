package pakang

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
    RunCmdOut(false, 0, "which", "which").OrFail("Could not run 'which'")

    if RunCmdOut(false, 0, "which", "apt-get").Ok() {
        return NewAptPM(extra)

    } else if RunCmdOut(false, 0, "which", "dnf").Ok() {
        return NewDnfPM("dnf", extra)

    } else if RunCmdOut(false, 0, "which", "yum").Ok() {
        return NewDnfPM("yum", extra)

    } else if RunCmdOut(false, 0, "which", "apk").Ok() {
        return NewApkPM(extra)

    } else if RunCmdOut(false, 0, "which", "pacman").Ok() {
        return NewPacmanPM(extra)

    } else if RunCmdOut(false, 0, "which", "zypper").Ok() {
        return NewZypperPM(extra)
    }

    Fail(10, "No package manager found.", nil)
    return nil // for the compiler ...
}
