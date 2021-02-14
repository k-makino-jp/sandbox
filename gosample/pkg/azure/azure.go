// azure is sample package
package azure

import (
	"context"
	"encoding/json"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
)

type azure interface {
	enqueue()
}

type azureImpl struct{}

type enqueuedMessage struct {
	Data string `json:"data"`
}

func (a *azureImpl) enqueue(queueEndpoint string, message enqueuedMessage) error {
	// create json data
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}
	jsonMessage := string(jsonBytes)

	// Create a request pipeline that is used to process HTTP(S) requests and responses..
	po := azqueue.PipelineOptions{
		Retry: azqueue.RetryOptions{
			Policy:        azqueue.RetryPolicyExponential, // Use exponential backoff as opposed to fixed
			MaxTries:      3,                              // Try at most 3 times to perform the operation (set to 1 to disable retries)
			TryTimeout:    time.Second * 3,                // Maximum time allowed for any single try
			RetryDelay:    time.Second * 1,                // Backoff amount for each retry (exponential or fixed)
			MaxRetryDelay: time.Second * 3,                // Max delay between retries
		},
	}
	pipeline := azqueue.NewPipeline(
		azqueue.NewAnonymousCredential(),
		po,
	)

	// Storage account queue service URL endpoint
	u, err := url.Parse(queueEndpoint)
	if err != nil {
		return err
	}

	// Create an ServiceURL object
	// serviceURL := azqueue.NewServiceURL(*u, pipeline)

	// All HTTP operations allow you to specify a Go context.Context object to control cancellation/timeout.
	ctx := context.TODO()

	// Create a URL
	queueURL := azqueue.NewQueueURL(*u, pipeline)
	// queueURL := serviceURL.NewQueueURL()

	// Create a URL
	messageURL := queueURL.NewMessagesURL()

	_, err = messageURL.Enqueue(ctx, jsonMessage, time.Second*0, time.Minute)
	if err != nil {
		return err
	}
	return nil
}

func NewAzure() *azureImpl {
	return &azureImpl{}
}
