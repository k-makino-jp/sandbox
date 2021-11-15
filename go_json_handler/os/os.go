// Package os implements a wrapper for os package.
// We can create os mock with below command.
//   mockgen -source=os/os.go -destination os/os_mock.go -package=os
package os

import (
	"io/fs"
	"os"
)

// OsInterface defines encoding/json methods.
type OsInterface interface {
	WriteFile(name string, data []byte, perm fs.FileMode) error
}

// Os implements OsInterface.
type Os struct {
}

// NewOs returns Os instance.
func NewOs() *Os {
	return &Os{}
}

// WriteFile wraps os.WriteFile.
func (o Os) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}
