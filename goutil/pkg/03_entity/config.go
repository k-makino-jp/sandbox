// package entity 暗号化向けパッケージ
package entity

type Config struct {
	endpoint string
}

func (c Config) GetEndpoint() string {
	return c.endpoint
}
