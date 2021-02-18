// azure is sample package
package azurev2

import (
	"context"
	"encoding/json"
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
	) (*azqueue.EnqueueMessageResponse, error)
}

type azure struct {
	enqueue messagesURLEnqueue
}

// Example:
// https://github.com/Azure/azure-storage-queue-go/blob/master/azqueue/zt_examples_test.go#L25
// Put messages: https://docs.microsoft.com/ja-jp/rest/api/storageservices/put-message
func (a azure) Enqueue(message Message) (int, error) {
	// create json data
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return 0, err
	}
	messageText := string(jsonBytes)

	// All HTTP operations allow you to specify a Go context.Context object to control cancellation/timeout.
	ctx := context.TODO()

	visibilityTimeout := time.Second * 0
	timeToLive := time.Minute
	// https://pkg.go.dev/github.com/Azure/azure-storage-queue-go/azqueue#MessagesURL.Enqueue
	enqueueMessageResponse, err := a.enqueue.Enqueue(
		ctx,
		messageText,
		visibilityTimeout,
		timeToLive, // メッセージの有効期限
	)
	statusCode := enqueueMessageResponse.StatusCode()
	if statusCode == http.StatusCreated {
		// fmt.Println(enqueueMessageResponse.Response())
		// Status returns the HTTP status message of the response, e.g. "200 OK".
		// fmt.Println(enqueueMessageResponse.StatusCode())
		return statusCode, nil
	} else if statusCode == http.StatusUnauthorized || statusCode == http.StatusForbidden {
		return statusCode, err
	}
	return statusCode, err
}

// NewAzure コンストラクタ
func NewAzure() *azure {
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
	url, _ := url.Parse(queueEndpoint)
	query := getSasQueryMap(sas{}, url.Query())
	url.RawQuery = query.Encode()

	// Create a messagesURL object
	// XxxURL object contains a URL and a request pipeline.
	messagesURL := azqueue.NewMessagesURL(*url, pipeline)
	return &azure{messagesURL}
}

type sas struct {
	st string
}

func getSasQueryMap(sas sas, query url.Values) url.Values {
	sasQueryMap := map[string]string{
		"sig": sas.st,
	}
	for key, value := range sasQueryMap {
		query.Set(key, value)
	}
	return query
}
