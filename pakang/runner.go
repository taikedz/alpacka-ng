package pakang

import (
    "fmt"
    "os"
    "os/exec"
)

func RunCmd(tokens ... string) (int, error) {
    return RunCmdOut(true, tokens...)
}

func RunCmdOut(dump bool, tokens ... string) (int, error) {
	if len(tokens) < 1 {
        fmt.Printf("FATAL - tokens not supplied")
        os.Exit(1)
    }

    cmd := exec.Command(tokens[0], tokens[1:]...)
    if dump {
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
    }
    if err := cmd.Start(); err != nil {
        return -1, fmt.Errorf("Execution error: %v\n", err)
    }

    // https://stackoverflow.com/a/10385867/2703818
    if err := cmd.Wait(); err != nil {
        if exiterr, ok := err.(*exec.ExitError); ok {
            return exiterr.ExitCode(), exiterr
        } else {
            return -1, err
        }
    }

    return 0, nil
}
