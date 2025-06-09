package pakang


type DnfPM struct {
    pm_cmd string
    extraflags []string
}

func NewDnfPM(pm_cmd string, flags []string) DnfPM {
    return DnfPM{pm_cmd, flags}
}

func (self DnfPM) Name() string { return "DNF (or yum)" }

func (self DnfPM) Help() []string {
    return []string{
        "clean : Clean the cache",
    }
}

func (self DnfPM) Extra(terms []string) {
    if ArrayHas("clean", self.extraflags) {
        self.clean()
    } else {
        self.Search(terms)
    }
}

func (self DnfPM) clean() {
    RunCmd(NEED_ROOT, self.pm_cmd, "clean").OrFail("Clean failed")
    RunCmd(NEED_ROOT, self.pm_cmd, "autoremove").OrFail("Auto remove failed")
}

func (self DnfPM) Search(terms []string) {
    cmd := []string{self.pm_cmd, "search"}
    cmd = append(cmd, terms...)
    RunCmd(0, cmd...)
}

func (self DnfPM) Show(pkg string) {
    RunCmd(0, self.pm_cmd, "info", pkg)
}

func (self DnfPM) Update() {
    // Do nothing. yum/dnf does this on its own. With no control, sadly.
}

func (self DnfPM) Install(yes bool, packages []string) {
    cmd := []string{self.pm_cmd, "install"}
    if yes {
        cmd = append(cmd, "-y")
    }
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...).OrFail("Install failed")
}

func (self DnfPM) Remove(packages []string) {
    cmd := []string{self.pm_cmd, "remove"}
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...).OrFail("Remove failed")
}

func (self DnfPM) Upgrade(yes bool) {
    cmd := []string{self.pm_cmd, "upgrade"}
    if yes {
        cmd = append(cmd, "-y")
    }
    RunCmd(NEED_ROOT, cmd...).OrFail("Upgrade failed")
}
