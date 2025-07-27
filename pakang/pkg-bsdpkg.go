package pakang

/* Initial implementation intends to manage _packages_ , not "ports"
 */

type FreebsdPkgPM struct {
	extraflags []string
}

func NewFreebsdPkgPM(flags []string) FreebsdPkgPM {
	return FreebsdPkgPM{flags}
}

func (pm FreebsdPkgPM) Name() string { return "zypper" }

func (pm FreebsdPkgPM) Help() []string {
	return []string{
		// FIXME:TODO
		// https://docs.freebsd.org/en/books/handbook/ports/#pkgng-intro
		"setup : perform initial setup and configuration of pkg(8)",
		"audit : perform security audit",
		"fetch=$src : fetch packages from stated source",
	}
}

func (pm FreebsdPkgPM) Search(terms []string) {
	if len(terms) == 0 {
		return
	}
	cmd := []string{"pkg", "search"}
	cmd = append(cmd, terms...)
	RunCmd(0, cmd...).OrFail("Search failed")
}

func (pm FreebsdPkgPM) NoAction(terms []string) {
	// FIXME:TODO
	pm.Search(terms)
}

func (pm FreebsdPkgPM) Clean() {
	RunCmd(NEED_ROOT, "pkg", "clean", "-a").OrFail("Clean failed")
}

func (pm FreebsdPkgPM) Show(pkg string) {
	// By default we prevent auto-update of package info - this is available with paf's own `-u` operation
	//  and allows accessing information without the delay of an automatic fetch
	// `paf minetest -s` : shows info on minetest, does not update indicies
	// `paf minetest -su` : same, with prior index update
	RunCmd(0, "pkg", "rquery", "-U", pkg).OrFail("Show package failed")
}

func (pm FreebsdPkgPM) Update() {
	RunCmd(NEED_ROOT, "pkg", "update", "-y").OrFail("Package index update failed")
}

func (pm FreebsdPkgPM) Install(yes bool, packages []string) {
	cmd := []string{"pkg", "install"}
	if yes {
		cmd = append(cmd, "-y")
	}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Install failed")
}

func (pm FreebsdPkgPM) Remove(packages []string) {
	cmd := []string{"pkg", "delete"}
	cmd = append(cmd, packages...)
	RunCmd(NEED_ROOT, cmd...).OrFail("Remove failed")
}

func (pm FreebsdPkgPM) Upgrade(yes bool) {
	cmd := []string{"pkg", "upgrade"}
	if yes {
		cmd = append(cmd, "-y")
	}
	RunCmd(NEED_ROOT, cmd...).OrFail("Upgrade failed")
}
