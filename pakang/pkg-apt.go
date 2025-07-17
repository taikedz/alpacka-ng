package pakang

import (
	"fmt"
)

type AptPM struct {
	extraflags []string
}

func NewAptPM(flags []string) AptPM {
	return AptPM{flags}
}

func (pm AptPM) Name() string { return "APT package manager" }

func (pm AptPM) Help() []string {
	return []string{
		"fix : fix broken dependencies",
		"ppa=$PPA_ID : Add a PPA",
		"desc : Display descriptions of specified packages",
	}
}

func (pm AptPM) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{"apt-cache", "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (pm AptPM) NoAction(terms []string) {
	if ArrayHas("fix", pm.extraflags) {
		pm.fixbroken()
	} else if ArrayHas("desc", pm.extraflags) {
		for _, pkg := range terms {
			fmt.Printf("\n--- %s ---\n", pkg)
			pm.Show(pkg)
		}
	} else if val, err := ExtractValueOfKey("ppa", pm.extraflags); err == nil {
		pm.addPpa(val)
	} else {
		pm.Search(terms)
	}
}

func (pm AptPM) Clean() {
	RunCmd(NEED_ROOT, "apt-get", "autoclean").OrFail("Auto clean failed")
	RunCmd(NEED_ROOT, "apt-get", "autoremove").OrFail("Atu remove failed")
}

func (pm AptPM) fixbroken() {
	RunCmd(NEED_ROOT, "apt-get", "-f", "install").OrFail("Install fix failed")
}

func (pm AptPM) addPpa(ppa_id string) {
	RunCmdOut(false, 0, "which", "add-apt-repository").OrFail("'add-apt-repository' command required, but not found on this system.")
	RunCmd(NEED_ROOT, "add-apt-repository", ppa_id).OrFail("Could not add PPA respoitory")
}

func (pm AptPM) Show(pkg string) {
	if ArrayHas("desc", pm.extraflags) {
		res := RunCmdOut(false, 0, "apt-cache", "show", pkg)
		FailIf(res.GetError(), 1, "Could not get info for '%s'", pkg)

		desc := extractSection("Description", res.Stdout)
		fmt.Printf("%s\n", desc)
	} else {
		RunCmd(0, "apt-cache", "show", pkg).OrFail("Error")
	}
}

func (pm AptPM) Update() {
	RunCmd(NEED_ROOT, "apt-get", "update").OrFail("Could not update package index")
}

func (pm AptPM) Install(yes bool, packages []string) {
	cmd := []string{"apt-get", "install"}
	if yes {
		cmd = append(cmd, "-y")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install operation failed")
}

func (pm AptPM) Remove(packages []string) {
	cmd := []string{"apt-get", "remove"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Package removal failed")
}

func (pm AptPM) Upgrade(yes bool) {
	cmd := []string{"apt-get", "upgrade"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Upgarde failed")
}
