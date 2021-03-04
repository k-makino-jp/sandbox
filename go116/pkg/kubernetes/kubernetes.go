// Package kubernetes Kubernetes REST API処理用パッケージ
package kubernetes

import (
	"context"
	"fmt"
	"gosample/pkg/downwardapi"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

const (
	namespaceKubeSystem = "kube-system"
)

// Kubernetes Kubernetes REST API処理用インターフェース
type Kubernetes interface{}

type kubernetes struct {
	clientset           k8s.Interface
	clientGoRest        clientGoRest
	clientGoKubernetes  clientGoKubernetes
	podInfo             downwardapi.PodInfo
	pod                 *corev1.Pod
	namespacePod        *corev1.Namespace
	nodeList            *corev1.NodeList
	namespaceKubeSystem *corev1.Namespace
}

// Init 初期化関数
func (k *kubernetes) Init() error {
	// use service account token
	config, err := k.clientGoRest.InClusterConfig()
	if err != nil {
		return err
	}
	k.clientset, err = k.clientGoKubernetes.NewForConfig(config)
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

// https://pkg.go.dev/k8s.io/client-go/kubernetes/typed/core/v1#PodInterface
// GET /api/v1/namespaces/{namespace}/pods/{name}
//   https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#-strong-read-operations-pod-v1-core-strong-
// error: IsForbidden
//   https://pkg.go.dev/k8s.io/apimachinery/pkg/api/errors#IsForbidden
func (k *kubernetes) getPods() (err error) {
	apiCaller := func() error {
		k.pod, err = k.clientset.CoreV1().Pods(k.podInfo.Namespace).Get(context.TODO(), k.podInfo.PodName, metav1.GetOptions{})
		if errors.IsForbidden(err) {
			fmt.Println("Forbidden")
		}
		return err
	}
	// func cannot be compared, so we dont use mock
	return retry.RetryOnConflict(retry.DefaultRetry, apiCaller)
}

// GET /api/v1/nodes
// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#list-node-v1-core
func (k *kubernetes) getNodeList() (err error) {
	apiCaller := func() error {
		k.nodeList, err = k.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		return err
	}
	return retry.RetryOnConflict(retry.DefaultRetry, apiCaller)
}

// GET /api/v1/namespaces/{name}
// name: name of the kube-system namespace
// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#-strong-read-operations-namespace-v1-core-strong-
func (k *kubernetes) getNamespaceKubeSystem() (err error) {
	apiCaller := func() error {
		k.namespaceKubeSystem, err = k.clientset.CoreV1().Namespaces().Get(context.TODO(), namespaceKubeSystem, metav1.GetOptions{})
		return err
	}
	return retry.RetryOnConflict(retry.DefaultRetry, apiCaller)
}

// GET /api/v1/namespaces/{name}
// name: name of the pod namespace
// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#-strong-read-operations-namespace-v1-core-strong-
func (k *kubernetes) getNamespacePod() (err error) {
	apiCaller := func() error {
		k.namespacePod, err = k.clientset.CoreV1().Namespaces().Get(context.TODO(), k.podInfo.Namespace, metav1.GetOptions{})
		return err
	}
	return retry.RetryOnConflict(retry.DefaultRetry, apiCaller)
}

// NewKubernetes コンストラクタ
func NewKubernetes(podInfo downwardapi.PodInfo) *kubernetes {
	return &kubernetes{
		clientset:           nil,
		clientGoRest:        clientGoRestImpl{},
		clientGoKubernetes:  clientGoKubernetesImpl{},
		podInfo:             podInfo,
		pod:                 &corev1.Pod{},
		namespacePod:        &corev1.Namespace{},
		nodeList:            &corev1.NodeList{},
		namespaceKubeSystem: &corev1.Namespace{},
	}
}
