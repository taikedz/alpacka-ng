package pakang

import (
    "os"
    "fmt"
    "strings"

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

type Manifest struct {
    Alpacka Sections
}

type Sections struct {
    Variants []Variant
    PackageGroups map[string][]string `yaml:"package-groups"`
}

type Variant struct {
    Release string // to allow shorthands, we treat this as a plain string
    Groups string
}

func LoadManifest(mfest_path string) Manifest {
    data, err := os.ReadFile(mfest_path)
    FailIf(err, 1, "Could not read manifest %s", mfest_path)

    manifest := Manifest{}
    yaml.Unmarshal(data, &manifest)
    return manifest
}

func (self Manifest) GetPackageGroups() []string {
    osr := LoadOsRelease()

    for _, variant := range self.Alpacka.Variants {
        comp_strs := SplitStringMultichar(variant.Release, ", ")
        comp_strs = ExcludeStr(comp_strs, []string{""})

        matched := true
        for _,str := range comp_strs {
            param, op, value := getComparison(str)
            switch op {
            case ">=":
                matched = matched && osr.ParamGteValueInts(param, value)
            case "<=":
                matched = matched && osr.ParamLteValueInts(param, value)
            case "==":
                matched = matched && osr.Param(param) == value
            case "=~":
                matched = matched && osr.ParamContains(param, value)
            default:
                Fail(1, fmt.Sprintf("Unrecognised release comparison '%s' (%s)", op, str), nil)
            }
        }
        if matched {
            groups := SplitStringMultichar(variant.Groups, ", ") // FIXME - use regex
            groups = ExcludeStr(groups, []string{""})
            return groups
        }
    }

    return nil
}

func getComparison(compstr string) (string, string, string) {
    for _,op := range []string{">=", "<=", "==", "=~"} {
        if strings.Index(compstr, op) > 0 {
            tokens := strings.SplitN(compstr, op, 2)
            return tokens[0], op, tokens[1]
        }
    }
    Fail(1, fmt.Sprintf("Could not split comparison: %s", compstr), nil)
    return "","","" //for compiler
}

func (self Manifest) GetPackages() ([]string, error) {
    pgroups := self.GetPackageGroups()
    if pgroups == nil { return nil, fmt.Errorf("No package groups matched") }
    packages := []string{}

    for name,packs := range self.Alpacka.PackageGroups {
        if ! ArrayHas(name, pgroups) { continue }
        for _,p := range packs {
            if ! ArrayHas(p, packages) {
                packages = append(packages, p)
            }
        }
    }

    return packages, nil
}
