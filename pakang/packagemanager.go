package pakang

import (
	"fmt"
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

func GetPackageManager(specific_pm string, extra []string) PackageManager {
	if specific_pm != "" {
		if !checkFor(specific_pm) {
			Fail(1, fmt.Sprintf("%s not available on this host", specific_pm), nil)
		}
		switch specific_pm {
		case "snap":
			return NewSnapPM(extra)
		case "flatpak":
			return NewFlatpakPM(extra)
		case "brew":
			return NewBrewPM(extra)
		}
	}

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
