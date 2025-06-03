package pakang

import (
    "fmt"
    "os"
    "os/exec"
)

const NEED_ROOT int = 1

func RunCmd(flags int, tokens ... string) (int, error) {
    return RunCmdOut(true, flags, tokens...)
}

func RunCmdOut(dump bool, flags int, tokens ... string) (int, error) {
	if len(tokens) < 1 {
        fmt.Printf("FATAL - tokens not supplied")
        os.Exit(1)
    }

    var t_cmd string
    var t_args []string

    if flags & NEED_ROOT == NEED_ROOT {
        // Hard-coding use of sudo for now. If we get cleverer in the future,
        //   deal with it starting from here ...
        t_cmd = "sudo"
        t_args = tokens
    } else {
        t_cmd = tokens[0]
        t_args = tokens[1:]
    }

    cmd := exec.Command(t_cmd, t_args...)
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
