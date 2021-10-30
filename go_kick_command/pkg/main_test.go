package main

import (
	"cmdctl/pkg/cmd"
	"cmdctl/pkg/docker"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	SuccessfulExitCode = 0
	FailureExitCode    = 1
)

func Test_ExecuteOsCommandFromWithinGolang(t *testing.T) {
	testName := "ExecuteOsCommandFromWithinGolang"
	t.Run(testName, func(t *testing.T) {
		exitCode, stdoutStderr, _ := cmd.ExecOsCommand("docker exec busybox ls -l -a")
		assert.Equal(t, exitCode, SuccessfulExitCode)
		assert.Contains(t, stdoutStderr, ".dockerenv")
		assert.Contains(t, stdoutStderr, "bin")

		exitCode, stdoutStderr, _ = cmd.ExecOsCommand("failedToThisCommand")
		assert.Equal(t, exitCode, 127)
		assert.Contains(t, stdoutStderr, "sh: 1: failedToThisCommand: not found")
	})
}

func Test_SetEnvironmentVariablesInContainer(t *testing.T) {
	testName := "SetEnvironmentVariablesInContainer"
	t.Run(testName, func(t *testing.T) {
		containerName := "busybox"
		envKeyValueMap := map[string]string{
			"AZURESTORAGEACCOUNT": "develop-st-jpeast",
			"AZUREQUEUESTORAGE":   "develop-queue-jpeast",
		}
		cmdExecMainProcess := "env"
		testCmd := docker.CreateDockerExecCommand(containerName, envKeyValueMap, cmdExecMainProcess)
		exitCode, stdoutStderr, _ := cmd.ExecOsCommand(testCmd)
		assert.Equal(t, exitCode, SuccessfulExitCode)
		assert.Contains(t, stdoutStderr, "develop-queue-jpeast")
	})
}

// func SetEnvInContainer(containerName string, keyValueMap map[string]string) error {
// 	for key, value := range keyValueMap {
// 		cmdSetEnv := fmt.Sprintf("docker exec %s sh -c 'export %s=%s'", containerName, key, value)
// 		exitCode, stdoutStderr, err := cmd.ExecOsCommand(cmdSetEnv)
// 		if err != nil {
// 			return fmt.Errorf("exit code = %d, stdout or stderr = %s, err = %s", exitCode, stdoutStderr, err.Error())
// 		}
// 	}
// 	return nil
// }

// envKeyValueMap := map[string]string{
// 	"AZURESTORAGEACCOUNT": "develop-st-jpeast",
// 	"AZUREQUEUESTORAGE":   "develop-queue-jpeast",
// }
// if err := SetEnvInContainer(containerName, envKeyValueMap); err != nil {
// 	t.Errorf("Failed to set environment variables in container. %v", err)
// 	return
// }

// func Test_SetEnvironmentVariablesInContainer(t *testing.T) {
// 	testName := "SetEnvironmentVariablesInContainer"
// 	t.Run(testName, func(t *testing.T) {
// 		exitCode, stdoutStderr, _ := cmd.ExecOsCommand("docker exec busybox sh -c '/main.sh'")
// 		assert.Equal(t, exitCode, SuccessfulExitCode)
// 		assert.Contains(t, stdoutStderr, "develop-queue-jpeast")
// 	})
// }
