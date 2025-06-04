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
        "ppa=$PPA_ID : Add a PPA"
    }
}

func (self AptPM) Search(terms []string) {
    cmd := []string{"apt-cache", "search"}
    cmd = append(cmd, terms...)
    RunCmd(0, cmd...).OrFail("Search failed")
}

func (self AptPM) Extra(terms []string) {
    if ArrayHas("clean", self.extraflags) {
        self.clean()
    } else if ArrayHas("fix", self.extraflags) {
        self.fixbroken()
    } else if val, err := ExtractValueOfKey("ppa", self.extraflags); err != nil {
        self.addPpa(val)
    } else {
        self.Search(terms)
    }
}

func (self AptPM) clean() {
    RunCmd(NEED_ROOT, "apt-get", "autoclean").OrFail("Auto clean failed")
    RunCmd(NEED_ROOT, "apt-get", "autoremove").OrFail("Atu remove failed")
}

func (self AptPM) fixbroken() {
    RunCmd(NEED_ROOT, "apt-get", "-f", "install").OrFail("Install fix failed")
}

func (self AptPM) addPp(ppa_id string) {
    RunCmdOut(false, 0, "which", "add-apt-repository").OrFail("'add-apt-repository' command required, but not found on this system.")
    RunCmd(NEED_ROOT, "add-apt-repository", ppa_id).OrFail("Could not add PPA respoitory")
}

func (self AptPM) Show(pkg string) {
    RunCmd(0, "apt-cache", "show", pkg).OrFail("Error")
}

func (self AptPM) Update() {
    RunCmd(NEED_ROOT, "apt-get", "update").OrFail("Could not update package index")
}

func (self AptPM) Install(yes bool, packages []string) {
    cmd := []string{"apt-get", "install"}
    if yes {
        cmd = append(cmd, "-y")
    }
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...).OrFail("Install operation failed")
}

func (self AptPM) Remove(packages []string) {
    cmd := []string{"apt-get", "remove"}
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...).OrFail("Package removal failed")
}

func (self AptPM) Upgrade(yes bool) {
    cmd := []string{"apt-get", "upgrade"}
    if yes {
        cmd = append(cmd, "-y")
    }
    RunCmd(NEED_ROOT, cmd...).OrFail("Upgarde failed")
}
