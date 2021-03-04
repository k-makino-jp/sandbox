package main

import (
	"fmt"

	"k8s.io/client-go/rest"
	// go install k8s.io/client-go/rest@v0.20.3
	// go install golang.org/x/tools/gopls@latest
)

// update go.mod
//   go mod tidy: add missing and remove unused modules
//   go build command uses require version
// download module
//   go install k8s.io/client-go/rest@v0.20.3
// go build

// Kubernetes Kubernetes REST API処理用インターフェース
type clientGoRest interface {
	InClusterConfig() (*rest.Config, error)
}

type clientGoRestImpl struct{}

func (r clientGoRestImpl) InClusterConfig() (*rest.Config, error) {
	return rest.InClusterConfig()
}

func main() {
	fmt.Println("Hello world")
}
