// Package kubernetes Kubernetes REST API処理用パッケージ
// ---
// mock作成コマンド
// mockgen -source=clientgo_rest.go -destination=mock_kubernetes/mock_clientgo_rest.go -package=mock_kubernetes
package kubernetes

import (
	"k8s.io/client-go/rest"
)

// Kubernetes Kubernetes REST API処理用インターフェース
type clientGoRest interface {
	InClusterConfig() (*rest.Config, error)
}

type clientGoRestImpl struct{}

func (r clientGoRestImpl) InClusterConfig() (*rest.Config, error) {
	return rest.InClusterConfig()
}
