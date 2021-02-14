// package usecase 暗号化向けパッケージ
package usecase

type ConfigInteractor struct {
	configRepository configRepository
}

func (c *ConfigInteractor) Read() {
	c.configRepository.Read()
}
