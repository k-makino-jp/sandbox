// Package kubernetes Kubernetes REST API処理用パッケージ
package kubernetes

import (
	"context"
	"gosample/pkg/downwardapi"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
)

const (
	namespaceKubeSystem = "kube-system"
)

// Kubernetes Kubernetes REST API処理用インターフェース
type Kubernetes interface{}

type kubernetes struct {
	clientset           k8s.Interface
	podInfo             downwardapi.PodInfo
	pod                 *corev1.Pod
	namespacePod        *corev1.Namespace
	nodeList            *corev1.NodeList
	namespaceKubeSystem *corev1.Namespace
}

// https://pkg.go.dev/k8s.io/client-go/kubernetes/typed/core/v1#PodInterface
func (k *kubernetes) getPods() (err error) {
	// GET /api/v1/namespaces/{namespace}/pods/{name}
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#-strong-read-operations-pod-v1-core-strong-
	apiCaller := func() error {
		k.pod, err = k.clientset.CoreV1().Pods(k.podInfo.Namespace).Get(context.TODO(), k.podInfo.PodName, metav1.GetOptions{})
		return err
	}
	err = retry.RetryOnConflict(retry.DefaultRetry, apiCaller)
	return err
}

func (k *kubernetes) getNodeList() (err error) {
	// GET /api/v1/nodes
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#list-node-v1-core
	apiCaller := func() error {
		k.nodeList, err = k.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		return err
	}
	err = retry.RetryOnConflict(retry.DefaultRetry, apiCaller)
	return err
}

func (k *kubernetes) getNamespaceKubeSystem() (err error) {
	// GET /api/v1/namespaces/{name}
	// name: name of the kube-system namespace
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#-strong-read-operations-namespace-v1-core-strong-
	apiCaller := func() error {
		k.namespaceKubeSystem, err = k.clientset.CoreV1().Namespaces().Get(context.TODO(), namespaceKubeSystem, metav1.GetOptions{})
		return err
	}
	err = retry.RetryOnConflict(retry.DefaultRetry, apiCaller)
	return err
}

func (k *kubernetes) getNamespacePod() (err error) {
	// GET /api/v1/namespaces/{name}
	// name: name of the pod namespace
	// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#-strong-read-operations-namespace-v1-core-strong-
	apiCaller := func() error {
		k.namespacePod, err = k.clientset.CoreV1().Namespaces().Get(context.TODO(), k.podInfo.Namespace, metav1.GetOptions{})
		return err
	}
	err = retry.RetryOnConflict(retry.DefaultRetry, apiCaller)
	return err
}

// GetInfo Kubernetes情報取得関数
func (k kubernetes) GetInfo() (Info, error) {
	if err := k.getPods(); err != nil {
		return Info{}, err
	}
	if err := k.getNodeList(); err != nil {
		return Info{}, err
	}
	if err := k.getNamespaceKubeSystem(); err != nil {
		return Info{}, err
	}
	if err := k.getNamespacePod(); err != nil {
		return Info{}, err
	}
	return Info{}, nil
}

// NewKubernetes コンストラクタ
func NewKubernetes(podInfo downwardapi.PodInfo) (*kubernetes, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &kubernetes{
		clientset: clientset,
		podInfo:   podInfo,
	}, err
}
