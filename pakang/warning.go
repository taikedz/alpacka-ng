package pakang

import (
	"fmt"
	"os"
)

const warning_path string = "/etc/alpacka/warnings"

func SetWarning(name string, message string) error {
	if !IsRootUser() {
		return fmt.Errorf("root user is required to set warnings")
	}
	if !fileExists(warning_path) {
		// https://gosamples.dev/create-directory/
		err := os.MkdirAll(warning_path, 0755)
		FailIf(err, 1, "Could not create global warnings dir")
	}
	return writeWarningFile(fileForWarning(name), message)
}

func GetWarning(name string) (string, error) {
	warnfile := fileForWarning(name)
	if fileExists(warnfile) {
		return readWarningFile(warnfile)
	} else {
		return "", nil
	}
}

func fileForWarning(name string) string {
	return fmt.Sprintf("%s/%s", warning_path, name)
}

func readWarningFile(filepath string) (string, error) {
	// https://golangdocs.com/reading-files-in-golang
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func writeWarningFile(filepath string, data string) error {
	// https://gobyexample.com/writing-files
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	f.Sync()

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	// Note if there is a stat error we just assume it does not exist
	//   - not structly true, but good enough for us.
	return err == nil
}
