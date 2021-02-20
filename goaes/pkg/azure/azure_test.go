package azure

import (
	"context"
	"encoding/json"
	"gosample/pkg/azure/mock_azure"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
	"github.com/golang/mock/gomock"
)

func Test_azure_Enqueue(t *testing.T) {
	// create controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// create mock
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
	// Configure sas
	URL, _ := url.Parse(queueEndpoint)
	query := url.Values{}
	query.Set("sv", "sv")
	URL.RawQuery = query.Encode()

	// Create a messagesURL object
	// XxxURL object contains a URL and a request pipeline.
	sampleMessagesURL := azqueue.NewMessagesURL(*URL, pipeline)
	type args struct {
		sas Sas
	}
	tests := []struct {
		name string
		args args
		want *azure
	}{
		{
			name: "NewAzure インスタンスを生成するとき インスタンスが返ること",
			args: args{
				sas: Sas{
					Sv: "sv",
				},
			},
			want: &azure{
				messagesURLEnqueue: sampleMessagesURL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAzure(tt.args.sas); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAzure() = %v, want %v", got, tt.want)
			}
		})
	}
}
