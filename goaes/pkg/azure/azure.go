// azure is sample package
package azure

import (
	"context"
	"encoding/json"
	"fmt"
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

type azure struct{}

// Example:
// https://github.com/Azure/azure-storage-queue-go/blob/master/azqueue/zt_examples_test.go#L25
// Put messages: https://docs.microsoft.com/ja-jp/rest/api/storageservices/put-message
func (a azure) Enqueue(message Message) error {
	// create json data
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	messageText := string(jsonBytes)

	// Create a request pipeline that is used to process HTTP(S) requests and responses..
	pipelineOption := azqueue.PipelineOptions{
		Retry: azqueue.RetryOptions{
			Policy:        azqueue.RetryPolicyExponential, // Use exponential backoff as opposed to fixed
			MaxTries:      3,                              // Try at most 3 times to perform the operation (set to 1 to disable retries)
			TryTimeout:    time.Second * 3,                // Maximum time allowed for any single try
			RetryDelay:    time.Second * 1,                // Backoff amount for each retry (exponential or fixed)
			MaxRetryDelay: time.Second * 3,                // Max delay between retries
		},
	}

	// Create a request pipeline object configured with credentials and with pipeline options. Once created,
	// a pipeline object is goroutine-safe and can be safely used with many XxxURL objects simultaneously.
	pipeline := azqueue.NewPipeline(
		azqueue.NewAnonymousCredential(), // A pipeline always requires some credential object
		pipelineOption,
	)

	// Storage account queue service URL endpoint
	url, err := url.Parse(queueEndpoint)
	if err != nil {
		return err
	}

	// All HTTP operations allow you to specify a Go context.Context object to control cancellation/timeout.
	ctx := context.TODO()

	// Create a URL
	messageURL := azqueue.NewMessagesURL(*url, pipeline)

	visibilityTimeout := time.Second * 0
	timeToLive := time.Minute
	// https://pkg.go.dev/github.com/Azure/azure-storage-queue-go/azqueue#MessagesURL.Enqueue
	_, err = messageURL.Enqueue(
		ctx,
		messageText,
		visibilityTimeout,
		timeToLive, // メッセージの有効期限
	)
	// fmt.Println(enqueueMessageResponse)
	// Status returns the HTTP status message of the response, e.g. "200 OK".
	// fmt.Println(enqueueMessageResponse.StatusCode())
	fmt.Println(err)
	//
	return err
}

// NewAzure コンストラクタ
func NewAzure() *azure {
	return &azure{}
}
