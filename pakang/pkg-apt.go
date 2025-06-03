package pakang

/* Define RunCmd(runmodes int, command string, fmttokens .... string)
 *
 * Combine the command and fmttokens with Sprintf
 * Define runmodes for NEED_ROOT and USE_LESS , determine these dynamically somehow...
 */

type AptPM struct {
    extraflags []string
}

func NewAptPM(flags []string) AptPM {
    return AptPM{flags}
}

func (self AptPM) Show(pkg string) {
    RunCmd("apt-cache", "show", pkg)
}

func (self AptPM) Update() {
    RunCmd("sudo", "apt-get", "update")
}

func (self AptPM) Install(yes bool, packages []string) {
    cmd := []string{"sudo", "apt-get"}
    if yes {
        cmd = append(cmd, "-y")
    }
    cmd = append(cmd, "install")
    cmd = append(cmd, packages...)
    RunCmd(cmd...)
}

func (self AptPM) Remove(packages []string) {
    cmd := []string{"sudo", "apt-get", "remove"}
    cmd = append(cmd, packages...)
    RunCmd(cmd...)
}

func (self AptPM) Upgrade(yes bool) {
    cmd := []string{"sudo", "apt-get"}
    if yes {
        cmd = append(cmd, "-y")
    }
    cmd = append(cmd, "upgrade")
    RunCmd(cmd...)
}
