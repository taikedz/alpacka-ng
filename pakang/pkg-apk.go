package pakang

// Support for alpine's `apk` utility

type ApkPM struct {
	extraflags []string
}

func NewApkPM(flags []string) ApkPM {
	return ApkPM{flags}
}

func (self ApkPM) Name() string { return "apk (alpine)" }

func (self ApkPM) Help() []string {
	return []string{
		"clean : Clean the cache",
	}
}

func (self ApkPM) Search(terms []string) {
	cmd := []string{"apk", "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (self ApkPM) Extra(terms []string) {
	if ArrayHas("clean", self.extraflags) {
		self.clean()
	} else {
		self.Search(terms)
	}
}

func (self ApkPM) clean() {
	RunCmd(NEED_ROOT, "apk", "clean").OrFail("Auto clean failed")
}

func (self ApkPM) Show(pkg string) {
	RunCmd(0, "apk", "info", pkg).OrFail("Error")
}

func (self ApkPM) Update() {
	RunCmd(NEED_ROOT, "apk", "update").OrFail("Could not update package index")
}

func (self ApkPM) Install(yes bool, packages []string) {
	cmd := []string{"apk", "add"}
	if yes {
		cmd = append(cmd, "-y")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install operation failed")
}

func (self ApkPM) Remove(packages []string) {
	cmd := []string{"apk", "del"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Package removal failed")
}

func (self ApkPM) Upgrade(yes bool) {
	cmd := []string{"apk", "upgrade"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Upgarde failed")
}
