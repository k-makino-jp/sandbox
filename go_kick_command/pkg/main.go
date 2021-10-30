package main

import (
	// "cmdctl/pkg/cmd"

	"fmt"
	"log"
	"os/exec"
)

// ExecOsCommand executes os command.
func ExecOsCommand(cmd string) (stdoutStderr string, err error) {
	execCmd := exec.Command("sh", "-c", cmd)
	stdoutStderrBytes, err := execCmd.CombinedOutput()
	return string(stdoutStderrBytes), err
}

func main() {
	stdoutStderr, err := ExecOsCommand("ls -l")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(stdoutStderr))

	stdoutStderr, err = ExecOsCommand("docker exec busybox ls -l -a")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(stdoutStderr))

}
