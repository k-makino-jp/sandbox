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
	Stat(name string) (os.FileInfo, error)
	IsNotExist(err error) bool
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, data []byte, perm fs.FileMode) error
}

// Os implements OsInterface.
type Os struct {
}

// NewOs returns Os instance.
func NewOs() *Os {
	return &Os{}
}

// Stat wraps os.Stat.
func (o Os) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

// IsNotExist wraps os.IsNotExist.
func (o Os) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

// ReadFile wraps os.ReadFile.
func (o Os) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

// WriteFile wraps os.WriteFile.
func (o Os) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}
