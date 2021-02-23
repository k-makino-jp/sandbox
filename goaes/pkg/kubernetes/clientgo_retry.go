// Package kubernetes Kubernetes REST API処理用パッケージ
// ---
// mockgen -source=clientgo_retry.go -destination=mock_kubernetes/mock_clientgo_retry.go -package=mock_kubernetes
package kubernetes

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
)

type clientGoRetry interface {
	RetryOnConflict(backoff wait.Backoff, fn func() error) error
}

type clientGoRetryImpl struct{}

func (r clientGoRetryImpl) RetryOnConflict(backoff wait.Backoff, fn func() error) error {
	return retry.RetryOnConflict(backoff, fn)
}
