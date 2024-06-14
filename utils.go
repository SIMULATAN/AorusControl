package main

import (
	"fmt"
	"io"
	"os/exec"
)

func ExecCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	bytes, err := io.ReadAll(stdErr)
	if err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Exit code is %d\n", exitError.ExitCode())
		}
		fmt.Println(string(bytes[:len(bytes)-1]))
		return err
	}
	return nil
}
