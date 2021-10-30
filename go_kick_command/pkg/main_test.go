package main

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SuccessfulExitCode = 0
	FailureExitCode    = 1
)

// ExecOsCommand executes os command.
func ExecOsCommand(cmd string) (exitCode int, stdoutStderr string, err error) {
	execCmd := exec.Command("sh", "-c", cmd)
	stdoutStderrBytes, err := execCmd.CombinedOutput()
	return execCmd.ProcessState.ExitCode(), string(stdoutStderrBytes), err
}

func Test_main(t *testing.T) {
	testName := "ExecuteOsCommandFromWithinGolang"
	t.Run(testName, func(t *testing.T) {
		exitCode, stdoutStderr, _ := ExecOsCommand("docker exec busybox ls -l -a")
		assert.Equal(t, exitCode, SuccessfulExitCode)
		assert.Contains(t, stdoutStderr, ".dockerenv")
		assert.Contains(t, stdoutStderr, "bin")

		exitCode, stdoutStderr, _ = ExecOsCommand("failedToThisCommand")
		assert.Equal(t, exitCode, 127)
		assert.Contains(t, stdoutStderr, "sh: 1: failedToThisCommand: not found")
	})
}
