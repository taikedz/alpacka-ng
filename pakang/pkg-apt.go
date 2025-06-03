package pakang

type AptPM struct {
    extraflags []string
}

func NewAptPM(flags []string) AptPM {
    return AptPM{flags}
}

func (self AptPM) Help() []string {
    return []string{
        "clean : Clean the cache",
        "fix : fix broken dependencies",
    }
}

func (self AptPM) Search(terms []string) {
    cmd := []string{"apt-cache", "search"}
    cmd = append(cmd, terms...)
    RunCmd(0, cmd...)
}

func (self AptPM) Extra(terms []string) {
    if ArrayHas("clean", self.extraflags) {
        self.clean()
    } else if ArrayHas("fix", self.extraflags) {
        self.fixbroken()
    } else {
        self.Search(terms)
    }
}

func (self AptPM) clean() {
    RunCmd(NEED_ROOT, "apt-get", "autoclean")
    RunCmd(NEED_ROOT, "apt-get", "autoremove")
}

func (self AptPM) fixbroken() {
    RunCmd(NEED_ROOT, "apt-get", "-f", "install")
}

func (self AptPM) Show(pkg string) {
    RunCmd(0, "apt-cache", "show", pkg)
}

func (self AptPM) Update() {
    RunCmd(NEED_ROOT, "apt-get", "update")
}

func (self AptPM) Install(yes bool, packages []string) {
    cmd := []string{"apt-get", "install"}
    if yes {
        cmd = append(cmd, "-y")
    }
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...)
}

func (self AptPM) Remove(packages []string) {
    cmd := []string{"apt-get", "remove"}
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...)
}

func (self AptPM) Upgrade(yes bool) {
    cmd := []string{"apt-get", "upgrade"}
    if yes {
        cmd = append(cmd, "-y")
    }
    RunCmd(NEED_ROOT, cmd...)
}
