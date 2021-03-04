// Package kubernetes Kubernetes REST API処理用パッケージ
package kubernetes

import (
	"errors"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	overFlowText     = "An overflow occurred"
	infixContainerID = "://"
	cgroupFilePath   = "/proc/self/cgroup"
)

// Info Kubernetes情報
type Info struct {
	Pod                 *corev1.Pod
	NamespacePod        *corev1.Namespace
	NodeList            *corev1.NodeList
	NamespaceKubeSystem *corev1.Namespace
	ioUtil              ioUtil
}

// GetNamespaceKubeSystemID kube-system取得
func (i Info) GetNamespaceKubeSystemID() string {
	return string(i.NamespaceKubeSystem.UID)
}

// CalculateTotalVCPUOfNodeList Node vCPU数合計計算
func (i Info) CalculateTotalVCPUOfNodeList() int64 {
	var totalVCPU resource.Quantity
	for _, node := range i.NodeList.Items {
		totalVCPU.Add(*node.Status.Capacity.Cpu())
	}
	// (A) Maximum value for an int64: 9,223,372,036,854,775,807
	// (B) Kubernetes supports clusters with up to 5,000 nodes
	// (C) Maximum logical cpus of Red Hat Enterprise Linux 8: 8,192 [core] => 8,192,000 [milli core]
	// (D) = (B) + (C) = 5,000 * 8,192,000 = 40,960,000,000
	// (A) >> (D)
	return totalVCPU.MilliValue()
}

// GetNumberOfNodes Node数取得
func (i Info) GetNumberOfNodes() int {
	return len(i.NodeList.Items)
}

// GetNamespacePodID Namespace取得
func (i Info) GetNamespacePodID() string {
	return string(i.NamespacePod.UID)
}

// GetVCPUOfOwnNode Node vCPU数取得
func (i Info) GetVCPUOfOwnNode() int64 {
	var vcpu resource.Quantity
	for _, node := range i.NodeList.Items {
		if node.Name == i.Pod.Spec.NodeName {
			vcpu = *node.Status.Capacity.Cpu()
		}
	}
	return vcpu.MilliValue()
}

// fp, err := os.Open(cgroupFilePath)
// if err != nil {
// 	return "", err
// }
// defer fp.Close()
// scanner := bufio.NewScanner(fp)
// for scanner.Scan() {
// 	cgroupLine := scanner.Text()
// }
// if err := scanner.Err(); err != nil {
// 	return "", err
// }

// 0123456789012345678901234567890123456789012
// Alfa Bravo Charlie Delta Echo Foxtrot Golf

// GetOwnContainerID container ID取得
func (i Info) GetOwnContainerID() (string, error) {
	// read cgroup file
	cgroupBytes, err := i.ioUtil.ReadFile(cgroupFilePath)
	if err != nil {
		return "", err
	}
	for _, containerStatus := range i.Pod.Status.ContainerStatuses {
		// Container ID Format:       <container_runtime>://<container_id>
		// Infix:                                        ://
		// Last Index:                                   *
		// Last Index + Infix Length:                       *
		// Container ID After Infix:                        <container_id>
		lastIndex := strings.LastIndex(containerStatus.ContainerID, infixContainerID)
		containerIDAfterInfix := containerStatus.ContainerID[lastIndex+len(infixContainerID):]
		// validate container_id
		if strings.Contains(string(cgroupBytes), containerIDAfterInfix) {
			return containerStatus.ContainerID, nil
		}
	}
	return "", errors.New("hoge")
}

// GetOwnContainerName コンテナ名取得
func (i Info) GetOwnContainerName(containerID string) (string, error) {
	for _, containerStatus := range i.Pod.Status.ContainerStatuses {
		if containerID == containerStatus.ContainerID {
			return containerStatus.Name, nil
		}
	}
	return "", errors.New("No such container" + containerID)
}

// GetLimitVCPUOfOwnContainer limit vCPU数取得
func (i Info) GetLimitVCPUOfOwnContainer(containerName string) (int64, error) {
	for _, container := range i.Pod.Spec.Containers {
		if containerName == container.Name {
			limitVCPU := *container.Resources.Limits.Cpu()
			return limitVCPU.MilliValue(), nil
		}
	}
	return 0, errors.New("No such container" + containerName)
}
