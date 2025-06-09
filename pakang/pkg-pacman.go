package pakang


type PacmanPM struct {
    extraflags []string
    updated *bool
}

func NewPacmanPM(flags []string) PacmanPM {
    updated := false
    return PacmanPM{flags, &updated}
}

func (self PacmanPM) Name() string { return "pacman" }

func (self PacmanPM) Help() []string {
    return []string{
        // "clean : Clean the cache",
        // "fix : fix broken dependencies",
    }
}

func (self PacmanPM) Search(terms []string) {
    cmd := []string{"pacman", "-Ss"}
    cmd = append(cmd, terms...)
    RunCmd(0, cmd...).OrFail("Search failed")
}

func (self PacmanPM) Extra(terms []string) {
    if ArrayHas("clean", self.extraflags) {
        self.clean()
    } else {
        self.Search(terms)
    }
}

func (self PacmanPM) clean() {
    RunCmd(NEED_ROOT, "pacman", "-Scc").OrFail("Clean failed")
}

func (self PacmanPM) Show(pkg string) {
    RunCmd(0, "pacman", "-Si", pkg).OrFail("Show failed")
}

func (self PacmanPM) Update() {
    RunCmd(NEED_ROOT, "pacman", "-Syy").OrFail("PIndex update failed")
    *self.updated = true
}

func (self PacmanPM) Install(yes bool, packages []string) {
    cmd := []string{"pacman", "-S"}
    if yes {
        cmd = append(cmd, "-y")
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
        cmd = append(cmd, "-y")
    }
    if *self.updated {
        cmd = append(cmd, "-Su")
    } else {
        cmd = append(cmd, "-Syu")
    }
    RunCmd(NEED_ROOT, cmd...).OrFail("Upgrade failed")
}
