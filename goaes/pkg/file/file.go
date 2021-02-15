// package ファイル処理向けパッケージ
package file

import (
	"io/ioutil"
	"os"
)

func Read(filepath string) ([]byte, error) {
	plaintext, err := ioutil.ReadFile(filepath)
	return plaintext, err
}

func Write(filepath string, writedata []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filepath, writedata, perm)
}
