// Package downwardapi Pod情報取得処理用パッケージ
// ---
// mock作成コマンド
// mockgen -self_package=downwardapi -source=ioutil.go -destination=ioutil_mock.go -package=downwardapi
package downwardapi

import "io/ioutil"

type ioUtil interface {
	ReadFile(filename string) ([]byte, error)
}

type ioUtilImpl struct{}

func (i ioUtilImpl) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
