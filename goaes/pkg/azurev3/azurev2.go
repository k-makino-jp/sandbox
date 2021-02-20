// azure is sample package
package azurev2

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

// https://pkg.go.dev/github.com/Azure/azure-storage-queue-go/azqueue#MessagesURL.Enqueue
type messagesURLEnqueue interface {
	// azqueue.MessagesURL.Enqueue
	Enqueue(
		ctx context.Context,
		messageText string,
		visibilityTimeout time.Duration,
		timeToLive time.Duration,
	) (*azqueue.EnqueueMessageResponse, error)
}

// https://pkg.go.dev/github.com/Azure/azure-storage-queue-go/azqueue#EnqueueMessageResponse
// type enqueueMessageResponse interface {
// 	Date() time.Time
// 	RequestID() string
// 	Response() *http.Response
// 	Status() string
// 	StatusCode() int
// 	Version() string
// }

type azure struct {
	messagesURLEnqueue messagesURLEnqueue
}

// type assertion
func isEnqueueMessageResponsePointer(interfaceImpl interface{}) bool {
	_, ok := interfaceImpl.(*enqueueMessageResponse)
	return ok
}

func (a azure) convertStructToJSONText(v interface{}) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func (a azure) enqueue(messageText string) (*azqueue.EnqueueMessageResponse, error) {
	// All HTTP operations allow you to specify a Go context.Context object to control cancellation/timeout.
	ctx := context.TODO()

	// https://pkg.go.dev/github.com/Azure/azure-storage-queue-go/azqueue#MessagesURL.Enqueue
	visibilityTimeout := time.Second * 0
	timeToLive := time.Minute
	enqueueMessageResponse, err := a.messagesURLEnqueue.Enqueue(
		ctx,
		messageText,
		visibilityTimeout,
		timeToLive, // メッセージの有効期限
	)
	fmt.Println(enqueueMessageResponse.StatusCode())
	return nil, err
}

// func (a azure) checkStatusCode(enqueueMessageResponse azqueue.EnqueueMessageResponse) (int, error) {
// 	statusCode := enqueueMessageResponse.StatusCode()
// 	if statusCode == http.StatusCreated {
// 		return statusCode, nil
// 	}
// 	return statusCode, errors.New("Server connection error occurred")
// }

// Example:
// https://github.com/Azure/azure-storage-queue-go/blob/master/azqueue/zt_examples_test.go#L25
// Put messages: https://docs.microsoft.com/ja-jp/rest/api/storageservices/put-message
func (a azure) Enqueue(message Message) (int, error) {
	messageText, err := a.convertStructToJSONText(message)
	if err != nil {
		return 0, err
	}
	var e enqueueMessageResponse
	e, err = a.enqueue(messageText)
	if err != nil {
		return 0, err
	}
	en := e.(azqueue.EnqueueMessageResponse)
	return a.checkStatusCode(en)
}

// NewAzure コンストラクタ
func NewAzure(sas Sas) *azure {
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

	// Configure sas
	URL, _ := url.Parse(queueEndpoint)
	query := url.Values{}
	query.Set("sv", sas.Sv)
	URL.RawQuery = query.Encode()

	// Create a messagesURL object
	// XxxURL object contains a URL and a request pipeline.
	messagesURL := azqueue.NewMessagesURL(*URL, pipeline)
	return &azure{messagesURL}
}

type Sas struct {
	Sv string
}
