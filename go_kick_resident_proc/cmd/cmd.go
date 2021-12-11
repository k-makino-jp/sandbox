package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"
)

// ExecOsCommand executes os command.
func ExecOsCommand(cmd string, envKeyValueMap map[string]string) (exitCode int, stdoutStderr string, err error) {
	execCmd := exec.Command("sh", "-c", cmd)
	execCmd.Env = os.Environ()
	for key, value := range envKeyValueMap {
		keyValue := fmt.Sprintf("%s=%s", key, value)
		execCmd.Env = append(execCmd.Env, keyValue)
	}
	stdoutStderrBytes, err := execCmd.CombinedOutput()
	return execCmd.ProcessState.ExitCode(), string(stdoutStderrBytes), err
}

// ExecOsCommandWithTimeout executes os command considering timeout.
func ExecOsCommandWithTimeout(cmd string, envKeyValueMap map[string]string, timeout time.Duration) (exitCode int, stdoutStderr string, err error) {
	execCmd := exec.Command("sh", "-c", cmd)
	execCmd.Env = os.Environ()
	for key, value := range envKeyValueMap {
		keyValue := fmt.Sprintf("%s=%s", key, value)
		execCmd.Env = append(execCmd.Env, keyValue)
	}

	// when exec.CombinedOutput used and interrupt occurred, messages in stdout fails to get.
	// in order to avoid above problem, we use new buffer.
	var stdoutStderrBuffer bytes.Buffer
	execCmd.Stdout = &stdoutStderrBuffer
	execCmd.Stderr = &stdoutStderrBuffer

	go func() {
		err = execCmd.Run()
	}()

	// waiting for specified time
	t := time.NewTimer(timeout)
	<-t.C

	// kill parent process
	parentProcess, err := process.NewProcess(int32(execCmd.Process.Pid))
	if err != nil {
		return execCmd.ProcessState.ExitCode(), stdoutStderrBuffer.String(), err
	}
	if err := execCmd.Process.Kill(); err != nil {
		return execCmd.ProcessState.ExitCode(), stdoutStderrBuffer.String(), err
	}

	// kill child processes
	childProcesses, err := parentProcess.Children()
	if err != nil {
		return execCmd.ProcessState.ExitCode(), stdoutStderrBuffer.String(), err
	}
	for _, childProcess := range childProcesses {
		name, _ := childProcess.Name()
		fmt.Println("INFO: kill child process = ", name)
		if err := childProcess.Kill(); err != nil {
			return execCmd.ProcessState.ExitCode(), stdoutStderrBuffer.String(), err
		}
	}
	return execCmd.ProcessState.ExitCode(), stdoutStderrBuffer.String(), err
}

// ExecOsCommandWithTimeout executes os command considering timeout.
func ExecOsCommandWithTimeoutLinux(cmd string, envKeyValueMap map[string]string, timeout time.Duration) (exitCode int, stdoutStderr string, err error) {
	execCmd := exec.Command("sh", "-c", cmd)
	execCmd.Env = os.Environ()
	for key, value := range envKeyValueMap {
		keyValue := fmt.Sprintf("%s=%s", key, value)
		execCmd.Env = append(execCmd.Env, keyValue)
	}
	execCmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// when exec.CombinedOutput used and interrupt occurred, messages in stdout fails to get.
	// in order to avoid above problem, we use new buffer.
	var stdoutStderrBuffer bytes.Buffer
	execCmd.Stdout = &stdoutStderrBuffer
	execCmd.Stderr = &stdoutStderrBuffer

	go func() {
		err = execCmd.Run()
	}()

	// waiting for specified time
	t := time.NewTimer(timeout)
	<-t.C

	// kill parent and child processes
	if err := execCmd.Process.Kill(); err != nil {
		return execCmd.ProcessState.ExitCode(), stdoutStderrBuffer.String(), err
	}
	err := syscall.Kill(-execCmd.Process.Pid, syscall.SIGKILL)
	return execCmd.ProcessState.ExitCode(), stdoutStderrBuffer.String(), err
}
