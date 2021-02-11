// http is sample package
package http

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type HttpClient interface {
	Get()
}

type httpClient struct {
	client         resty.Client
	retryCount     int
	retryWaitMin   time.Duration
	retryWaitMax   time.Duration
	requestTimeout time.Duration
}

// Get HTTP Request method
// status code == 429 or >= 500: do retry
func (h httpClient) Get(
	baseURL string, header, query map[string]string) (
	respBody string, err error, statusCode int) {
	// Configure a Resty Client
	resp, err := h.client.
		// SetLogger(httpLogger{}).
		SetRetryCount(h.retryCount).
		SetRetryWaitTime(h.retryWaitMin).
		SetRetryMaxWaitTime(h.retryWaitMax).
		SetTimeout(h.requestTimeout).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				statusCode := r.StatusCode()
				if statusCode == http.StatusTooManyRequests || statusCode >= 500 {
					return true
				}
				return false
			},
		).
		R().
		SetQueryParams(query).
		SetHeaders(header).
		Get(baseURL)
	return string(resp.Body()), err, resp.StatusCode()
}

var restyNew = resty.New

func NewHttpClient(retryCount int, retryWaitMin, retryWaitMax, requestTimeout time.Duration) *httpClient {
	return &httpClient{
		retryCount:     retryCount,
		retryWaitMin:   retryWaitMin,
		retryWaitMax:   retryWaitMax,
		requestTimeout: requestTimeout,
		client:         *restyNew(),
	}
}
