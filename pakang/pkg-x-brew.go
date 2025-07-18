package pakang

type Brew struct {
	extraflags []string
}

func NewBrewPM(flags []string) Brew {
	return Brew{flags}
}

func (pm Brew) Name() string { return "Homebrew package manager" }

func (pm Brew) Help() []string {
	return []string{}
}

func (pm Brew) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{"brew", "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (pm Brew) NoAction(terms []string) {
	pm.Search(terms)
}

func (pm Brew) Clean() {
	RunCmd(0, "brew", "cleanup")
}

func (pm Brew) Show(pkg string) {
	RunCmd(0, "brew", "info", pkg).OrFail("Error")
}

func (pm Brew) Update() {
}

func (pm Brew) Install(yes bool, packages []string) {
	cmd := []string{"brew", "install"}
	if yes {
		cmd = append(cmd, "-y")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install operation failed")
}

func (pm Brew) Remove(packages []string) {
	cmd := []string{"brew", "uninstall"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Package removal failed")
}

func (pm Brew) Upgrade(yes bool) {
	cmd := []string{"brew", "upgrade"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Refresh (upgrade) failed")
}
