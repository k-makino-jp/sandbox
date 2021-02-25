// Package kubernetes Kubernetes REST API処理用パッケージ
package kubernetes

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Info Kubernetes情報
type Info struct {
	Pod                 *corev1.Pod
	NamespacePod        *corev1.Namespace
	NodeList            *corev1.NodeList
	NamespaceKubeSystem *corev1.Namespace
}

// CalculateTotalVCPUOfNodeList Node vCPU数合計計算
func (i Info) CalculateTotalVCPUOfNodeList() (int64, error) {
	var totalVCPU resource.Quantity
	for _, node := range i.NodeList.Items {
		totalVCPU.Add(*node.Status.Capacity.Cpu())
	}
	// (A) Maximum value for an int64: 9,223,372,036,854,775,807
	// (B) Kubernetes supports clusters with up to 5,000 nodes
	// (C) Maximum logical cpus of Red Hat Enterprise Linux 8: 8,192 [core] => 8,192,000 [milli core]
	// (D) = (B) + (C) = 5,000 * 8,192,000 = 40,960,000,000
	// (A) >> (D)
	value, isConvertible := totalVCPU.AsInt64()
	if !isConvertible {
		return 0, errors.New("Overflow")
	}
	return value, nil
}

// GetVCPUOfOwnNode Node vCPU数取得
func (i Info) GetVCPUOfOwnNode(podName string) (int64, error) {
	var vcpu resource.Quantity
	for _, node := range i.NodeList.Items {
		if node.Name == i.Pod.Spec.NodeName {
			vcpu = *node.Status.Capacity.Cpu()
		}
	}
	value, isConvertible := vcpu.AsInt64()
	if !isConvertible {
		return 0, errors.New("Overflow")
	}
	return value, nil
}

// GetLimitVCPUOfOwnContainer limit vCPU数取得
func (i Info) GetLimitVCPUOfOwnContainer(podName string) (int64, error) {
	var limitVCPU resource.Quantity
	for _, container := range i.Pod.Spec.Containers {
		limitVCPU = *container.Resources.Limits.Cpu()
	}
	value, isConvertible := limitVCPU.AsInt64()
	if !isConvertible {
		return 0, errors.New("Overflow")
	}
	return value, nil
}

// GetOwnContainerID container ID取得
func (i Info) GetOwnContainerID(podName string) []string {
	var containerIDList []string
	for _, containerStatus := range i.Pod.Status.ContainerStatuses {
		containerIDList = append(containerIDList, containerStatus.ContainerID)
	}
	return containerIDList
}

// GetKubernetesID kube-system namespace UID 取得
func (i Info) GetKubernetesID() string {
	return string(i.NamespaceKubeSystem.UID)
}
