package pakang

type PackageManager interface {
    Update(yes *bool)
    Install(yes *bool, packages []string)
    Remove(packages []string)
    Upgrade(yes *bool)
}
