package pakang

import (
    "fmt"
    "os"
    "strings"
    "slices"

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

func Main() {
    progname := os.Args[0]
    parser := goargs.NewParser(fmt.Sprintf("%s - Unified package manager command\n\n%s [OPTS] [PACKAGES ...]\n\nOPTS:", progname, progname))

    update := parser.Bool("update-index", false, "Update the package index (relevant package managers)")
    parser.SetShortFlag('u', "update-index")
    yes := parser.Bool("yes", false, "Automatcially accept (install/upgrade)")
    parser.SetShortFlag('y', "yes")

    op_modes := map[rune]string{
        'S': "search",
        'i': "install",
        'r': "remove",
        'g': "upgrade",
        's': "show",
        'm': "manifest",
    }
    mode := parser.Mode("action", "search", op_modes, "Action")
    manifestfile := parser.String("manifest", "", "Manifest file path (requires '-m' mode)")
    parser.SetShortFlag('M', "manifest")

    var extraflags *[]string
    pman := GetPackageManager(nil)
    pman_help := pman.Help()
    if len(pman_help) > 0 {
        extraflags = parser.Appender("extra", "Custom flags for system-specific package manager")
        parser.SetShortFlag('x', "extra")
        pman_help = slices.Insert(pman_help, 0, "", "Extra flags (-x|--extra)")
        parser.SetPostHelptext(strings.Join(pman_help, "\n"))
    }

    parser.SetHelpOnEmptyArgs(true)

    if err := parser.ParseCliArgs(); err != nil {
        Fail(1, "Invalid args", err)
    }

    pman = GetPackageManager(*extraflags)

    if *update {
        pman.Update()
    }

    switch *mode {
    case "install":
        pman.Install(*yes, parser.Args())
    case "remove":
        pman.Remove(parser.Args())
    case "upgrade":
        pman.Upgrade(*yes)
    case "show":
        for _, pkg := range parser.Args() {
            pman.Show(pkg)
        }
    case "manifest":
        if len(*manifestfile) == 0 {
            Fail(1, "No manifest file specified", nil)
        }
        //installManifest(pman, manifestfile) // TODO
    default:
        pman.Extra(parser.Args())
    }
}
