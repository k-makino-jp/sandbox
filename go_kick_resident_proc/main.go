package main

import (
	"fmt"
	"integrationtest/cmd"
	"log"
	"time"
)

func main() {
	// execCmd := "./script/sleep.sh"
	// execCmd := "echo start > test.txt; sleep 5; echo end"
	execCmd := "echo start; sleep 100; echo end"
	exitCode, stdoutStderr, err := cmd.ExecOsCommandWithTimeout(execCmd, nil, 3*time.Second)
	if err != nil {
		log.Fatalln("FATAL:", exitCode, stdoutStderr, err)
	}
	fmt.Println("=== exit code: ", exitCode)
	fmt.Println("=== stdout")
	fmt.Println(stdoutStderr)
}
