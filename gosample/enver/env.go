package env

import (
	"github.com/kelseyhightower/envconfig"
)

type env struct {
	Endpoint string `envconfig:"ENDPOINT" required:"true"`
	Level    string `envconfig:"LEVEL" default:"INFO"`
}

var Env env

func (e *env) hasValidLevel() {
	switch e.Level {
	case "INFO":
		return
	default:
		e.Level = "INFO"
		return
	}
}

func Init() error {
	err := envconfig.Process("", &Env)
	if err != nil {
		return err
	}
	Env.hasValidLevel()
	return nil
}
