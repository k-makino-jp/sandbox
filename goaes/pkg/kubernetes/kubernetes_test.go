// Package kubernetes Kubernetes REST API処理用パッケージ

package kubernetes

import (
	"errors"
	"gosample/pkg/downwardapi"
	"gosample/pkg/kubernetes/mock_kubernetes"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	fakecorev1 "k8s.io/client-go/kubernetes/typed/core/v1/fake"
	"k8s.io/client-go/rest"
	k8stesting "k8s.io/client-go/testing"
)

func Test_kubernetes_Init(t *testing.T) {
	// variables
	errInClusterConfig := errors.New("InClusterConfig error occurred")
	errNewForConfig := errors.New("NewForConfig error occurred")
	// create controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClientGoRest := mock_kubernetes.NewMockclientGoRest(ctrl)
	mockclientGoKubernetes := mock_kubernetes.NewMockclientGoKubernetes(ctrl)
	// test
	tests := []struct {
		name      string
		k         *kubernetes
		wantErr   error
		testSetup func()
	}{
		{
			name: "kubernetes.Init 初期化するとき 構造体が初期化されること",
			k: &kubernetes{
				clientGoRest:       mockClientGoRest,
				clientGoKubernetes: mockclientGoKubernetes,
			},
			wantErr: nil,
			testSetup: func() {
				mockClientGoRest.EXPECT().InClusterConfig().Return(&rest.Config{}, nil).Times(1)
				mockclientGoKubernetes.EXPECT().NewForConfig(&rest.Config{}).Return(&k8s.Clientset{}, nil).Times(1)
			},
		},
		{
			name: "kubernetes.Init InClusterConfigに失敗するとき Errorが返ること",
			k: &kubernetes{
				clientGoRest: mockClientGoRest,
			},
			wantErr: errInClusterConfig,
			testSetup: func() {
				mockClientGoRest.EXPECT().InClusterConfig().Return(&rest.Config{}, errInClusterConfig).Times(1)
			},
		},
		{
			name: "kubernetes.Init NewForConfigに失敗するとき Errorが返ること",
			k: &kubernetes{
				clientGoRest:       mockClientGoRest,
				clientGoKubernetes: mockclientGoKubernetes,
			},
			wantErr: errNewForConfig,
			testSetup: func() {
				mockClientGoRest.EXPECT().InClusterConfig().Return(&rest.Config{}, nil).Times(1)
				mockclientGoKubernetes.EXPECT().NewForConfig(&rest.Config{}).Return(&k8s.Clientset{}, errNewForConfig).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			if err := tt.k.Init(); err != tt.wantErr {
				t.Errorf("kubernetes.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kubernetes_GetInfo(t *testing.T) {
	tests := []struct {
		name    string
		k       kubernetes
		want    Info
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.k.GetInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("kubernetes.GetInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kubernetes.GetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

// functions
// newFakeClient := func() *fake.Clientset {
// 	client := fake.NewSimpleClientset()
// 	pod := &corev1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{Namespace: corev1.NamespaceDefault, Name: "podname"},
// 	}
// 	client.CoreV1().Pods(corev1.NamespaceDefault).Create(context.TODO(), pod, metav1.CreateOptions{})
// pod = &corev1.Pod{
// 	ObjectMeta: metav1.ObjectMeta{Namespace: "kube-system", Name: "podname"},
// }
// 	client.CoreV1().Pods("kube-system").Create(context.TODO(), pod, metav1.CreateOptions{})
// 	return client
// }

// reactionFunc := func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
// 	fmt.Println("ACTION")
// 	switch {
// 	case action.Matches("get", "pods"):
// 		fmt.Println("ACTION")
// 		groupResource := schema.GroupResource{
// 			Group:    "",
// 			Resource: "",
// 		}
// 		return false, nil, k8serrors.NewForbidden(groupResource, "forbidden", errForbidden)
// 	}
// 	return true, pod, nil
// }
// fakeClient.AddReactor("get", "pods", reactionFunc)

func newFakeClient(verb, resourceName string, resource runtime.Object, err error) *fake.Clientset {
	clientset := fake.NewSimpleClientset()
	clientset.CoreV1().(*fakecorev1.FakeCoreV1).
		PrependReactor(verb, resourceName, func(action k8stesting.Action) (bool, runtime.Object, error) {
			return true, resource, err
		})
	return clientset
}

func Test_kubernetes_getPods(t *testing.T) {
	// variables
	errForbidden := k8serrors.NewForbidden(schema.GroupResource{}, "Error", errors.New("Forbidden error occurred"))
	// configure fake client
	pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Namespace: "kube-system", Name: "podname"}}
	// test
	tests := []struct {
		name    string
		k       *kubernetes
		wantErr error
	}{
		{
			name: "kubernetes.getPods Pod情報を取得するとき Pod情報が取得されること",
			k: &kubernetes{
				clientset: newFakeClient("get", "pods", pod, nil),
				podInfo:   downwardapi.PodInfo{PodName: "podname", Namespace: corev1.NamespaceDefault},
			},
			wantErr: nil,
		},
		{
			name: "kubernetes.getPods 403Forbiddenが発生するとき Errorが返ること",
			k: &kubernetes{
				clientset: newFakeClient("get", "pods", pod, errForbidden),
				podInfo:   downwardapi.PodInfo{PodName: "podname", Namespace: corev1.NamespaceDefault},
			},
			wantErr: errForbidden,
		},
		// {
		// 	name: "kubernetes.getPods 403Forbiddenが発生するとき Errorが返ること",
		// 	k: &kubernetes{
		// 		clientset: fake.NewSimpleClientset(),
		// 		podInfo:   downwardapi.PodInfo{PodName: "podname", Namespace: corev1.NamespaceDefault},
		// 	},
		// 	wantErr: errors.New(`pods "podname" not found`),
		// },
		// {
		// 	name: "kubernetes.getPods Pod情報の取得に失敗するとき Errorが返ること",
		// 	k: &kubernetes{
		// 		clientset: newFakeClient(),
		// 		podInfo:   downwardapi.PodInfo{PodName: "podname", Namespace: "kube-system"},
		// 	},
		// 	// wantErr: k8serrors.NewConflict(schema.GroupResource{Resource: "test"}, "other", nil),
		// 	wantErr: errors.New(`pods "podname" not found`),
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.getPods(); err != tt.wantErr {
				t.Errorf("kubernetes.getPods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kubernetes_getNodeList(t *testing.T) {
	tests := []struct {
		name    string
		k       *kubernetes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.getNodeList(); (err != nil) != tt.wantErr {
				t.Errorf("kubernetes.getNodeList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kubernetes_getNamespaceKubeSystem(t *testing.T) {
	tests := []struct {
		name    string
		k       *kubernetes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.getNamespaceKubeSystem(); (err != nil) != tt.wantErr {
				t.Errorf("kubernetes.getNamespaceKubeSystem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_kubernetes_getNamespacePod(t *testing.T) {
	tests := []struct {
		name    string
		k       *kubernetes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.getNamespacePod(); (err != nil) != tt.wantErr {
				t.Errorf("kubernetes.getNamespacePod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewKubernetes(t *testing.T) {
	type args struct {
		podInfo downwardapi.PodInfo
	}
	tests := []struct {
		name string
		args args
		want *kubernetes
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKubernetes(tt.args.podInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKubernetes() = %v, want %v", got, tt.want)
			}
		})
	}
}
