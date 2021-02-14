// Repository
// Layer: Interface Adapter: GateWays
// Role: Data Lifecycle Operation (CRUD)
package repository

import "sandbox/goutil/pkg/03_entity/config"

type ConfigRepository interface {
	Read() *config.Config
}

type Config struct {
}

func (c Config) Read() *config.Config {
	return &config.Config{}
}
