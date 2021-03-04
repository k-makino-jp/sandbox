// Package kubernetes Kubernetes REST API処理用パッケージ

package kubernetes

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestInfo_GetNamespaceKubeSystemID(t *testing.T) {
	namespaceKubeSystem := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{UID: "kube-system"}}
	tests := []struct {
		name string
		i    Info
		want string
	}{
		{
			name: "Info.GetNamespaceKubeSystemID",
			i:    Info{NamespaceKubeSystem: namespaceKubeSystem},
			want: "kube-system",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.GetNamespaceKubeSystemID(); got != tt.want {
				t.Errorf("Info.GetNamespaceKubeSystemID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_CalculateTotalVCPUOfNodeList(t *testing.T) {
	nodeList := &corev1.NodeList{
		Items: []corev1.Node{
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0001"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("100m")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0002"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("1000m")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0003"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("1")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0004"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("12.345")}},
			},
		},
	}
	tests := []struct {
		name string
		i    Info
		want int64
	}{
		{
			name: "Info.CalculateTotalVCPUOfNodeList TotalVCPUが返るとき TotalVCPUが返ること",
			i:    Info{NodeList: nodeList},
			want: 100 + 1000 + 1000 + 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.CalculateTotalVCPUOfNodeList()
			if got != tt.want {
				t.Errorf("Info.CalculateTotalVCPUOfNodeList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetNumberOfNodes(t *testing.T) {
	nodeList := &corev1.NodeList{
		Items: []corev1.Node{
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0001"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("100m")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0002"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("1000m")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0003"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("1")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0004"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("12.345")}},
			},
		},
	}
	tests := []struct {
		name string
		i    Info
		want int
	}{
		{
			name: "Info.GetNumberOfNodes Node数が返るとき Node数が返ること",
			i:    Info{NodeList: nodeList},
			want: len(nodeList.Items),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.GetNumberOfNodes(); got != tt.want {
				t.Errorf("Info.GetNumberOfNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetNamespacePodID(t *testing.T) {
	namespacePod := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{UID: "pod-namespace-id"}}
	tests := []struct {
		name string
		i    Info
		want string
	}{
		{
			name: "Info.GetNamespacePodID",
			i:    Info{NamespacePod: namespacePod},
			want: "pod-namespace-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.GetNamespacePodID(); got != tt.want {
				t.Errorf("Info.GetNamespacePodID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetVCPUOfOwnNode(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "podname"},
		Spec:       corev1.PodSpec{NodeName: "rhocp-east-0001-app-0003"},
	}
	nodeList := &corev1.NodeList{
		Items: []corev1.Node{
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0001"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("100m")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0002"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("1234m")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0003"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("1")}},
			},
			{
				ObjectMeta: metav1.ObjectMeta{Name: "rhocp-east-0001-app-0004"},
				Status:     corev1.NodeStatus{Capacity: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("12.345")}},
			},
		},
	}
	tests := []struct {
		name string
		i    Info
		want int64
	}{
		{
			name: "Info.GetVCPUOfOwnNode コンテナが属するNodeのvCPU数を取得するとき コンテナが属するNodeのvCPU数が返ること",
			i:    Info{Pod: pod, NodeList: nodeList},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.GetVCPUOfOwnNode()
			if got != tt.want {
				t.Errorf("Info.GetVCPUOfOwnNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetOwnContainerID(t *testing.T) {
	cgroupBytes := []byte(`
12:rdma:/
11:blkio:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
10:perf_event:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
9:cpuset:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
8:freezer:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
7:devices:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
6:memory:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
5:net_cls,net_prio:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
4:pids:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
3:cpu,cpuacct:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
2:hugetlb:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
1:name=systemd:/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podd4977652_ffc4_424b_90c0_bbb5e35e5ea3.slice/docker-abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02.scope
0::/system.slice/containerd.service
`)
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "podname"},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{ContainerID: "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz01"},
				{ContainerID: "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02"},
				{ContainerID: "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz03"},
				{ContainerID: "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz04"},
			},
		},
	}
	// variables
	errIoUtilReadFile := errors.New("ioUtil.ReadFile error occurred")
	// configure mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIoUtil := NewMockioUtil(ctrl)

	tests := []struct {
		name      string
		i         Info
		want      string
		wantErr   bool
		testSetup func()
	}{
		{
			name:      "Info.GetOwnContainerID コンテナIDを取得するとき コンテナIDが返ること",
			i:         Info{Pod: pod, ioUtil: mockIoUtil},
			want:      "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02",
			wantErr:   false,
			testSetup: func() { mockIoUtil.EXPECT().ReadFile(cgroupFilePath).Return(cgroupBytes, nil).Times(1) },
		},
		{
			name:      "Info.GetOwnContainerID cgroupファイル読み込みに失敗するとき 予期しないエラーが返ること",
			i:         Info{Pod: pod, ioUtil: mockIoUtil},
			want:      "",
			wantErr:   true,
			testSetup: func() { mockIoUtil.EXPECT().ReadFile(cgroupFilePath).Return(nil, errIoUtilReadFile).Times(1) },
		},
		{
			name:      "Info.GetOwnContainerID コンテナIDの照合に失敗するとき 予期しないエラーが返ること",
			i:         Info{Pod: pod, ioUtil: mockIoUtil},
			want:      "",
			wantErr:   true,
			testSetup: func() { mockIoUtil.EXPECT().ReadFile(cgroupFilePath).Return([]byte(""), nil).Times(1) },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			got, err := tt.i.GetOwnContainerID()
			if (err != nil) != tt.wantErr {
				t.Errorf("Info.GetOwnContainerID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Info.GetOwnContainerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetOwnContainerName(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "podname"},
		Status: corev1.PodStatus{
			ContainerStatuses: []corev1.ContainerStatus{
				{Name: "nginx-container-01", ContainerID: "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz01"},
				{Name: "nginx-container-02", ContainerID: "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02"},
			},
		},
	}
	type args struct {
		containerID string
	}
	tests := []struct {
		name    string
		i       Info
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Info.GetOwnContainerName コンテナ名を取得するとき コンテナ名が返ること",
			i:       Info{Pod: pod},
			args:    args{containerID: "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz02"},
			want:    "nginx-container-02",
			wantErr: false,
		},
		{
			name:    "Info.GetOwnContainerName 存在しないコンテナIDを与えたとき Errorが返ること",
			i:       Info{Pod: pod},
			args:    args{containerID: "docker://abcdefghijklmnopqrstuvwxyz1234567890abcdefghijklmnopqrstuvwxyz99"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GetOwnContainerName(tt.args.containerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Info.GetOwnContainerName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Info.GetOwnContainerName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetLimitVCPUOfOwnContainer(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "podname"},
		Status:     corev1.PodStatus{ContainerStatuses: []corev1.ContainerStatus{{Name: "nginx-container-01"}, {Name: "nginx-container-02"}}},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:      "nginx-container-01",
					Resources: corev1.ResourceRequirements{Limits: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("1000m")}},
				},
				{
					Name: "nginx-container-02",
				},
			},
		},
	}
	type args struct {
		containerName string
	}
	tests := []struct {
		name    string
		i       Info
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "Info.GetLimitVCPUOfOwnContainer コンテナのlimitVCPUを取得するとき コンテナのlimitVCPUが返ること",
			i:       Info{Pod: pod},
			args:    args{containerName: "nginx-container-01"},
			want:    1000,
			wantErr: false,
		},
		{
			name:    "Info.GetLimitVCPUOfOwnContainer 当該コンテナのlimit_vcpu未設定のとき 0が返ること",
			i:       Info{Pod: pod},
			args:    args{containerName: "nginx-container-02"},
			want:    0,
			wantErr: false,
		},
		{
			name:    "Info.GetLimitVCPUOfOwnContainer 存在しないコンテナ名を与えたとき Errorが返ること",
			i:       Info{Pod: pod},
			args:    args{containerName: "nginx-container-99"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GetLimitVCPUOfOwnContainer(tt.args.containerName)
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
