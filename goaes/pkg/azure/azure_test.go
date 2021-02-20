package azure

import (
	"context"
	"encoding/json"
	"errors"
	"gosample/pkg/azure/mock_azure"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
	"github.com/golang/mock/gomock"
)

func Test_azure_InitPipeline(t *testing.T) {
	sampleSas := Sas{}
	URL, _ := url.Parse(queueEndpoint)
	query := url.Values{}
	query.Set("sv", sampleSas.Sv)
	URL.RawQuery = query.Encode()
	type args struct {
		sas Sas
	}
	tests := []struct {
		name      string
		a         *azure
		args      args
		wantAzure *azure
	}{
		{
			name: "azure.InitPipeline Pipelineを初期化するとき Pipelineが初期化されること",
			a:    &azure{},
			args: args{sas: sampleSas},
			wantAzure: &azure{
				messagesURLEnqueue:     azqueue.NewMessagesURL(*URL, pipeline),
				enqueueMessageResponse: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.InitPipeline(tt.args.sas)
			if !reflect.DeepEqual(tt.a, tt.wantAzure) {
				t.Errorf("InitPipeline() got = %v, wantAzure %v", tt.a, tt.wantAzure)
			}
		})
	}
}

func Test_azure_Enqueue(t *testing.T) {
	// create controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// // create mock
	mockMessagesURLEnqueue := mock_azure.NewMockmessagesURLEnqueue(ctrl)
	mockEnqueueMessageResponse := mock_azure.NewMockenqueueMessageResponse(ctrl)
	type args struct {
		message Message
	}
	tests := []struct {
		name         string
		a            azure
		args         args
		want         int
		wantErr      bool
		testSetup    func()
		testTeardown func()
	}{
		{
			name: "azure.Enqueue AzureQueueStorageにEnqueueするとき 正常にEnqueueが完了すること",
			a: azure{
				messagesURLEnqueue:     mockMessagesURLEnqueue,
				enqueueMessageResponse: mockEnqueueMessageResponse,
			},
			args:    args{message: Message{}},
			want:    201,
			wantErr: false,
			testSetup: func() {
				jsonBytes, _ := json.Marshal(Message{})
				messageText := string(jsonBytes)
				ctx := context.TODO()
				visibilityTimeout := time.Second * 0
				timeToLive := time.Minute
				sampleEnqueueMessageResponse := azqueue.EnqueueMessageResponse{}
				mockMessagesURLEnqueue.
					EXPECT().
					Enqueue(ctx, messageText, visibilityTimeout, timeToLive).
					Return(&sampleEnqueueMessageResponse, nil).
					Times(1)
				mockEnqueueMessageResponse.
					EXPECT().
					StatusCode(&sampleEnqueueMessageResponse).
					Return(201).
					Times(1)
			},
			testTeardown: func() {},
		},
		{
			name: "azure.Enqueue リクエストに失敗したとき Errorが返ること",
			a: azure{
				messagesURLEnqueue:     mockMessagesURLEnqueue,
				enqueueMessageResponse: mockEnqueueMessageResponse,
			},
			args:    args{message: Message{}},
			want:    0,
			wantErr: true,
			testSetup: func() {
				jsonBytes, _ := json.Marshal(Message{})
				messageText := string(jsonBytes)
				ctx := context.TODO()
				visibilityTimeout := time.Second * 0
				timeToLive := time.Minute
				sampleEnqueueMessageResponse := azqueue.EnqueueMessageResponse{}
				mockMessagesURLEnqueue.
					EXPECT().
					Enqueue(ctx, messageText, visibilityTimeout, timeToLive).
					Return(&sampleEnqueueMessageResponse, errors.New("Azure Queue Storage Error Occurred")).
					Times(1)
				mockEnqueueMessageResponse.
					EXPECT().
					StatusCode(&sampleEnqueueMessageResponse).
					Return(0).
					Times(1)
			},
			testTeardown: func() {},
		},
		{
			name: "azure.Enqueue StatusCode201以外が返るとき Errorが返ること",
			a: azure{
				messagesURLEnqueue:     mockMessagesURLEnqueue,
				enqueueMessageResponse: mockEnqueueMessageResponse,
			},
			args:    args{message: Message{}},
			want:    500,
			wantErr: true,
			testSetup: func() {
				jsonBytes, _ := json.Marshal(Message{})
				messageText := string(jsonBytes)
				ctx := context.TODO()
				visibilityTimeout := time.Second * 0
				timeToLive := time.Minute
				sampleEnqueueMessageResponse := azqueue.EnqueueMessageResponse{}
				mockMessagesURLEnqueue.
					EXPECT().
					Enqueue(ctx, messageText, visibilityTimeout, timeToLive).
					Return(&sampleEnqueueMessageResponse, nil).
					Times(1)
				mockEnqueueMessageResponse.
					EXPECT().
					StatusCode(&sampleEnqueueMessageResponse).
					Return(500).
					Times(1)
			},
			testTeardown: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			got, err := tt.a.Enqueue(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("azure.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("azure.Enqueue() = %v, want %v", got, tt.want)
			}
			tt.testTeardown()
		})
	}
}

func TestNewAzure(t *testing.T) {
	tests := []struct {
		name string
		want *azure
	}{
		{
			name: "NewAzure インスタンスを生成するとき インスタンスが返ること",
			want: &azure{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAzure(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAzure() = %v, want %v", got, tt.want)
			}
		})
	}
}
