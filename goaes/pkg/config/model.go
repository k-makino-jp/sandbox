// package config 設定ファイル処理用パッケージ
package config

import (
	"encoding/json"
	"io/ioutil"
)

// Configer 設定ファイル処理用インターフェース
type Configer interface {
	Read()
	Get()
}

type config struct {
	Endpoint string `json:"endpoint"`
}

type configer struct {
	config config
}

// Read 設定ファイル読み込み関数
func (c *configer) Read(configFilePath string) error {
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, &c.config); err != nil {
		return err
	}
	return nil
}

// Get 設定ファイル情報取得関数
func (c *configer) Get() config {
	return c.config
}

// NewCongier コンストラクタ
func NewConfiger() *configer {
	return &configer{}
}
