// Package kubernetes Kubernetes REST API処理用パッケージ

package kubernetes

import (
	"reflect"
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func TestInfo_CalculateTotalVCPUOfNodeList(t *testing.T) {
	tests := []struct {
		name    string
		i       Info
		want    int64
		wantErr error
	}{
		{
			name:    "Info.CalculateTotalVCPUOfNodeList 合計vCPU数計算 合計vCPU数が返ること",
			i:       Info{NodeList: &corev1.NodeList{}},
			want:    1000,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.CalculateTotalVCPUOfNodeList()
			if err != tt.wantErr {
				t.Errorf("Info.CalculateTotalVCPUOfNodeList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Info.CalculateTotalVCPUOfNodeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetVCPUOfOwnNode(t *testing.T) {
	type args struct {
		podName string
	}
	tests := []struct {
		name    string
		i       Info
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GetVCPUOfOwnNode(tt.args.podName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Info.GetVCPUOfOwnNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Info.GetVCPUOfOwnNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetLimitVCPUOfOwnContainer(t *testing.T) {
	type args struct {
		podName string
	}
	tests := []struct {
		name    string
		i       Info
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GetLimitVCPUOfOwnContainer(tt.args.podName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Info.GetLimitVCPUOfOwnContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Info.GetLimitVCPUOfOwnContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetOwnContainerID(t *testing.T) {
	type args struct {
		podName string
	}
	tests := []struct {
		name string
		i    Info
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.GetOwnContainerID(tt.args.podName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Info.GetOwnContainerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetKubernetesID(t *testing.T) {
	tests := []struct {
		name string
		i    Info
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.GetKubernetesID(); got != tt.want {
				t.Errorf("Info.GetKubernetesID() = %v, want %v", got, tt.want)
			}
		})
	}
}
