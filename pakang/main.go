package pakang

import (
    "fmt"
    "os"
    "strings"
    "slices"

    "github.com/taikedz/goargs/goargs"
)

func Main(progname string) {
    parser := goargs.NewParser(fmt.Sprintf("%s - Unified package manager command\n\n%s [OPTS] [PACKAGES ...]\n\nOPTS:", progname, progname))

    update := parser.Bool("update-index", false, "Update the package index (relevant package managers)")
    parser.SetShortFlag('u', "update-index")
    yes := parser.Bool("yes", false, "Automatcially accept (install/upgrade)")
    parser.SetShortFlag('y', "yes")
    print_version := parser.Bool("version", false, "Show version and exit")
    parser.SetShortFlag('V', "version")

    warning_message := parser.String("warning-message", "", "A warning message")
    parser.SetShortFlag('W', "warning-message")
    warning_action := parser.Choices("warning-action", []string{"", "install", "upgrade", "remove", "manifest"}, "Action for the warning message")
    parser.SetShortFlag('A', "warning-action")
    override_warning := parser.Bool("ignore-warnings", false, "Use this flag to ignore the warning.")

    op_modes := map[rune]string{
        'S': "search",
        'i': "install",
        'r': "remove",
        'g': "upgrade",
        's': "show",
        'm': "manifest",
        'w': "warn",
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

    if *print_version {
        fmt.Printf("%s %s\n", progname, VERSION)
        return
    }

    pman = GetPackageManager(*extraflags)

    switch *mode {
    case "install":
        WarningCheck(*mode, *override_warning)
        if *update { pman.Update() }
        pman.Install(*yes, parser.Args())
    case "remove":
        WarningCheck(*mode, *override_warning)
        if *update { pman.Update() }
        pman.Remove(parser.Args())
    case "upgrade":
        WarningCheck(*mode, *override_warning)
        if *update { pman.Update() }
        pman.Upgrade(*yes)
    case "show":
        for _, pkg := range parser.Args() {
            pman.Show(pkg)
        }
    case "manifest":
        WarningCheck("install", *override_warning)
        if len(*manifestfile) == 0 {
            Fail(1, "No manifest file specified", nil)
        }
        if *update { pman.Update() }
        installManifest(pman, manifestfile)
    case "warn":
        doWarningAction(*warning_message, *warning_action)
    default:
        pman.Extra(parser.Args())
    }
}

func doWarningAction(message, action string) {
    if action != "" {
        if message == "" {
            text, err := GetWarning(action)
            if err != nil {
                Fail(1, "Could not read warning file", err)
            }
            if text != "" {
                fmt.Printf("Warning for %s:\n%s\n", action, text)
            }
            os.Exit(0)
        }
        
        if message == "." {
            message = ""
        }

        err := SetWarning(action, message)
        if err != nil {
            Fail(1, "Could not set warning", err)
        }
    } else if message != "" {
        Fail(1, "Action -A must be specified to set a warning with -W", nil)
    } else {
        Fail(1, "-w mode specified without -A action or -W message", nil)
    }
}

func installManifest(pman PackageManager, manifest_path string) {
    manifest := LoadManifest(manifest_path)
    packs := manifest.GetPackages()
    pman.Install(true, packs)
}

func WarningCheck(name string, override_warning bool) {
    warntext, err := GetWarning(name)
    if err != nil {
        Fail(1, "Failed to read warning file", err)
    }
    if warntext == "" {
        // No warning
        return
    }
    // A warning was found - print it, and abort.
    fmt.Printf("!!! WARNING: %s\n\n", warntext)
    if !override_warning {
        Fail(10, "Abort", nil)
    }
}
