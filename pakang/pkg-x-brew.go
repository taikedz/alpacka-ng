package pakang

type BrewPM struct {
	extraflags []string
}

func NewBrewPM(flags []string) BrewPM {
	return BrewPM{flags}
}

func (pm BrewPM) Name() string { return "Homebrew package manager" }

func (pm BrewPM) Help() []string {
	return []string{}
}

func (pm BrewPM) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{"brew", "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (pm BrewPM) NoAction(terms []string) {
	pm.Search(terms)
}

func (pm BrewPM) Clean() {
	RunCmd(0, "brew", "cleanup")
}

func (pm BrewPM) Show(pkg string) {
	RunCmd(0, "brew", "info", pkg).OrFail("Error")
}

func (pm BrewPM) Update() {
}

func (pm BrewPM) Install(yes bool, packages []string) {
	cmd := []string{"brew", "install"}
	if yes {
		cmd = append(cmd, "-y")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install operation failed")
}

func (pm BrewPM) Remove(packages []string) {
	cmd := []string{"brew", "uninstall"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Package removal failed")
}

func (pm BrewPM) Upgrade(yes bool) {
	cmd := []string{"brew", "upgrade"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Refresh (upgrade) failed")
}
