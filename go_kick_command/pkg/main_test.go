package main

import (
	"fmt"
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

func Test_ExecuteOsCommandFromWithinGolang(t *testing.T) {
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

func SetEnvInContainer(containerName string, keyValueMap map[string]string) error {
	for key, value := range keyValueMap {
		cmdSetEnv := fmt.Sprintf("docker exec %s sh -c 'export %s=%s'", containerName, key, value)
		exitCode, stdoutStderr, err := ExecOsCommand(cmdSetEnv)
		if err != nil {
			return fmt.Errorf("exit code = %d, stdout or stderr = %s, err = %s", exitCode, stdoutStderr, err.Error())
		}
	}
	return nil
}

// envKeyValueMap := map[string]string{
// 	"AZURESTORAGEACCOUNT": "develop-st-jpeast",
// 	"AZUREQUEUESTORAGE":   "develop-queue-jpeast",
// }
// if err := SetEnvInContainer(containerName, envKeyValueMap); err != nil {
// 	t.Errorf("Failed to set environment variables in container. %v", err)
// 	return
// }

func Test_SetEnvironmentVariablesInContainer(t *testing.T) {
	testName := "SetEnvironmentVariablesInContainer"
	t.Run(testName, func(t *testing.T) {
		exitCode, stdoutStderr, _ := ExecOsCommand("docker exec busybox sh -c '/main.sh'")
		assert.Equal(t, exitCode, SuccessfulExitCode)
		assert.Contains(t, stdoutStderr, "develop-queue-jpeast")
	})
}
