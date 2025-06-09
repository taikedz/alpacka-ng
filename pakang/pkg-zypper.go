package pakang

type ZypperPM struct {
    extraflags []string
}

func NewZypperPM(flags []string) ZypperPM {
    return ZypperPM{flags}
}

func (self ZypperPM) Name() string { return "zypper" }

func (self ZypperPM) Help() []string {
    return []string{
        "clean : Clean the cache",
        "fix : fix broken dependencies",
    }
}

func (self ZypperPM) Search(terms []string) {
    cmd := []string{"zypper", "search"}
    cmd = append(cmd, terms...)
    RunCmd(0, cmd...).OrFail("Search failed")
}

func (self ZypperPM) Extra(terms []string) {
    if ArrayHas("clean", self.extraflags) {
        self.clean()
    } else {
        self.Search(terms)
    }
}

func (self ZypperPM) clean() {
    RunCmd(NEED_ROOT, "zypper", "clean").OrFail("Clean failed")
}

func (self ZypperPM) Show(pkg string) {
    RunCmd(0, "zypper", "info", pkg).OrFail("Show package failed")
}

func (self ZypperPM) Update() {
    RunCmd(NEED_ROOT, "zypper", "refresh").OrFail("Package index update failed")
}

func (self ZypperPM) Install(yes bool, packages []string) {
    cmd := []string{"zypper", "install"}
    if yes {
        cmd = append(cmd, "-y")
    }
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...).OrFail("Install failed")
}

func (self ZypperPM) Remove(packages []string) {
    cmd := []string{"zypper", "remove"}
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...).OrFail("Remove failed")
}

func (self ZypperPM) Upgrade(yes bool) {
    cmd := []string{"zypper", "update"}
    if yes {
        cmd = append(cmd, "-y")
    }
    RunCmd(NEED_ROOT, cmd...).OrFail("Index update failed")
}
