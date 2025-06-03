package pakang


type AptPM struct {
    extraflags []string
}

func NewAptPM(flags []string) AptPM {
    return AptPM{flags}
}

func (self AptPM) Show(pkg string) {
    RunCmd(0, "apt-cache", "show", pkg)
}

func (self AptPM) Update() {
    RunCmd(NEED_ROOT, "sudo", "apt-get", "update")
}

func (self AptPM) Install(yes bool, packages []string) {
    cmd := []string{"apt-get"}
    if yes {
        cmd = append(cmd, "-y")
    }
    cmd = append(cmd, "install")
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...)
}

func (self AptPM) Remove(packages []string) {
    cmd := []string{"apt-get", "remove"}
    cmd = append(cmd, packages...)
    RunCmd(NEED_ROOT, cmd...)
}

func (self AptPM) Upgrade(yes bool) {
    cmd := []string{"apt-get"}
    if yes {
        cmd = append(cmd, "-y")
    }
    cmd = append(cmd, "upgrade")
    RunCmd(NEED_ROOT, cmd...)
}
