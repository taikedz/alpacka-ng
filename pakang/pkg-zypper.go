package pakang

type ZypperPM struct {
	extraflags []string
}

func NewZypperPM(flags []string) ZypperPM {
	return ZypperPM{flags}
}

func (pm ZypperPM) Name() string { return "zypper" }

func (pm ZypperPM) Help() []string {
	return []string{}
}

func (pm ZypperPM) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{"zypper", "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (pm ZypperPM) NoAction(terms []string) {
	pm.Search(terms)
}

func (pm ZypperPM) Clean() {
	RunCmd(NEED_ROOT, "zypper", "clean").OrFail("Clean failed")
}

func (pm ZypperPM) Show(pkg string) {
	RunCmd(0, "zypper", "info", pkg).OrFail("Show package failed")
}

func (pm ZypperPM) Update() {
	RunCmd(NEED_ROOT, "zypper", "refresh").OrFail("Package index update failed")
}

func (pm ZypperPM) Install(yes bool, packages []string) {
	cmd := []string{"zypper", "install"}
	if yes {
		cmd = append(cmd, "-y")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install failed")
}

func (pm ZypperPM) Remove(packages []string) {
	cmd := []string{"zypper", "remove"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Remove failed")
}

func (pm ZypperPM) Upgrade(yes bool) {
	cmd := []string{"zypper", "update"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Upgrade failed")
}
