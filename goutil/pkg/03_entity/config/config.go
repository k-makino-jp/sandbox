package config

type Config struct {
	endpoint string
}

func (c Config) GetEndpoint() string {
	return c.endpoint
}
