package pakang

type FlatpakPM struct {
	extraflags []string
}

func NewFlatpakPM(flags []string) FlatpakPM {
	return FlatpakPM{flags}
}

func (pm FlatpakPM) Name() string { return "Flatpak package manager" }

func (pm FlatpakPM) Help() []string {
	return []string{}
}

func (pm FlatpakPM) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{"flatpak", "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (pm FlatpakPM) NoAction(terms []string) {
	pm.Search(terms)
}

func (pm FlatpakPM) Clean() {
}

func (pm FlatpakPM) Show(pkg string) {
	RunCmd(0, "flatpak", "remote-info", pkg).OrFail("Error")
}

func (pm FlatpakPM) Update() {
}

func (pm FlatpakPM) Install(yes bool, packages []string) {
	cmd := []string{"flatpak", "install"}
	if yes {
		cmd = append(cmd, "-y")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install operation failed")
}

func (pm FlatpakPM) Remove(packages []string) {
	cmd := []string{"flatpak", "uninstall"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Package removal failed")
}

func (pm FlatpakPM) Upgrade(yes bool) {
	cmd := []string{"flatpak", "update"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Refresh (upgrade) failed")
}
