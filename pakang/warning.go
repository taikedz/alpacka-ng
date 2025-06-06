package pakang

import (
	"os"
    "fmt"
    "io/ioutil"
)

const warning_path string = "/etc/alpacka/warnings"

func SetWarning(name string, message string) error {
	if ! IsRootUser() {
		return fmt.Errorf("Root user is required to set warnings.")
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
	warnfile := fmt.Sprintf("%s/%s", warning_path, name)
}
 
func readWarningFile(filepath string) (string, error) {
	// https://golangdocs.com/reading-files-in-golang
    content, err := ioutil.ReadFile(filepath)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

func writeWarningFile(filepath string, data string) error {
	// https://gobyexample.com/writing-files
	f, err := os.Create(filepath)
	if err != nil { return err }
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil { return err }
	f.Sync()

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	// Note if there is a stat error we just assume it does not exist
	//   - not structly true, but good enough for us.
	return err == nil
}