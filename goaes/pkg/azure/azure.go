// Package azure REST
package azure

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
)

var (
	queueEndpoint = "https://myaccount.queue.core.windows.net/images-to-download"
)

// Azure Queue Storage処理向けインターフェース
type Azure interface {
	Enqueue()
}

type messagesURLEnqueue interface {
	// azqueue.MessagesURL.Enqueue
	Enqueue(
		ctx context.Context,
		messageText string,
		visibilityTimeout time.Duration,
		timeToLive time.Duration,
	) (*azqueue.EnqueueMessageResponse, error) // cannot convert struct to interface
}

type enqueueMessageResponse interface {
	// use azqueue.EnqueueMessageResponse
	Response(*azqueue.EnqueueMessageResponse) []byte
	StatusCode(*azqueue.EnqueueMessageResponse) int
}

type azure struct {
	messagesURLEnqueue     messagesURLEnqueue
	enqueueMessageResponse enqueueMessageResponse
}

// Enqueue Azure Queue Storage REST API "Put messages"
func (a azure) Enqueue(message Message) (int, error) {
	// convert struct to json text
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return 0, err
	}
	messageText := string(jsonBytes)

	// Put messages
	ctx := context.TODO()
	visibilityTimeout := time.Second * 0
	timeToLive := time.Minute
	enqueueMessageResponse, err := a.messagesURLEnqueue.Enqueue(
		ctx,
		messageText,
		visibilityTimeout, // 可視性タイムアウト
		timeToLive,        // メッセージの有効期限
	)
	if err != nil {
		return 0, err
	}
	// status code check
	statusCode := a.enqueueMessageResponse.StatusCode(enqueueMessageResponse)
	if statusCode == http.StatusCreated {
		return statusCode, nil
	}
	return statusCode, errors.New("Server connection error occurred")
}

var (
	// Create a request pipeline object configured with credentials and with pipeline options. Once created,
	// a pipeline object is goroutine-safe and can be safely used with many XxxURL objects simultaneously.
	pipeline = azqueue.NewPipeline(
		azqueue.NewAnonymousCredential(), // A pipeline always requires some credential object
		azqueue.PipelineOptions{
			Retry: azqueue.RetryOptions{
				Policy:        azqueue.RetryPolicyExponential, // Use exponential backoff as opposed to fixed
				MaxTries:      3,                              // Try at most 3 times to perform the operation (set to 1 to disable retries)
				TryTimeout:    time.Second * 3,                // Maximum time allowed for any single try
				RetryDelay:    time.Second * 1,                // Backoff amount for each retry (exponential or fixed)
				MaxRetryDelay: time.Second * 3,                // Max delay between retries
			},
		},
	)
)

// NewAzure コンストラクタ
func NewAzure(sas Sas) *azure {
	// Configure sas
	URL, _ := url.Parse(queueEndpoint)
	query := url.Values{}
	query.Set("sv", sas.Sv)
	URL.RawQuery = query.Encode()

	// Create a messagesURL object
	// XxxURL object contains a URL and a request pipeline.
	messagesURL := azqueue.NewMessagesURL(*URL, pipeline)
	return &azure{
		messagesURLEnqueue: messagesURL,
	}
}
