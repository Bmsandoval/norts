package utils

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

const ShellToUse = "bash"

func ExecGetOutput(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func ExecGetOutputArray(command string) (error, []string, string) {
	err, stdout, stderr := ExecGetOutput(command)

	trimmedOutput := strings.TrimSpace(stdout)
	splitOutput := strings.Split(trimmedOutput,"\n")

	return err, splitOutput, stderr
}

func Exec(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// required for things like ssh, kubectl-exec, and vim
func ExecNotCapturingOutput(command string, args []string) error {
	c := exec.Command(command, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	err := c.Run()
	if err != nil {
		return err
	}
	return nil
}
