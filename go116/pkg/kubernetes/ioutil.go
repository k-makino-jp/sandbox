// Package kubernetes a
// ---
// mock作成コマンド
// mockgen -self_package=kubernetes -source=ioutil.go -destination=ioutil_mock.go -package=kubernetes
package kubernetes

import "io/ioutil"

type ioUtil interface {
	ReadFile(filename string) ([]byte, error)
}

type ioUtilImpl struct{}

func (i ioUtilImpl) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
