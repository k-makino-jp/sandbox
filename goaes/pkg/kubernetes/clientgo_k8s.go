// Package kubernetes Kubernetes REST API処理用パッケージ
// ---
// mock作成コマンド
// mockgen -source=clientgo_k8s.go -destination=mock_kubernetes/mock_clientgo_k8s.go -package=mock_kubernetes
package kubernetes

import (
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type clientGoKubernetes interface {
	NewForConfig(c *rest.Config) (*k8s.Clientset, error)
}

type clientGoKubernetesImpl struct{}

func (o clientGoKubernetesImpl) NewForConfig(c *rest.Config) (*k8s.Clientset, error) {
	return k8s.NewForConfig(c)
}
