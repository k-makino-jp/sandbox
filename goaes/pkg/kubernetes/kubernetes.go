// Package kubernetes Kubernetes REST API処理用パッケージ
package kubernetes

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Kubernetes Kubernetes REST API処理用インターフェース
type Kubernetes interface{}

type kubernetes struct {
	clientset           k8s.Interface
	pod                 *corev1.Pod
	namespacePod        *corev1.Namespace
	nodeList            *corev1.NodeList
	namespaceKubeSystem *corev1.Namespace
}

func (k *kubernetes) getPods() error {
	// /api/v1/pods/{podInfo.podName}
	// https://pkg.go.dev/k8s.io/client-go/kubernetes/typed/core/v1#PodInterface
	// ctx := context.TODO()
	name := "podName"
	opts := metav1.GetOptions{
		TypeMeta:        metav1.TypeMeta{},
		ResourceVersion: "",
	}
	pod, err := k.clientset.CoreV1().Pods("namespace").Get(context.TODO(), name, opts)
	k.pod = pod
	return err
}

func (k *kubernetes) getNodeList() error {
	// /api/v1/nodes
	listOptions := metav1.ListOptions{
		TypeMeta:             metav1.TypeMeta{},
		LabelSelector:        "",
		FieldSelector:        "",
		Watch:                false,
		AllowWatchBookmarks:  false,
		ResourceVersion:      "",
		ResourceVersionMatch: "",
		TimeoutSeconds:       new(int64),
		Limit:                0,
		Continue:             "",
	}
	nodeList, err := k.clientset.CoreV1().Nodes().List(context.TODO(), listOptions)
	k.nodeList = nodeList
	return err
}

func (k *kubernetes) getNamespaceKubeSystem() error {
	// kube-system
	namespaceKubeSystem, err := k.clientset.CoreV1().Namespaces().Get(context.TODO(), "kube-system", metav1.GetOptions{})
	k.namespaceKubeSystem = namespaceKubeSystem
	return err
}

func (k *kubernetes) getNamespacePod() error {
	// kube-system
	namespacePod, err := k.clientset.CoreV1().Namespaces().Get(context.TODO(), "podnamespace", metav1.GetOptions{})
	k.namespacePod = namespacePod
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
func NewKubernetes() (*kubernetes, error) {
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
	return &kubernetes{clientset: clientset}, err
}
