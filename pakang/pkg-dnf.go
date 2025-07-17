package pakang

type DnfPM struct {
	pm_cmd     string
	extraflags []string
}

func NewDnfPM(pm_cmd string, flags []string) DnfPM {
	return DnfPM{pm_cmd, flags}
}

func (pm DnfPM) Name() string { return "DNF (or yum)" }

func (pm DnfPM) Help() []string {
	return []string{}
}

func (pm DnfPM) NoAction(terms []string) {
	pm.Search(terms)
}

func (pm DnfPM) Clean() {
	RunCmd(NEED_ROOT, pm.pm_cmd, "clean").OrFail("Clean failed")
	RunCmd(NEED_ROOT, pm.pm_cmd, "autoremove").OrFail("Auto remove failed")
}

func (pm DnfPM) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{pm.pm_cmd, "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...)
}

func (pm DnfPM) Show(pkg string) {
	RunCmd(0, pm.pm_cmd, "info", pkg)
}

func (pm DnfPM) Update() {
	// Do nothing. yum/dnf does this on its own. With no control, sadly.
}

func (pm DnfPM) Install(yes bool, packages []string) {
	cmd := []string{pm.pm_cmd, "install"}
	if yes {
		cmd = append(cmd, "-y")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install failed")
}

func (pm DnfPM) Remove(packages []string) {
	cmd := []string{pm.pm_cmd, "remove"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Remove failed")
}

func (pm DnfPM) Upgrade(yes bool) {
	cmd := []string{pm.pm_cmd, "upgrade"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Upgrade failed")
}
