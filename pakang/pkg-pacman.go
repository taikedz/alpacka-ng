package pakang

type PacmanPM struct {
	extraflags []string
}

func NewPacmanPM(flags []string) PacmanPM {
	return PacmanPM{flags}
}

func (self PacmanPM) Name() string { return "pacman" }

func (self PacmanPM) Help() []string {
	return nil
}

func (self PacmanPM) Search(terms []string) {
	if len(terms) == 0 { return }
	cmd := []string{"pacman", "-Ss"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (self PacmanPM) NoAction(terms []string) {
	self.Search(terms)
}

func (self PacmanPM) Clean() {
	RunCmd(NEED_ROOT, "pacman", "-Scc").OrFail("Clean failed")
}

func (self PacmanPM) Show(pkg string) {
	RunCmd(0, "pacman", "-Si", pkg).OrFail("Show failed")
}

func (self PacmanPM) Update() {
	RunCmd(NEED_ROOT, "pacman", "-Sy").OrFail("Package index update failed")
}

func (self PacmanPM) Install(yes bool, packages []string) {
	cmd := []string{"pacman", "-S"}
	if yes {
		cmd = append(cmd, "--noconfirm")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install failed")
}

func (self PacmanPM) Remove(packages []string) {
	cmd := []string{"pacman", "-Rs"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Remove failed")
}

func (self PacmanPM) Upgrade(yes bool) {
	cmd := []string{"pacman"}
	if yes {
		cmd = append(cmd, "--noconfirm")
	}
	cmd = append(cmd, "-Su")
	RunCmd(NEED_ROOT, cmd...).OrFail("Upgrade failed")
}
