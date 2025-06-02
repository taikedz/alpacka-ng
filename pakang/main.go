package pakang

import (
    "fmt"
    "strings"

    "github.com/taikedz/goargs/goargs"
)

func Fail(code int, message string, err error) {
    if err == nil {
        fmt.Println(message)
    } else {
        fmt.Printf("%s : %v\n", message, err)
    }
    os.Exit(code)
}

func Main(progname string) {
    modes_help := []string{
    "[OPTS] TERMS : search for terms in package descriptions",
    "actions:",
    "[OPTS] {--install|-i} PACKAGES : install specified packages",
    "[OPTS] {--remove|-r} PACKAGES : remove specified packages",
    "[OPTS] {--upgrade|-g} : upgrade packages on system",
    "[OPTS] {--show|-s} PACKAGE : show information on given package",
    "[OPTS] {--install-manifest|-M} MANIFEST : install from a manifest file"
    }
    opts_help := []string {
    "-i , -r, -g, and -M, are mutually exclusive - use only one at a time.",
    "",
    "OPTS:",
    "{--update-index|-u} : update package index before action",
    "{--yes|-y} : automatically accept (for -i, -g)",
    "{--extra|-x} EXTRAFLAGS : extra flag, for custom package manager support. Can be specified multiple times.",
    }

    var final_help []string
    for _,hstr := range(modes_help) {
        final_help = append(final_help, fmt.Sprintf("%s %s\n", progname, hstr))
    }
    final_help = append(final_help, strings.Join(opts_help, "\n"))

    parser := goargs.NewParser(strings.Join(final_help, "\n"))
    parser.ParseCliArgs()

    update := parser.Bool("update-index", false, "Update the package index (supported package managers)")
    parser.SetShortFlag('u', "update-index")
    yes := parser.Bool("yes", false, "Automatcially accept")
    parser.SetShortFlag('y', "yes")
    extraflags := parser.Appender("extra", "", "An extra standalone flag, can be specified multiple times; e.g. --extra=Frobulate --extra This=That")
    parser.SetShortFlag('x', "extra")

    install := parser.Bool("install", false, "Install the packages")
    parser.SetShortFlag('i', "install")
    remove := parser.Bool("remove", false, "Remove the packages")
    parser.SetShortFlag('r', "remove")
    upgrade := parser.Bool("upgrade", false, "Upgrade system packages")
    parser.SetShortFlag('g', "upgrade")
    show := parser.Bool("show", false, "Show info for package")
    parser.SetShortFlag('s', "show")
    manifest := parser.String("install-manifest", "", "Install from manifest file")
    parser.SetShortFlag('M', "install-manifest")

    assert_OnlyOneUsed(install, remove, upgrade, manifest, show)
    assert_FlagNotUsedWithMode("yes", yes, remove, show)

    parser.ParseCliArgs()

    pman := GetPackageManager(extraflags)

    if *update {
        pman.Update()
    }
    
    if *install {
        pman.Install(yes, parser.Args())
    } else if *upgrade {
        pman.Upgrade(yes)
    } else if *remove {
        pman.Remove(parser.Args())
    } else if *show {
        var pkg string
        chopmOne("packages", &pkg)
        pman.Show(pkg)
    } else if *manifest {
        var manifestfile string
        chompOne("manifest files", &manifestfile)
        //installManifest(&pman, manifestfile) // TODO
    ]
}

func chompOne(name string, ref interface{}) {
    remains, err := parser.UnpackArgs(0, ref)
    if err != nil {
        Fail(1, "Internal error", err)
    }
    if len(remains) > 0 {
        Fail(1, fmt.Sprintf("Too many %s specified", name), nil)
    }
}

func assert_OnlyOneUsed(flags ... *bool) {
    found_prior := false
    for _,flag in range(flags) {
        if flag {
            if found_prior {
                Fail(1, "Mode set more than once.", nil)
            }
            found_prior = true
        }
    }
}

func assert_FlagNotUsedWithMode(flagname string, flag_set *bool, modes ... *bool) {
    mode_set := false
    for _,mode range(modes) {
        mode_set = mode_set || mode
    }
    if mode_set && flag_set {
        Fail(1, fmt.Sprintf("%s used with incompatible mode", flagname), nil)
    }
}
