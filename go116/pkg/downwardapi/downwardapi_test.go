// Package downwardapi Pod情報取得処理用パッケージ

package downwardapi

import (
	"errors"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_downwardAPI_GetPodInfo(t *testing.T) {
	// variables
	errIoUtilReadFile := errors.New("ioUtil.ReadFile error occurred")
	// create controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// // create mock
	mockIoUtil := NewMockioUtil(ctrl)
	tests := []struct {
		name      string
		d         downwardAPI
		want      PodInfo
		wantErr   error
		testSetup func()
	}{
		{
			name:    "downwardAPI.GetPodInfo Pod情報を取得するとき Pod情報が返ること",
			d:       downwardAPI{podNameFilePath: podNameFilePath, namespaceFilePath: namespaceFilePath, ioUtil: mockIoUtil},
			want:    PodInfo{PodName: "podname", Namespace: "namespace"},
			wantErr: nil,
			testSetup: func() {
				gomock.InOrder(
					mockIoUtil.EXPECT().ReadFile(podNameFilePath).Return([]byte("podname"), nil).Times(1),
					mockIoUtil.EXPECT().ReadFile(namespaceFilePath).Return([]byte("namespace"), nil).Times(1),
				)
			},
		},
		{
			name:    "downwardAPI.GetPodInfo Pod情報ファイル読み込みに失敗したとき Errorが返ること",
			d:       downwardAPI{podNameFilePath: podNameFilePath, namespaceFilePath: namespaceFilePath, ioUtil: mockIoUtil},
			want:    PodInfo{},
			wantErr: errIoUtilReadFile,
			testSetup: func() {
				gomock.InOrder(
					mockIoUtil.EXPECT().ReadFile(podNameFilePath).Return(nil, errIoUtilReadFile).Times(1),
				)
			},
		},
		{
			name:    "downwardAPI.GetPodInfo Namespace情報ファイル読み込みに失敗したとき Errorが返ること",
			d:       downwardAPI{podNameFilePath: podNameFilePath, namespaceFilePath: namespaceFilePath, ioUtil: mockIoUtil},
			want:    PodInfo{},
			wantErr: errIoUtilReadFile,
			testSetup: func() {
				gomock.InOrder(
					mockIoUtil.EXPECT().ReadFile(podNameFilePath).Return([]byte("podname"), nil).Times(1),
					mockIoUtil.EXPECT().ReadFile(namespaceFilePath).Return(nil, errIoUtilReadFile).Times(1),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			got, err := tt.d.GetPodInfo()
			if err != tt.wantErr {
				t.Errorf("downwardAPI.GetPodInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("downwardAPI.GetPodInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDonwardAPI(t *testing.T) {
	tests := []struct {
		name string
		want *downwardAPI
	}{
		{
			name: "NewDownwardAPI インスタンスを生成するとき インスタンスが返ること",
			want: &downwardAPI{
				podNameFilePath:   podNameFilePath,
				namespaceFilePath: namespaceFilePath,
				ioUtil:            ioUtilImpl{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDonwardAPI(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDonwardAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}
