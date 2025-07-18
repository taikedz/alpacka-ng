package pakang

import "fmt"

// update readme too ;-)
const VER_MAJ int = 1
const VER_MIN int = 1
const VER_PATCH int = 0

func GetVersionString() string {
	return fmt.Sprintf("v%d.%d.%d", VER_MAJ, VER_MIN, VER_PATCH)
}

func GetVersionInts() []int {
	return []int{VER_MAJ, VER_MIN, VER_PATCH}
}
