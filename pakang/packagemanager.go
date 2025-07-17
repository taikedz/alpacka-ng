package pakang

import (
	"os"
)

type PackageManager interface {
	Update()
	NoAction(terms []string)
	Search(terms []string)
	Show(pkg string)
	Install(yes bool, packages []string)
	Upgrade(yes bool)
	Remove(packages []string)
	Clean()
	Help() []string
	Name() string
}

func checkFor(cmd string) bool {
	// set PAF_TEST_PMAN to an explicit package manager (for testing)
	if pman, present := os.LookupEnv("PAF_TEST_PMAN"); present {
		return pman == cmd
	} else {
		return RunCmdOut(false, 0, "which", cmd).Ok()
	}
}

func checkPmanRequirements() {
	RunCmdOut(false, 0, "which", "which").OrFail("Could not run 'which'")
}

func GetPackageManager(extra []string) PackageManager {
	if checkFor("apt-get") {
		return NewAptPM(extra)

	} else if checkFor("dnf") {
		return NewDnfPM("dnf", extra)

	} else if checkFor("yum") {
		return NewDnfPM("yum", extra)

	} else if checkFor("pacman") {
		return NewPacmanPM(extra)

	} else if checkFor("zypper") {
		return NewZypperPM(extra)
	}

	Fail(10, "No package manager found.", nil)
	return nil // for the compiler ...
}
