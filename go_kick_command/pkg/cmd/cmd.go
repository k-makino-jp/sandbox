package cmd

import "os/exec"

// ExecOsCommand executes os command.
func ExecOsCommand(cmd string) (exitCode int, stdoutStderr string, err error) {
	execCmd := exec.Command("sh", "-c", cmd)
	stdoutStderrBytes, err := execCmd.CombinedOutput()
	return execCmd.ProcessState.ExitCode(), string(stdoutStderrBytes), err
}
