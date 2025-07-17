package pakang

type PacmanPM struct {
	extraflags []string
}

func NewPacmanPM(flags []string) PacmanPM {
	return PacmanPM{flags}
}

func (pm PacmanPM) Name() string { return "pacman" }

func (pm PacmanPM) Help() []string {
	return []string{}
}

func (pm PacmanPM) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{"pacman", "-Ss"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (pm PacmanPM) NoAction(terms []string) {
	pm.Search(terms)
}

func (pm PacmanPM) Clean() {
	RunCmd(NEED_ROOT, "pacman", "-Scc").OrFail("Clean failed")
}

func (pm PacmanPM) Show(pkg string) {
	RunCmd(0, "pacman", "-Si", pkg).OrFail("Show failed")
}

func (pm PacmanPM) Update() {
	RunCmd(NEED_ROOT, "pacman", "-Sy").OrFail("Package index update failed")
}

func (pm PacmanPM) Install(yes bool, packages []string) {
	cmd := []string{"pacman", "-S"}
	if yes {
		cmd = append(cmd, "--noconfirm")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install failed")
}

func (pm PacmanPM) Remove(packages []string) {
	cmd := []string{"pacman", "-Rs"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Remove failed")
}

func (pm PacmanPM) Upgrade(yes bool) {
	cmd := []string{"pacman"}
	if yes {
		cmd = append(cmd, "--noconfirm")
	}
	cmd = append(cmd, "-Su")
	RunCmd(NEED_ROOT, cmd...).OrFail("Upgrade failed")
}
