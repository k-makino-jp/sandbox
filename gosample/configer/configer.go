package configer

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configer interface {
	Read()
	Get()
}

type ModelConfig struct {
	Endpoint string `json:"endpoint"`
}

type ConfigerImpl struct {
	configFilePath string
	config         ModelConfig
}

func (c *ConfigerImpl) Read() {
	bytes, err := ioutil.ReadFile(c.configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bytes, &c.config); err != nil {
		log.Fatal(err)
	}
}

func (c *ConfigerImpl) Get() ModelConfig {
	return c.config
}

func NewConfigerImpl(configFilePath string) *ConfigerImpl {
	return &ConfigerImpl{
		configFilePath: configFilePath,
	}
}
