package pakang

type SnapPM struct {
	extraflags []string
}

func NewSnapPM(flags []string) SnapPM {
	return SnapPM{flags}
}

func (pm SnapPM) Name() string { return "Snap package manager" }

func (pm SnapPM) Help() []string {
	return []string{
		"classic | c : use classic confinement (unconfined)",
	}
}

func (pm SnapPM) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{"snap", "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (pm SnapPM) NoAction(terms []string) {
	pm.Search(terms)
}

func (pm SnapPM) Clean() {
}

func (pm SnapPM) Show(pkg string) {
	RunCmd(0, "snap", "info", pkg).OrFail("Error")
}

func (pm SnapPM) Update() {
}

func (pm SnapPM) Install(yes bool, packages []string) {
	cmd := []string{"snap", "install"}
	if yes {
		cmd = append(cmd, "-y")
	}
	if ArrayHas("classic", pm.extraflags) || ArrayHas("c", pm.extraflags) {
		cmd = append(cmd, "--classic")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install operation failed")
}

func (pm SnapPM) Remove(packages []string) {
	cmd := []string{"snap", "remove"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Package removal failed")
}

func (pm SnapPM) Upgrade(yes bool) {
	cmd := []string{"snap", "refresh"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Refresh (upgrade) failed")
}
