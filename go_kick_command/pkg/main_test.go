package main

import (
	"cmdctl/pkg/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Validation of stdout / stderr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, "HelloWorld", "World")
			stdoutStderr, err := cmd.ExecOsCommand("docker exec busybox ls -l -a")
			assert.NoError(t, err)
			assert.Contains(t, stdoutStderr, ".dockerenv")
			assert.Contains(t, stdoutStderr, "bin")
		})
	}
}
