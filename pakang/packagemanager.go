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
	Name() string
}

var found_pm PackageManager = nil

func checkFor(cmd string) bool {
	return RunCmdOut(false, 0, "which", cmd).Ok()
}

func GetPackageManager(extra []string) PackageManager {
	RunCmdOut(false, 0, "which", "which").OrFail("Could not run 'which'")

	if checkFor("apt-get") {
		return NewAptPM(extra)

	} else if checkFor("dnf") {
		return NewDnfPM("dnf", extra)

	} else if checkFor("yum") {
		return NewDnfPM("yum", extra)

	} else if checkFor("apk") {
		return NewApkPM(extra)

	} else if checkFor("pacman") {
		return NewPacmanPM(extra)

	} else if checkFor("zypper") {
		return NewZypperPM(extra)
	}

	Fail(10, "No package manager found.", nil)
	return nil // for the compiler ...
}
