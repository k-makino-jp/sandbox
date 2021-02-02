package enver

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Enver is interface
type Enver interface {
	Initialize()
	GetEnv() ModelEnv
}

// ModelEnv is struct
type ModelEnv struct {
	Endpoint string `envconfig:"ENDPOINT"`
}

// EnvImpl is struct (implements Enver)
type EnvImpl struct {
	ModelEnv
}

// Initialize is function
func (e *EnvImpl) Initialize() {
	err := envconfig.Process("", &e.ModelEnv)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GetEnv is getter
func (e *EnvImpl) GetEnv() ModelEnv {
	return e.ModelEnv
}

// NewEnvImpl creates instance
func NewEnvImpl() *EnvImpl {
	return &EnvImpl{}
}

func HowToUsePkgEnv() {
	e := NewEnvImpl()
	e.Initialize()
	env := e.GetEnv()
	fmt.Println(env)
}
