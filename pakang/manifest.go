package pakang

import (
	"os"
	"fmt"
	"gopkg.in/yaml.v3"
)

/* Load a YAML manifest file
 * Filter out based on os-release content '>=', '<=', '==', '=~' (substring)
 * Combine all package names found
 * Pass to Install(packages)
 */

/*
Choice of YAML package:

Multiple tutorials reference `gopkg.in/yaml.v3`
I'm following this one for example: https://zetcode.com/golang/yaml/

I had no idea who runs that site (gopkg.in), and who publishes that package.

Apparently gopkg.in has been around sice the early days of Go - so much so that the go tooling
itself does a _special_ version recognition allowing gopkg.in to specify `/yaml.v3` whereas
all other sources need to specify as a subdir: `/yaml/v3`, Github included.
This is because gopkg.in is not a package repository itself, but a URL convetionizer from before
the days of Go Modules. Go's package version resolver seems to have stabilised with a couple
conventions I need still to explore...

At time of looking up (June 2025) it is run by an employee at Canonical - the WhoIs lookup confirms this:
https://www.whois.com/whois/gopkg.in

The yaml package itself eventually leads to here: https://github.com/go-yaml/yaml

gopkg.in and the yaml package originated with Gustavo Niemeyer , and the yaml package was archived in April this year (!)

What a palava trying to establish chain of trust... Lazy-me will just accept `gopkg.in/yaml` as being a Canonical project
but it still grates me...
*/

type ReleaseSpec struct {
	Release string // to allow shorthands, we treat this as a plain string
	Groups string
}

func LoadManifest(mfest_path string) {
	data, err := os.ReadFile(mfest_path)
	FailIf(err, 1, "Could not read manifest %s", mfest_path)

	release := ReleaseSpec{}
	yaml.Unmarshal(data, &release)
	fmt.Printf("%#v\n", release)
}